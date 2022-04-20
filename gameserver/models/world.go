package models

import (
	"l2gogameserver/gameserver/interfaces"
	"math"
	"sync"
)

type BackwardToLocation struct {
	TargetX int32
	TargetY int32
	TargetZ int32
	OriginX int32
	OriginY int32
	OriginZ int32
}
type OnlineCharacters struct {
	Char map[int32]*Character
	Mu   sync.Mutex
}

var GraciaMaxX = -166168
var GraciaMaxZ = 6105
var GraciaMinZ = -895

var TileXMin = 11
var TileYMin = 10
var TileXMax = 26
var TileYMax = 26

var WorldSizeX = 26 - 11 + 1
var WorldSizeY = 26 - 10 + 1

var MapMinX = (TileXMin - 20) << 15
var MapMaxX = ((TileXMax - 19) << 15) - 1

var MapMinY = (TileYMin - 18) << 15
var MapMaxY = ((TileYMax - 17) << 15) - 1

var MapMinZ = -16384
var MapMaxZ = 16383

var ShiftBy = 12
var ShiftByForZ = 10

var OffsetX = math.Abs(float64(MapMinX >> ShiftBy))
var OffsetY = math.Abs(float64(MapMinY >> ShiftBy))
var OffsetZ = math.Abs(float64(MapMinZ >> ShiftByForZ))

var RegionsX = int32((float64(MapMaxX >> ShiftBy)) + OffsetX)
var RegionsY = int32((float64(MapMaxY >> ShiftBy)) + OffsetY)
var RegionsZ = int32((float64(MapMaxZ >> ShiftByForZ)) + OffsetZ)

var World [][][]*WorldRegion

func NewWorld() {
	World = make([][][]*WorldRegion, RegionsX+1)
	for i := 0; i <= int(RegionsX); i++ {
		wj := make([][]*WorldRegion, RegionsY+1)
		for j := 0; j <= int(RegionsY); j++ {
			wz := make([]*WorldRegion, 1)
			for z := 0; z < 1; z++ { //todo Z из конфига??
				wz[z] = NewWorldRegion(int32(i), int32(j), int32(z))
			}
			wj[j] = wz
		}
		World[i] = wj
	}
	qw := World
	_ = qw
}
func GetNeighbors(regX, regY, regZ, deepH, deepV int) []interfaces.WorldRegioner {
	neighbors := make([]*WorldRegion, 0, 9)
	deepH *= 2
	deepV *= 2
	var rx, ry, rz int

	for x := 0; x <= deepH; x++ {
		for y := 0; y <= deepH; y++ {
			for z := 0; z <= deepV; z++ {

				if x%2 == 0 {
					rx = regX + (-x / 2)
				} else {
					rx = regX + (x - x/2)
				}

				if y%2 == 0 {
					ry = regY + (-y / 2)
				} else {
					ry = regY + (y - y/2)
				}
				rz = 0
				if validRegion(rx, ry, rz) {
					if len(World[rx][ry]) > 1 {
						if z%2 == 0 {
							rz = regZ + (-z / 2)
						} else {
							rz = regZ + (z - z/2)
						}

						if !validRegion(rx, ry, rz) {
							continue
						}
					} else {
						z = deepV + 1
					}

					qw := rx
					qe := ry
					qq := rz
					_, _, _ = qw, qe, qq
					if World[rx][ry][rz] != nil {
						neighbors = append(neighbors, World[rx][ry][rz])
					}
				}
			}
		}
	}
	ret := make([]interfaces.WorldRegioner, len(neighbors))
	for i, d := range neighbors {
		ret[i] = d
	}

	return ret
}

func GetRegion(x, y, z int32) *WorldRegion {
	_x := int(x>>ShiftBy) + int(OffsetX)
	_y := int(y>>ShiftBy) + int(OffsetY)
	_z := 0
	if validRegion(_x, _y, _z) {
		if len(World[_x][_y]) > 1 {
			_z = int(z>>ShiftByForZ) + int(OffsetZ)
		}

		if World[_x][_y][_z] == nil {
			World[_x][_y][_z] = NewWorldRegion(int32(_x), int32(_y), int32(_z))
		}
		return World[_x][_y][_z]
	}
	return nil
}

func CalculateDistance(ox, oy, oz, mx, my, mz int32, includeZAxis, squared bool) float64 {
	var distance float64
	if includeZAxis {
		distance = math.Pow(float64(ox-mx), 2) + math.Pow(float64(oy-my), 2) + math.Pow(float64(oz-mz), 2)
	} else {
		distance = math.Pow(float64(ox-mx), 2) + math.Pow(float64(oy-my), 2)
	}

	if squared {
		return distance
	}

	return math.Sqrt(distance)
}

func validRegion(x, y, z int) bool {
	return x >= 0 && x < int(RegionsX) && y >= 0 && y < int(RegionsY) && z >= 0 && z < int(RegionsZ)
}
