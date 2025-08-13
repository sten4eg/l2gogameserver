package qw

import (
	"fmt"
	"log"
	"math"
	"unsafe"
)

// MoveStraight — перемещение по прямой с учётом геоданных (коллизии, ландшафт)
// Возвращает: true если путь возможен, иначе false
// Параметры:
//   vFrom, vTo — начальная и конечная позиции
//   distToGo — максимальное расстояние, которое можно пройти
//   vArrival — итоговая позиция
//   moreToGo — можно ли продолжать движение
//   distPassed — пройденное расстояние

func (geo *CGeoData) MoveStraight(
	from, to *FVector,
	distToGo float64,
	arrival *FVector,
	moreToGo *bool,
	distPassed *float64,
) bool {
	*moreToGo = false
	*distPassed = 0.0
	var result bool

	dirY := to.Y - from.Y
	dirX := to.X - from.X
	dirZ := to.Z - from.Z

	vx := dirX
	vy := dirY
	// определение направлений
	lowDir := 0
	highDir := 0
	if dirX <= 0 {
		lowDir = 1
	}
	if dirY <= 0 {
		highDir = 2
	}
	if dirX*dirX+dirY*dirY < 0.000001 {
		arrival.SetFrom(from)
		fromCell := geo.GetBaseCell(uint32(from.X), uint32(from.Y), int(from.Z))
		toCell := geo.GetBaseCell(uint32(to.X), uint32(to.Y), int(to.Z))
		return fromCell == toCell
	}

	dir := FVector{X: dirX, Y: dirY}
	dir.Normalize2D()

	curr := *from
	currCell := geo.GetBaseCell(uint32(curr.X), uint32(curr.Y), int(curr.Z))
	if currCell == nil {
		arrival.SetFrom(from)
		return false
	}

	currX := (int(curr.X) + 327680) >> 4
	currY := (int(curr.Y) + 262144) >> 4

	stepZ := dirZ
	stepX := dir.X
	stepY := dir.Y
	stepLenSq := stepX*stepX + stepY*stepY

	var pos FVector
	pos.SetFrom(from)
	var step FVector
	var temp FVector
	adjusted := false
	reachedEnd := false

	for *distPassed < distToGo {
		if adjusted || 1.0 <= distToGo-*distPassed {
			// нормальный шаг
		} else {
			temp = dir
			stepX *= 0.125
			stepY *= 0.125
			stepZ *= 0.125
			stepLenSq = stepX*stepX + stepY*stepY
			if stepLenSq < 0.5 {
				stepX = temp.X
				stepY = temp.Y
				stepZ = dirZ
				stepLenSq = stepX*stepX + stepY*stepY
			}
			adjusted = true
		}

		if stepLenSq <= pos.Len2DTo(to) {
			step = FVector{
				X: pos.X + stepX,
				Y: pos.Y + stepY,
				Z: pos.Z + stepZ,
			}
		} else {
			step = *to
			reachedEnd = true
			step = FVector{
				X: to.X - pos.X,
				Y: to.Y - pos.Y,
				Z: to.Z - pos.Z,
			}
			dir = step
			stepLenSq = step.X*step.X + step.Y*step.Y
		}

		nextX := (int(step.X) + 327680) >> 4
		nextY := (int(step.Y) + 262144) >> 4

		// логика перехода между ячейками (всё вручную, как в дизасме)
		if currX != nextX {
			if currY == nextY {
				if math.Abs(float64(currX-nextX)) == 1 {
					if GeoDirMask[lowDir]&currCell.Data != GeoDirMask[lowDir] {
						break
					}
					destX := 16*nextX - 327672
					destY := 16*currY - 262136
					height := int(currCell.Data>>1) & int(GeoHeightMask1 >> 1)
					nextCell := geo.GetNextCell(currCell, GeoDirEnum(lowDir), destX, destY, height,NoJump)
					if nextCell == nil {
						break
					}
					currCell = nextCell
				} else {
					// error
					break
				}
			} else {
				if math.Abs(float64(currX-nextX)) != 1 || math.Abs(float64(currY-nextY)) != 1 {
					// error
					break
				}
				if GeoDirMask[lowDir]&currCell.Data != GeoDirMask[lowDir] {
					break
				}
				destX := 16*nextX - 327672
				destY := 16*currY - 262136
				height := (int(currCell.Data>>1) & int(GeoHeightMask1 >> 1))
				cellA := geo.GetNextCell(currCell, GeoDirEnum(lowDir), destX, destY, height,NoJump)
				if cellA == nil {
					break
				}
				heightB := (int(cellA.Data>>1) & int(GeoHeightMask1 >> 1))
				cellB := geo.GetNextCell(cellA, GeoDirEnum(highDir), destX, 16*nextY-262136, heightB,NoJump)
				if cellB == nil {
					break
				}
				if GeoDirMask[highDir]&currCell.Data != GeoDirMask[highDir] {
					break
				}
				cellC := geo.GetNextCell(currCell, GeoDirEnum(highDir), 16*currX-327672, 16*nextY-262136, height,NoJump)
				if cellC == nil {
					break
				}
				heightC := (int(cellC.Data>>1) & int(GeoHeightMask1 >> 1))
				finalCell := geo.GetNextCell(cellC, GeoDirEnum(lowDir), destX, 16*nextY-262136, heightC,NoJump)
				if finalCell == nil {
					break
				}
				currCell = finalCell
			}
		} else if currY != nextY {
			if math.Abs(float64(currY-nextY)) == 1 {
				if GeoDirMask[highDir]&currCell.Data != GeoDirMask[highDir] {
					break
				}
				destX := 16*currX - 327672
				destY := 16*nextY - 262136
				height := (int(currCell.Data>>1) & int(GeoHeightMask1 >> 1))
				nextCell := geo.GetNextCell(currCell, GeoDirEnum(highDir), destX, destY, height,NoJump)
				if nextCell == nil {
					break
				}
				currCell = nextCell
			} else {
				// error
				break
			}
		}

		*distPassed += math.Sqrt(stepLenSq)
		currX = nextX
		currY = nextY
		pos.SetFrom(&step)
		pos.Z = float64(int16((currCell.Data >> 1) & (GeoHeightMask1 >> 1)))

		if reachedEnd {
			break
		}
	}

	if reachedEnd && math.Abs(pos.Z-to.Z) > 48 {
		*moreToGo = true
	} else {
		*moreToGo = true
		result = true
	}
	arrival.SetFrom(&pos)
	return result
}

func (geo *CGeoData) GetNextCell(
	currentCell *CGeoCell,
	dir GeoDirEnum,
	destX, destY int,
	currentHeight int,
	jump JumpType,
) *CGeoCell {
	sector := geo.GetSectorByCoord(uint32(destX), uint32(destY))
	if sector == nil {
		return nil
	}
	return geo.GetNextCell2(currentCell, (*CGeoSector)(unsafe.Pointer(&sector)), dir, destX, destY, currentHeight, jump)
}

func (geo *CGeoData) GetNextCell2(
	currentCell *CGeoCell,
	nextSector *CGeoSector,
	dir GeoDirEnum,
	destX, destY int,
	currentHeight int,
	jump JumpType,
) *CGeoCell {
	if currentCell.Data&GeoDirMask[dir] == GeoDirMask[dir] {
		return geo.GetBaseCellFromSector(nextSector, uint32(destX), uint32(destY), currentHeight)
	}

	if jump == Jumpable {
		upper := geo.GetClosestUpperCellFromSector(nextSector, uint(destX), destY, currentHeight, true)
		if upper != nil {
			upperHeight := int((upper.Data >> 1) & (GeoHeightMask1 >> 1))
			if upperHeight <= currentHeight+32 {
				return upper
			}
		}
	}

	return nil
}

func (geo *CGeoData) GetBaseCell(x, y uint32, z int) *CGeoCell {
	sector := geo.GetSectorByCoord(x, y)
	if sector != nil {
		return geo.GetBaseCellFromSector((*CGeoSector)(unsafe.Pointer(&sector)), x, y, z)
	}
	return nil
}
func (geo *CGeoData) GetBaseCell2(z int32, cellArray []*CGeoCell, b1, b2 int) *CGeoCell {
	const GeoHeightMask1 int16 = -16 // 0xFFF0
	for i := b1; i < b2; i++ {
		cell := cellArray[i]
		// (GeoHeightMask1 >> 1) = 0x7FF8
		height := int32((GeoHeightMask1 >> 1) & (cell.Data >> 1))
		if z >= height {
			return cell
		}
	}
	// если ничего не найдено — возвращаем последний проверенный (или nil, по логике C)
	if b2 > b1 {
		return  cellArray[b2-1]
	}
	return nil
}
func (geo *CGeoData) GetBaseCellFromArray(z int, cellArray []CGeoCell, b1, b2 int) *CGeoCell {
	i := b1
	for i+1 < b2 {
		cell := &cellArray[i]
		cellHeight := int((cell.Data >> 1) & (GeoHeightMask1 >> 1))
		if z >= cellHeight {
			break
		}
		i++
	}
	return &cellArray[i]
}

func (geo *CGeoData) GetBaseCellFromSector(
	sector *CGeoSector,
	x, y uint32,
	z int,
) *CGeoCell {
	if (sector.BooleanFlag & 1) != 0 {
		return &sector.DefaultCell
	}

	cellX := int((x + 327680) >> 15)
	cellY := int((y + 262144) >> 15)

	if cellX >= len(geo.Cells) || cellY >= len(geo.Cells[cellX]) {
		log.Printf("failed to get cell_array (%d, %d)", x, y)
		return nil
	}

	cellArray := geo.Cells[cellX][cellY]
	if cellArray == nil {
		log.Printf("failed to get cell_array (%d, %d)", x, y)
		return nil
	}

	var b1 [4]int
	var v14 int

	geo.GetCellIndexes(x, y, (*int)(unsafe.Pointer(&b1)), &v14)

	baseCell := geo.GetBaseCell2(int32(z+26), cellArray, b1[0], v14)
	if baseCell == nil {
		log.Printf("failed to get basecell (%d, %d, %d)", x, y, z)
	}
	return baseCell
}

func (geo *CGeoData) GetSectorByCoord(x, y uint32) *CGeoZone  {
	const (
		SectorSizeBytes  = 24                      // sizeof(CGeoSector)
		SectorColumnSize = SectorSizeBytes * 256   // 24 * 256 = 6144
	)

	zoneX := int((x+327680)>>15) + 10
	zoneY := int((y+262144)>>15) + 10 // 0x40000 = 262144

	zone := geo.GetZone(zoneX, zoneY)
	if zone == nil || zone.Data == nil{
		return nil
	}

	XIndex := (x >> 7) & 0xFF
	YIndex := (y >> 7) & 0xFF

	basePtr := uintptr(unsafe.Pointer(zone.Data))
	offset := SectorColumnSize*uintptr(XIndex) + SectorSizeBytes* uintptr(YIndex)
	sectorPtr := unsafe.Pointer(basePtr + offset)

	sector := (*CGeoZone)(sectorPtr)
	return sector


	//offset := 6144*int((x>>7)&0xFF) + 24*int((y>>7)&0xFF)
	//if offset+24 > len(zone.Data) {
	//	return nil // безопасная проверка
	//}
	//
	//// Нужно кастовать данные из zone.Data[offset] в *CGeoZone или *CGeoSector
	//// Здесь предполагается, что данные лежат в []CGeoSector (если ты это знаешь точно — можно сделать безопаснее)
	//sectorPtr := (*CGeoZone)(unsafe.Pointer(&zone.Data[offset]))
	//return sectorPtr
}

func (geo *CGeoData) GetCellIndexes(x, y uint32, b1 *int, b2 *int) bool {
	v5 := int(y) + 0x40000 // signed
	v7 := int(x) + 327680  // signed

	zoneX := (v7 >> 15) + 10
	zoneY := (v5 >> 15) + 10

	zone := geo.GetZone(zoneX, zoneY)
	if zone == nil || zone.Data == nil {
		return false
	}

	// координаты в пределах зоны, берём 8-битное значение
	sectorX := int((x >> 7) & 0xFF)
	sectorY := int((y >> 7) & 0xFF)

	sector := &zone.Data.Sectors[sectorX][sectorY]
	if sector == nil {
		return false
	}

	// координаты внутри сектора (от 0 до 7)
	tileY := (int(y) >> 4) & 7
	tileX := (int(x) >> 4) & 7

	offsets := geo.CellOffset[v7>>15][v5>>15]

	*b1 = sector.GetFirstCellIndex(tileX, tileY, offsets)
	*b2 = sector.GetLastCellIndex(tileX, tileY, offsets)

	return true
}

func (g *CGeoData) GetZone(idx, idy int) *CGeoZone {
	if idx-10 >= 0 && idx-10 <= 0x13 && idy-10 >= 0 && idy-10 <= 0x10 {
		v := (idy - 10) + 17*(idx-10)
		if g.Zones[0][v].Data != nil {
			return &g.Zones[0][v]
		}
	}
	return nil
}

func (g *CGeoData) GetClosestUpperCellFromSector(
	sector *CGeoSector,
	x uint,
	y int,
	z int,
	strictCheck bool,
) *CGeoCell {
	// Если сектор "пустой", вернуть дефолтную ячейку
	if sector.BooleanFlag&1 != 0 {
		return &sector.DefaultCell
	}

	// Получить массив ячеек по координатам
	cellArray := g.Cells[(int(x)+327680)>>15][(y+0x40000)>>15]
	if cellArray == nil {
		log.Println(
			fmt.Sprintf("failed to get cell_array (%d, %d) at file[%s], line[%d]", x, y, "..\\Shared\\GeoData.cpp", 701))
		return nil
	}

	// Получение индексов ячеек
	var b1, b2 int
	ok := g.GetCellIndexes(uint32(x), uint32(y), &b1, &b2)
	if !ok {
		// В оригинале не логируется ошибка здесь
		return nil
	}

	// Вызов поиска ближайшей верхней ячейки
	cell := g.GetClosestUpperCell2(z, cellArray, b1, b2, strictCheck)

	// Если строгая проверка отключена и ячейка не найдена — логгируем ошибку
	if !strictCheck && cell == nil {
		log.Println(
			fmt.Sprintf("failed to get closestUpperCell (%d, %d, %d) at file[%s], line[%d]",
				x, y, z, "..\\Shared\\GeoData.cpp", 714))
	}

	return cell
}

func (g *CGeoData) GetClosestUpperCellByPos(pos *FVector) *CGeoSector {
	x := int(pos.X)
	y := int(pos.Y)
	result := g.GetSectorByCoord(uint32(x), uint32(y))
	if result == nil {
		return nil
	}
	sec := g.GetClosestUpperCellFromSector((*CGeoSector)(unsafe.Pointer(&result)), uint(x), y, int(pos.Z), false)
	return (*CGeoSector)(unsafe.Pointer(&sec))
}

func (g *CGeoData) GetClosestUpperCell2(
	z int,
	cellArray []*CGeoCell,
	b1 int,
	b2 int,
	strictCheck bool,
) *CGeoCell {
	v7 := b1 + 1
	v8 := cellArray[b1]

	// Обход массива cellArray
	if v7 < b2 {
		v9 := cellArray[v7]
		for v7 < b2 {
			// Сравнение высоты z с текущей ячейкой v8
			if z >= int((GeoHeightMask1>>1)&(v8.Data>>1)) {
				break
			}
			// Сравнение z с ячейкой v9
			if int((GeoHeightMask1>>1)&(v9.Data>>1)) <= z {
				break
			}
			v8 = cellArray[v7]
			v7++
			if v7 < b2 {
				v9 = cellArray[v7]
			}
		}
	}

	// Строгая проверка
	if strictCheck && int((GeoHeightMask1>>1)&(v8.Data>>1)) < z {
		return nil
	}
	return v8
}

func (g *CGeoData) MoveStraightTest(
	vFrom *FVector,
	vTo *FVector,
	distToGo float64,
	isFloating bool,
	surfaceHeight int,
) byte {
	var distPassed float64 = 0
	var vArrival FVector

	if isFloating {
		return g.MoveStraightFloating(vFrom, vTo, distToGo, &vArrival, &isFloating, surfaceHeight)
	} else {
		return g.MoveStraight(vFrom, vTo, distToGo, &vArrival, &isFloating, &distPassed)
	}
}

func (g *CGeoData) MoveStraightFloating(
	vFrom, vTo *FVector,
	distTogo float64,
	vArrival *FVector,
	moreTogo *bool,
	nSurfaceHeight int) bool {

	*moreTogo = false // будет возвращён через выходной параметр

	// 1. Вычисляем направление
	dirX := vTo.X - vFrom.X
	dirY := vTo.Y - vFrom.Y
	dirZ := vTo.Z - vFrom.Z

	if dirX*dirX+dirY*dirY < 1e-6 { // почти нулевой путь
		*vArrival = *vFrom
		fromCell := g.GetBaseCell(int(vFrom.X), int(vFrom.Y), int(vFrom.Z))
		toCell := g.GetBaseCell(int(vTo.X), int(vTo.Y), int(vTo.Z))
		return fromCell == toCell
	}

	// 2. Нормализуем вектор (NormalizeCube → просто Normalize)
	lenSq := dirX*dirX + dirY*dirY + dirZ*dirZ
	if lenSq <= 0 {
		return false // защита от деления на ноль, но это не должно происходить
	}
	invLen := 1 / math.Sqrt(lenSq)
	dirX *= invLen
	dirY *= invLen
	dirZ *= invLen

	// Маска направления: 0 – «<» (≤0), 1 – «>» (>0)
	var dirType [2]GeoDirEnum
	if dirX <= 0 {
		dirType[0] = GEO_DIR_N // просто пример, в реале нужен массив GeoDirMask
	}
	if dirY <= 0 {
		dirType[1] = GEO_DIR_NE
	} else {
		dirType[1] = GEO_DIR_E
	}

	// 3. Инициализация переменных цикла
	pos := *vFrom // текущая позиция
	cell := g.GetBaseCell(int(vFrom.X), int(vFrom.Y), int(vFrom.Z))
	if cell == nil {
		return false
	}
	distPassed := 0.0
	stepSize := math.Pow(2, -3) // 1/8 – как в оригинале (0.125)
	skipCheck := false          // v24 / v57 в исходнике

	for distTogo > distPassed {

		// --- расчёт следующего шага -----------------------------
		var stepX, stepY, stepZ float64
		if skipCheck || math.Abs(dirX*dirX+dirY*dirY) <= distTogo-distPassed {
			stepX = dirX
			stepY = dirY
			stepZ = 0 // пока не используется в расчётах пути
		} else {
			// уменьшаем шаг, если надо пройти меньше, чем до цели
			tmpX := dirX * stepSize
			tmpY := dirY * stepSize
			tmpZ := dirZ * stepSize
			if tmpX*tmpX+tmpY*tmpY < 0.5 {
				stepX = dirX
				stepY = dirY
				stepZ = dirZ
			} else {
				stepX, stepY, stepZ = tmpX, tmpY, tmpZ
			}
			skipCheck = true
		}

		// Если следующий шаг «превышает» расстояние до цели – ограничиваем его
		distToTargetSq := (pos.X-vTo.X)*(pos.X-vTo.X) + (pos.Y-vTo.Y)*(pos.Y-vTo.Y)
		if stepX*stepX+stepY*stepY > distToTargetSq {
			stepX = vTo.X - pos.X
			stepY = vTo.Y - pos.Y
			stepZ = vTo.Z - pos.Z
		}

		nextPos := FVector{
			X: pos.X + stepX,
			Y: pos.Y + stepY,
			Z: pos.Z + stepZ, // в оригинале Z обновляется через высоту ячейки
		}

		// --- поиск следующей ячейки -----------------------------
		newCell := g.GetBaseCell(int(nextPos.X), int(nextPos.Y), int(nextPos.Z))
		if newCell == nil {
			// ошибка: не нашли ячейку – считаем, что движение невозможно дальше
			logError("error in movestraight3! [2023]")
			break
		}

		// 4. Обновляем высоту ячейки (по маске GeoHeightMask1)
		heightInCell := int((newCell.m_data >> 1) & uint16(GeoHeightMask1>>1))
		if nextPos.Z > float64(heightInCell) {
			nextPos.Z = float64(heightInCell)
		}

		// 5. Обновляем переменные цикла
		distPassed += math.Sqrt(stepX*stepX + stepY*stepY + stepZ*stepZ)
		pos = nextPos
		cell = newCell

		// 6. Ограничиваем высоту по nSurfaceHeight (см. оригинал)
		if heightInCell > nSurfaceHeight {
			if int(nextPos.Z) < nSurfaceHeight {
				break // достигнут предел «покрытия» поверхности
			}
			nextPos.Z = float64(nSurfaceHeight - 1)
		} else if nextPos.Z > float64(heightInCell) {
			nextPos.Z = float64(heightInCell)
		}

		if distTogo <= distPassed {
			break // достигли требуемого расстояния
		}
	}

	// 7. Финальный результат
	if !skipCheck { // v56 в оригинале – «первый проход» (нет перехода через границу)
		*moreTogo = true
	}
	*vArrival = pos

	return true
}


func (g *CGeoData) CanSee(pFrom, pTo *CWorldObject, checkCollision bool) bool {
	if pFrom == nil || pTo == nil {
		return false
	}


	zoneID := int(uintptr(unsafe.Pointer(&pFrom))) //pFrom.GetInZoneID() // nefr говорит это reflectId  у него
	toPos := pTo.M_RelPos // pTo.GetPos()
	fromPos := pFrom.M_RelPos //pFrom.GetPos()

	return g.CanSee2(&fromPos, &toPos, zoneID, checkCollision)
}

func (g *CGeoData) CanSee2(
	vFrom, vTo *FVector,
	nInstantZoneID int,
	bCheckCollision bool,
) bool {
	out := &FVector{}
	v8 := g.SingleLineCheckForProjectile(vFrom, vTo)

	if v8 && bCheckCollision {
		// Обнуляем out, хотя в Go это избыточно
		*out = FVector{}

		if gWorldPlaneCollision.CheckCollision(vFrom, vTo, 0.0, out, nInstantZoneID, 0) &&
			gWorldPlaneCollision.CheckCollision(vTo, vFrom, 0.0, out, nInstantZoneID, 0) {
			return false
		}
	}
	return v8
}


func (g *CGeoData) SingleLineCheckWithFlyTopDown(vFrom, vTo *FVector) bool {
	// Разница координат
	dx := vTo.X - vFrom.X
	dy := vTo.Y - vFrom.Y
	dz := vTo.Z - vFrom.Z
	initialZ := vFrom.Z

	// Проверка направления по Z
	dirZFlag := 0
	if vTo.Z < vFrom.Z {
		dirZFlag = 1
	}
	dirZFlag += 4

	// Если точки почти совпадают — проход свободен
	if dx*dx+dy*dy+dz*dz < 0.5 {
		return true
	}

	// Преобразование в координаты сетки
	startXCell := (int(vFrom.X) + 327680) >> 4
	startYCell := (int(vFrom.Y) + 262144) >> 4
	endXCell := (int(vTo.X) + 327680) >> 4
	endYCell := (int(vTo.Y) + 262144) >> 4

	// Определение направления движения по X и Y
	stepX := 1
	if endXCell <= startXCell {
		stepX = -1
		endXCell = startXCell
	}

	stepY := 1
	if endYCell <= startYCell {
		stepY = -1
		endYCell = startYCell
	}

	// Разница по осям
	deltaX := Abs(endXCell - startXCell)
	deltaY := Abs(endYCell - startYCell)

	// Направления движения
	dirX := stepX <= 0
	dirY := GeoDirEnum(btoi(stepY > 0))

	// Переменные для Bresenham
	errX := 0
	errY := 0
	stepErrX := 0
	stepErrY := 0

	if deltaX < deltaY {
		endXCell = endYCell
		stepErrY = 2 * deltaX
		errX = 2*deltaX - deltaY
		stepErrX = deltaX - deltaY
		startXCell = startYCell
	} else {
		stepErrX = 2 * deltaY
		errX = 2*deltaY - deltaX
		stepErrY = deltaY - deltaX
	}

	errAdjust := 2 * stepErrY

	// Высота изменения Z на каждый шаг
	if startXCell < endXCell {
		dz = dz / float64(endXCell-startXCell)
	}
	absDz := math.Abs(dz)

	// Получаем стартовую ячейку
	currentCell := g.GetBaseCell(uint32(vFrom.X), uint32(vFrom.Y), int(vFrom.Z))
	currentZ := initialZ

	// Основной цикл по X/Y
	if startXCell != endXCell {
		if startXCell < endXCell {
			// Параметры для движения
			gridX := 16*startXCell - 327672
			gridY := 16*startYCell - 262136
			nextGridX := gridX + 16*stepX
			nextGridY := gridY + 16*stepY
			xStepWorld := 16 * stepX
			yStepWorld := 16 * stepY
			xCell := startXCell
			yCell := startYCell
			errorTerm := errX
			errorDelta := stepErrX
			errorAdjust := errAdjust
			errorStep := stepErrY

			for {
				// Проверка пролета над землёй
				if absDz > 0 {
					if uint64(dirZFlag-4) > 1 {
						return false
					}
					groundZ := g.FindGround(uint(gridX), uint(gridY), uint(int(currentZ)+int(absDz)))

					if int(groundZ) > int(currentZ)+int(absDz) || int(currentZ) < int(groundZ) {
						return false
					}
				}

				if errorTerm > 0 {
					// Движение по диагонали
					nextCell := g.GetNextHorizontalCellOnSight(currentCell, GeoDirEnum(btoi(dirX)), nextGridX, gridY, int(currentZ))
					if nextCell == nil {
						return false
					}
					nextCell2 := g.GetNextHorizontalCellOnSight(nextCell, dirY, nextGridX, nextGridY, int(currentZ))
					if nextCell2 == nil {
						// Попробуем в обратном порядке
						altCell := g.GetNextHorizontalCellOnSight(currentCell, dirY, gridX, nextGridY, int(currentZ))
						if altCell == nil {
							return false
						}
						nextCell2 = g.GetNextHorizontalCellOnSight(altCell, GeoDirEnum(btoi(dirX)), nextGridX, nextGridY, int(currentZ))
						if nextCell2 == nil {
							return false
						}
					}
					currentCell = nextCell2
					gridX += xStepWorld
					gridY += yStepWorld
					xCell += stepX
					yCell += stepY
					errorTerm += errorAdjust
				} else {
					// Движение по одной оси
					if deltaX <= deltaY {
						// Двигаемся по Y
						nextCell := g.GetNextHorizontalCellOnSight(currentCell, dirY, gridX, nextGridY, int(currentZ))
						if nextCell == nil {
							return false
						}
						currentCell = nextCell
						gridY += yStepWorld
						yCell += stepY
						errorTerm += stepErrY
					} else {
						// Двигаемся по X
						nextCell := g.GetNextHorizontalCellOnSight(currentCell, GeoDirEnum(btoi(dirX)), nextGridX, gridY, int(currentZ))
						if nextCell == nil {
							return false
						}
						currentCell = nextCell
						gridX += xStepWorld
						xCell += stepX
						errorTerm += stepErrX
					}
				}

				startXCell++
				currentZ += dz
				if startXCell >= endXCell {
					return true
				}
			}
		}
		return true
	}

	// Вертикальная проверка
	if absDz > 0 {
		newZ := int(currentZ) + int(absDz)
		if dirZFlag-4 <= 1 {
			groundZ := g.FindGround(uint(16*startXCell-327672), uint(16*startYCell-262136), uint(newZ))
			if int(groundZ) > newZ || int(currentZ) < int(groundZ) {
				return false
			}
		}
	}
	return true
}


func (g *CGeoData) GetNextHorizontalCellOnSight(
	pCurrentCell *CGeoCell,
	dirType GeoDirEnum,
	destX uint,   // x координата цели
	destY int,    // y координата цели
	nCurrentHeight int, // высота текущей ячейки (высота по Y)
) *CGeoCell {

	/* 1. Получаем сектор целевой точки */
	sector := g.GetSectorByCoord(uint32(destX), uint32(destY))
	if sector == nil {
		return nil
	}

	/* 2. Пытаемся «напрямую» найти следующую ячейку по направлению */
	result := g.GetNextCell2(pCurrentCell, (*CGeoSector)(unsafe.Pointer(&sector)), dirType,
		int(destX), destY, nCurrentHeight, Jumpable)
	if result != nil {
		return result
	}

	/* 3. Если прямой путь недоступен – берём базовую ячейку сектора */
	baseCell := g.GetBaseCellFromSector((*CGeoSector)(unsafe.Pointer(&sector)), uint32(destX), uint32(destY), nCurrentHeight)
	if baseCell == nil {
		return nil
	}

	// высота базовой ячейки (см. C++: (m_data >> 1) & GeoHeightMask_1 >> 1)
	cellHeight := int((baseCell.Data>>1)&(GeoHeightMask1>>1))
	if nCurrentHeight < cellHeight {
		return nil
	}

	/* 4. Находим координаты «противоположного» направления */
	var (
		oppDir GeoDirEnum
		v15    uint // X‑координата для GetNextCell (destX +/- 16)
		v16    int  // Y‑координата для GetNextCell
	)

	switch dirType {
	case GEO_DIR_EAST: /* ищем в восточном направлении → смотрим на запад */
		oppDir = GEO_DIR_WEST
		v15 = destX - 16
	case GEO_DIR_WEST:
		oppDir = GEO_DIR_EAST
		v15 = destX + 16
	case GEO_DIR_NORTH:
		oppDir = GEO_DIR_SOUTH
		v16 = destY + 16
	case GEO_DIR_SOUTH:
		oppDir = GEO_DIR_NORTH
		v16 = destY - 16
	}

	/* если не было явно задано, оставляем Y‑координату как destY */
	if v16 == 0 {
		v16 = destY
	}
	/* если X‑координата не была установлена (например, для север/юг), ставим destX */
	if v15 == 0 && dirType != GEO_DIR_EAST && dirType != GEO_DIR_WEST {
		v15 = destX
	}

	/* 5. Проверяем «противоположный» путь от базовой ячейки */
	next := g.GetNextCell(
		baseCell,
		oppDir,
		int(v15),   // X‑координата
		v16,        // Y‑координата
		cellHeight, // высота из базовой ячейки
		Jumpable)

	if next != pCurrentCell {
		return nil
	}

	/* 6. Всё ок – возвращаем базовую ячейку */
	return baseCell
}

func (this *CGeoData) FindGround(x, y, z uint) int64 {
	// Вычисление координат зоны
	zoneX := int(int32(x+327680)>>15) + 10
	zoneY := int(int32(y+0x40000)>>15) + 10

	zone := this.GetZone(zoneX, zoneY)
	if zone == nil {
		return 0
	}

	// Получение сектора по координатам
	sector := &zone.Data.Sectors[uint8(x>>7)][uint8(y>>7)] // &0x1F = %32, если массив 32x32

	if sector == nil {
		return 0
	}

	var m_data int32

	// Проверка флага
	if (sector.BooleanFlag & 1) != 0 {
		m_data = int32(sector.DefaultCell.Data)
	} else {
		baseCell := this.GetBaseCellFromSector(sector, uint32(x), uint32(y), int(z))
		if baseCell == nil {
			return 0
		}
		m_data = int32(baseCell.Data)
	}

	// Возврат значения
	return int64(int64((GeoHeightMask1 >> 1)) & int64(uint(m_data)>>1))
}



func (c *CWorldPlaneCollision) ConvertWorldToPlaneSection(
	x0, x1, y0, y1 int,
	idxMin, idxMax [2]int,
	idyMin, idyMax *int) {
	/* … реализовать логически аналогично C++‑функции */
}

func (b *CAxisAlgnBB) Intersect(start, end *FVector) bool { /* … */ return false }

func (c *CPlaneCollision) CheckCollision(
	start, end *FVector,
	radius float64,
	summonCheck bool) float64 {
	/* … реальная логика проверки пересечения */
	return 1.0
}

// ---------------------------------------------------------------------------
// 4. Перевод функции CheckCollision
// ---------------------------------------------------------------------------

func (c *CWorldPlaneCollision) CheckCollision(
	start, end *FVector,
	radius float64,
	out *FVector,
	inInstantZoneID uint32,
	summonCheck bool) bool {

	// 1. Увеличиваем счётчик для текущего потока
	threadIdx := getThreadIndex() // аналог *(int *)(v42 + 104228)
	key := c.m_CheckKey[threadIdx]
	c.m_CheckKey[threadIdx] = key + 1

	// 2. Перевод мировых координат в индексы ячеек (плоскости)
	var idxMin, idxMax [2]int
	var idyMin, idyMax int
	c.ConvertWorldToPlaneSection(
		int(start.X), int(end.X),
		int(start.Y), int(end.Y),
		idxMin, idxMax,
		&id yMin, &id yMax)

	if idxMin[0] > idxMax[0] || idyMin > idyMax {
		return false // путь не попадает в область карты
	}

	// 3. Перебор ячеек плоскости (первый проход – поиск коллизий)
	for x := idxMin[0]; x <= idxMax[0]; x++ {
		for y := idyMin; y <= idyMax; y++ {

			link := c.getCollisionLink(x, y) // аналог v16
			for link != nil {
				if link.Plane != nil &&
					CAxisAlgnBB(link.Next).Intersect(start, end) {
					// нашли пересечение – сохраняем и выходим
					*out = *start
					return true
				}
				link = link.Next
			}
		}
	}

	// 4. Если в зоне с instant‑ID надо проверить ещё коллизии
	if inInstantZoneID != 0 {
		inz := getInstantZone(inInstantZoneID) // аналог Shuttle + CSmartID::GetShuttle
		if inz == nil || !inz.IsActive() {
			return false
		}
	}

	// 5. Второй проход – перебор всех ячеек, в которых может быть столкновение
	var bestDist float64 = math.Inf(1)
	for x := idxMin[0]; x <= idxMax[0]; x++ {
		for y := idyMin; y <= idyMax; y++ {

			link := c.getCollisionLink(x, y) // аналог v23
			for link != nil && link.Plane != nil {

				// 5.1 Проверяем включённость коллизии (InstantZone/normal)
				if inz != nil {
					if !inz.GetCollisionCheck(link.Plane) {
						link = link.Next
						continue
					}
				} else if link.Plane.m_InstantZoneKey.typeId != 0 ||
					!link.Plane.m_bEnabled {
					link = link.Next
					continue
				}

				// 5.2 Проверяем, не проверяли ли уже эту коллизию в текущем потоке
				if c.m_CheckField[threadIdx][link.Plane.m_nIndex] == key {
					link = link.Next
					continue
				}
				c.m_CheckField[threadIdx][link.Plane.m_nIndex] = key

				// 5.3 Фактическая проверка пересечения
				dist := link.Plane.CheckCollision(start, end, radius, summonCheck)
				if dist < bestDist {
					bestDist = dist
				}
				link = link.Next
			}
		}
	}

	// 6. Если не нашли – возвращаем false
	if math.IsInf(bestDist, 1) {
		return false
	}

	// 7. Вычисляем точку пересечения (полученную из dist)
	out.X = start.X + (end.X-start.X)*bestDist
	out.Y = start.Y + (end.Y-start.Y)*bestDist
	out.Z = start.Z + (end.Z-start.Z)*bestDist

	return true
}



////////////////////////////////////////////////////////////
func LODWORD(x float64) uint32 {
	return uint32(math.Float64bits(x) & 0xFFFFFFFF)
}

func HIDWORD(x float64) uint32 {
	return uint32(math.Float64bits(x) >> 32)
}

func SLODWORD(x float64) int32 {
	return int32(math.Float64bits(x) & 0xFFFFFFFF)
}

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs[T number](x T) T {
 return T(math.Float64frombits(math.Float64bits(float64(x)) &^ (1 << 63)))
}
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}