package qw

import (
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
		fromCell := geo.GetBaseCell(int(from.X), int(from.Y), int(from.Z))
		toCell := geo.GetBaseCell(int(to.X), int(to.Y), int(to.Z))
		return fromCell == toCell
	}

	dir := FVector{X: dirX, Y: dirY}
	dir.Normalize2D()

	curr := *from
	currCell := geo.GetBaseCell(int(curr.X), int(curr.Y), int(curr.Z))
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
					height := (int(currCell.Data>>1) & (GeoHeightMask1 >> 1))
					nextCell := geo.GetNextCell(currCell, lowDir, destX, destY, height)
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
				height := (int(currCell.Data>>1) & (GeoHeightMask1 >> 1))
				cellA := geo.GetNextCell(currCell, lowDir, destX, destY, height)
				if cellA == nil {
					break
				}
				heightB := (int(cellA.Data>>1) & (GeoHeightMask1 >> 1))
				cellB := geo.GetNextCell(cellA, highDir, destX, 16*nextY-262136, heightB)
				if cellB == nil {
					break
				}
				if GeoDirMask[highDir]&currCell.Data != GeoDirMask[highDir] {
					break
				}
				cellC := geo.GetNextCell(currCell, highDir, 16*currX-327672, 16*nextY-262136, height)
				if cellC == nil {
					break
				}
				heightC := (int(cellC.Data>>1) & (GeoHeightMask1 >> 1))
				finalCell := geo.GetNextCell(cellC, lowDir, destX, 16*nextY-262136, heightC)
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
				height := (int(currCell.Data>>1) & (GeoHeightMask1 >> 1))
				nextCell := geo.GetNextCell(currCell, highDir, destX, destY, height)
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
	sector := geo.GetSectorByCoord(destX, destY)
	if sector == nil {
		return nil
	}
	return geo.getNextCellWithSector(currentCell, sector, dir, destX, destY, currentHeight, jump)
}

func (geo *CGeoData) getNextCellWithSector(
	currentCell *CGeoCell,
	nextSector *CGeoSector,
	dir GeoDirEnum,
	destX, destY int,
	currentHeight int,
	jump JumpType,
) *CGeoCell {
	if currentCell.Data&GeoDirMask[dir] == GeoDirMask[dir] {
		return geo.GetBaseCellFromSector(nextSector, destX, destY, currentHeight)
	}

	if jump == Jumpable {
		upper := geo.GetClosestUpperCellFromSector(nextSector, destX, destY, currentHeight, true)
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
		return geo.GetBaseCellFromSector(sector, x, y, z)
	}
	return nil
}
func (geo *CGeoData) GetBaseCell2(z int32, cellArray []CGeoCell, b1, b2 int) *CGeoCell {
	const GeoHeightMask1 int16 = -16 // 0xFFF0
	for i := b1; i < b2; i++ {
		cell := &cellArray[i]
		// (GeoHeightMask1 >> 1) = 0x7FF8
		height := int32((GeoHeightMask1 >> 1) & (cell.data >> 1))
		if z >= height {
			return cell
		}
	}
	// если ничего не найдено — возвращаем последний проверенный (или nil, по логике C)
	if b2 > b1 {
		return &cellArray[b2-1]
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

	if cellX >= len(geo.mCells) || cellY >= len(geo.mCells[cellX]) {
		log.Printf("failed to get cell_array (%d, %d)", x, y)
		return nil
	}

	cellArray := geo.mCells[cellX][cellY]
	if cellArray == nil {
		log.Printf("failed to get cell_array (%d, %d)", x, y)
		return nil
	}

	var b1 [4]int
	var b2 int

	geo.GetCellIndexes(x, y, b1[:], &b2)

	baseCell := geo.GetBaseCell2(z+26, cellArray, b1[0], b2)
	if baseCell == nil {
		log.Printf("failed to get basecell (%d, %d, %d)", x, y, z)
	}
	return baseCell
}

func (geo *CGeoData) GetSectorByCoord(x, y uint32) *CGeoZone {
	zoneX := int((x+327680)>>15) + 10
	zoneY := int((y+262144)>>15) + 10 // 0x40000 = 262144

	zone := geo.GetZone(zoneX, zoneY)
	if zone == nil {
		return nil
	}

	offset := 6144*int((x>>7)&0xFF) + 24*int((y>>7)&0xFF)
	if offset+24 > len(zone.Data) {
		return nil // безопасная проверка
	}

	// Нужно кастовать данные из zone.Data[offset] в *CGeoZone или *CGeoSector
	// Здесь предполагается, что данные лежат в []CGeoSector (если ты это знаешь точно — можно сделать безопаснее)
	sectorPtr := (*CGeoZone)(unsafe.Pointer(&zone.Data[offset]))
	return sectorPtr
}

func (geo *CGeoData) GetCellIndexes(x, y uint32, b1 *int, b2 *int) bool {
	v5 := int(y) + 0x40000 // signed
	v7 := int(x) + 327680  // signed

	zoneX := (v7 >> 15) + 10
	zoneY := (v5 >> 15) + 10

	zone := geo.GetZone(zoneX, zoneY)
	if zone == nil {
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

	offsets := geo.mCellOffset[v7>>15][v5>>15]

	*b1 = sector.GetFirstCellIndex(tileX, tileY, offsets)
	*b2 = sector.GetLastCellIndex(tileX, tileY, offsets)

	return true
}

func (g *CGeoData) GetZone(idx, idy int) *CGeoZone {
	if idx-10 >= 0 && idx-10 <= 0x13 && idy-10 >= 0 && idy-10 <= 0x10 {
		v := (idy - 10) + 17*(idx-10)
		if g.m_Zone[0][v].m_data != nil {
			return &g.m_Zone[0][v]
		}
	}
	return nil
}
