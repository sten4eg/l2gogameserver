package qw

import "math"

var GeoHeightMask1 int16 = -16

var GeoFieldMask2 = [16]int16{
	32768, 16384, 8192, 4096,
	2048, 1024, 512, 256,
	128, 64, 32, 16,
	8, 4, 2, 1,
}

var GeoDirMask = [4]int16{1, 2, 4, 8}

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
				if abs(currX-nextX) == 1 {
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
				if abs(currX-nextX) != 1 || abs(currY-nextY) != 1 {
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
			if abs(currY-nextY) == 1 {
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
