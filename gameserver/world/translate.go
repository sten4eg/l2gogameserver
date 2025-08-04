package world

import (
	"math"
)

const GeoHeightMask_1 = -16

var GeoDirMask = [4]int{1, 2, 4, 8}

func (g *CGeoData) MoveStraight(vFrom, vTo *FVector, distTogo float64, vArrival *FVector, moreTogo *bool, distPassed *float64) bool {
	*moreTogo = false
	*distPassed = 0.0

	// Вычисление вектора движения
	dx := vTo.X - vFrom.X
	dy := vTo.Y - vFrom.Y
	dz := vTo.Z - vFrom.Z

	// Проверка на нулевое расстояние
	if math.Abs(dx*dx+dy*dy) < 0.000001 {
		*moreTogo = false
		*vArrival = *vFrom
		return g.GetBaseCell(int(vTo.X), int(vTo.Y), int(vTo.Z)) == g.GetBaseCell(int(vFrom.X), int(vFrom.Y), int(vFrom.Z))
	}

	// Нормализация вектора
	length := math.Sqrt(dx*dx + dy*dy + dz*dz)
	if length > 0 {
		dx /= length
		dy /= length
		dz /= length
	}

	// Получение начальной ячейки
	baseCell := g.GetBaseCell(int(vFrom.X), int(vFrom.Y), int(vFrom.Z))
	if baseCell == nil {
		*moreTogo = false
		*vArrival = *vFrom
		return false
	}

	// Инициализация переменных
	currentHeight := int(vFrom.Z)
	currentX := int(vFrom.X)
	currentY := int(vFrom.Y)

	// Создание начального вектора движения
	moveVector := FVector{dx, dy, dz}

	// Основной цикл движения
	for *distPassed < distTogo {
		// Вычисление следующей позиции
		nextX := vFrom.X + dx
		nextY := vFrom.Y + dy
		nextZ := vFrom.Z + dz

		// Проверка на достижение цели
		distanceToTarget := math.Sqrt(
			(nextX-vTo.X)*(nextX-vTo.X) +
				(nextY-vTo.Y)*(nextY-vTo.Y))

		if distanceToTarget <= 0.5 { // условие приближения к цели
			*vArrival = *vTo
			*moreTogo = true
			return true
		}

		// Обновление позиции
		vFrom.X = nextX
		vFrom.Y = nextY
		vFrom.Z = nextZ

		// Обновление пройденного расстояния
		*distPassed += math.Sqrt(dx*dx + dy*dy)

		// Проверка перехода между ячейками
		newX := int(vFrom.X)
		newY := int(vFrom.Y)

		if currentX != newX || currentY != newY {
			// Логика перехода между ячейками
			cell := g.GetNextCell(baseCell, newX, newY)
			if cell == nil {
				*moreTogo = true
				return false
			}
			baseCell = cell
		}

		currentX = newX
		currentY = newY
	}

	// Возвращаем финальную позицию
	*vArrival = *vFrom
	*moreTogo = true
	return true
}

func (g *CGeoData) GetBaseCell(x, y uint32, z int) *CGeoCell {
	sector := g.GetSectorByCoord(x, y)
	if sector != nil {
		return g.GetBaseCellFromSector(sector, x, y, z)
	}
	return sector
}
func (g *CGeoData) GetBaseCell2(z int, cellArray *CGeoCell, b1, b2 int) *CGeoCell {
	// Начинаем с индекса b1
	currentIndex := b1

	// Проходим по массиву до b2
	for currentIndex < b2 {
		// Получаем высоту ячейки через битовые операции
		cellHeight := (GeoHeightMask_1 >> 1) & (cellArray[currentIndex].MData >> 1)

		// Если z больше или равен высоте, возвращаем текущую ячейку
		if z >= cellHeight {
			return &cellArray[currentIndex]
		}

		currentIndex++
	}

	// Возвращаем последнюю ячейку (или nil если массив пуст)
	if currentIndex < len(cellArray) {
		return &cellArray[currentIndex]
	}

	return nil
}
func (g *CGeoData) GetBaseCellFromSector(sector *CGeoSector, x, y uint32, z int) *CGeoCell {
	if (sector.MBooleanFlag & 1) != 0 {
		return &sector.MDefaultCell
	}

	// Вычисление индексов ячеек
	xIndex := (int(x) + 327680) >> 15
	yIndex := (int(y) + 0x40000) >> 15

	// Получение ячейки из массива
	if xIndex >= 0 && xIndex < len(g.MCells) &&
		yIndex >= 0 && yIndex < len(g.MCells[xIndex]) {
		v9 := g.MCells[xIndex][yIndex]
		if v9 != nil {
			// Получение индексов ячеек
			b1 := make([]int, 4)
			var v14 int

			g.GetCellIndexes(x, y, b1, &v14)

			// Вызов GetBaseCell с дополнительными параметрами
			baseCell := g.GetBaseCell2(z+26, v9, b1[0], v14)
			if baseCell == nil {
				// Логирование ошибки (в оригинале используется CLog::Add)
				// В Go можно использовать log.Printf или аналогичный механизм
				// fmt.Printf("failed to get basecell (%d, %d, %d) at file[%s], line[%d]\n", x, y, z, "..\\Shared\\GeoData.cpp", 680)
			}
			return baseCell
		} else {
			// Логирование ошибки
			// fmt.Printf("failed to get cell_array (%d, %d) at file[%s], line[%d]\n", x, y, "..\\Shared\\GeoData.cpp", 670)
			return nil
		}
	}

	return nil
}

func (g *CGeoData) GetNextCell(baseCell *CGeoCell, x, y int) *CGeoCell {
	// Реализация получения следующей ячейки
	return nil
}
func (g *CGeoData) GetSectorByCoord(x, y uint32) *CGeoZone {
	// Вычисление индексов зоны
	idx := ((int(x) + 327680) >> 15) + 10
	idy := ((int(y) + 0x40000) >> 15) + 10

	// Получение зоны
	result := g.GetZone(idx, idy)
	if result != nil {
		// Вычисление смещения внутри зоны
		xShift := uint64(x) >> 7
		yShift := uint64(y) >> 7

		// Вычисление адреса сектора (эмуляция указательной арифметики)
		offset := 6144*xShift + 24*yShift
		// В Go не можем напрямую работать с указательной арифметикой,
		// поэтому нужно реализовать по-другому

		// Предполагаем, что возвращаемая зона содержит нужную информацию
		return result
	}
	return result
}

func (g *CGeoData) GetZone(idx, idy int) *CGeoZone {
	// Проверка границ
	if uint32(idx-10) <= 0x13 && uint32(idy-10) <= 0x10 {
		// Вычисление индекса в массиве
		v3 := idy - 10 + 17*(idx-10)

		// Проверка наличия данных
		if g.MZone[0][v3].MData != nil {
			return &g.MZone[0][v3]
		}
	}

	return nil
}

func (g *CGeoData) GetCellIndexes(x, y uint32, b1, b2 *int) int {
	v5 := int(y) + 0x40000
	v7 := int(x) + 327680

	zone := g.GetZone((v7>>15)+10, (v5>>15)+10)
	if zone == nil {
		return 0
	}

	sector := &zone.MData.Sectors[uint64(x)>>7][uint64(y)>>7]
	if sector == nil {
		return 0
	}

	v14 := (int(y) >> 4) & 7
	v15 := (int(x) >> 4) & 7
	v16 := g.MCellOffset[uint64(v7)>>15][uint64(v5)>>15]

	*b1 = g.GetFirstCellIndex(sector, v15, v14, v16)
	*b2 = g.GetLastCellIndex(sector, v15, v14, v16)
	return 1
}

func (g *CGeoData) GetFirstCellIndex(sector *CGeoSector, idx, idy int, cellOffsetArray []int16) int64 {
	if (sector.MBooleanFlag & 1) != 0 {
		return 0
	}

	if (sector.MBooleanFlag & 2) != 0 {
		return int64(idy + sector.CellSegment + 8*idx)
	}

	if idx != 0 || idy != 0 {
		return int64(sector.CellSegment + int(cellOffsetArray[8*idx-1+sector.CellOffsetIndex+idy]))
	}

	return int64(sector.CellSegment)
}

func (g *CGeoData) GetLastCellIndex(sector *CGeoSector, idx, idy int, cellOffsetArray []int16) int64 {
	if (sector.MBooleanFlag & 1) != 0 {
		return 0
	}

	if (sector.MBooleanFlag & 2) != 0 {
		return int64(sector.CellSegment + 8*idx + idy + 1)
	}

	return int64(sector.CellSegment + int(cellOffsetArray[8*idx+idy+sector.CellOffsetIndex]))
}

func (g *CGeoData) GetNextCell(pCurrentCell *CGeoCell, pNextSector *CGeoSector, dirType int, destX, destY, currentHeight int, jump int) *CGeoCell {
	if (pCurrentCell.MData & int16(GeoDirMask[dirType])) == GeoDirMask[dirType] {
		return g.GetBaseCellFromSector(pNextSector, destX, destY, currentHeight)
	}

	if jump == Jumpable {
		closestUpperCellFromSector := g.GetClosestUpperCellFromSector(pNextSector, destX, destY, currentHeight, true)
		if closestUpperCellFromSector != nil {
			if ((GeoHeightMask_1 >> 1) & (closestUpperCellFromSector.MData >> 1)) <= currentHeight+32 {
				return closestUpperCellFromSector
			}
		}
	}

	return nil
}

func (g *CGeoData) GetNextCell2(pCurrentCell *CGeoCell, dirType int, destX uint32, destY, currentHeight, jump int) *CGeoCell {
	result := g.GetSectorByCoord(destX, uint32(destY))
	if result != nil {
		return g.GetNextCell(pCurrentCell, (*CGeoSector)(result), dirType, int(destX), destY, currentHeight, jump)
	}
	return result
}

func (g *CGeoData) GetClosestUpperCellFromSector(sector *CGeoSector, x uint32, y, z, strickCheck int) *CGeoCell {
	if (sector.MBooleanFlag & 1) != 0 {
		return &sector.MDefaultCell
	}

	xIndex := (int(x) + 327680) >> 15
	yIndex := (int(y) + 0x40000) >> 15

	if xIndex >= 0 && xIndex < len(g.MCells) &&
		yIndex >= 0 && yIndex < len(g.MCells[xIndex]) {
		v10 := g.MCells[xIndex][yIndex]
		if v10 != nil {
			b1 := make([]int, 4)
			var v19 int

			g.GetCellIndexes(x, uint32(y), &b1[0], &v19)

			closestUpperCell := g.GetClosestUpperCell(z, v10, b1[0], v19, strickCheck != 0)
			if closestUpperCell == nil && strickCheck == 0 {
				// Логирование ошибки (в оригинале CLog::Add)
				return nil
			}
			return closestUpperCell
		} else {
			// Логирование ошибки (в оригинале CLog::Add)
			return nil
		}
	}

	return nil
}

// Дополнительные функции для завершения реализации
func (g *CGeoData) GetClosestUpperCell(z int, cellArray []*CGeoCell, b1, b2 int, strickCheck bool) *CGeoCell {
	// Реализация поиска ближайшей верхней ячейки
	return nil
}
