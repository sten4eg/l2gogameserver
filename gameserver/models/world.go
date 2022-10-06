package models

import (
	"l2gogameserver/config"
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

const WORLD_SIZE_X = config.GeoLastX - config.GeoFirstX + 1
const WORLD_SIZE_Y = config.GeoLastY - config.GeoFirstY + 1

var MapMinX = (TileXMin - 20) << 15
var MapMaxX = ((TileXMax - 19) << 15) - 1

var MapMinY = (TileYMin - 18) << 15
var MapMaxY = ((TileYMax - 17) << 15) - 1

const MAP_MIN_Z = -16384
const MAP_MAX_Z = 16384

const SHIFT_BY = config.SHIFT_BY
const SHIFT_BY_Z = config.SHIFT_BY_Z

var OffsetX = math.Abs(float64(MapMinX >> SHIFT_BY))
var OffsetY = math.Abs(float64(MapMinY >> SHIFT_BY))
var OffsetZ = math.Abs(float64(MAP_MIN_Z >> SHIFT_BY_Z))

var RegionsX = int32((float64(MapMaxX >> SHIFT_BY)) + OffsetX)
var RegionsY = int32((float64(MapMaxY >> SHIFT_BY)) + OffsetY)
var RegionsZ = int32((float64(MAP_MAX_Z >> SHIFT_BY_Z)) + OffsetZ)

var World [][][]*WorldRegion

func NewWorld() {
	World = make([][][]*WorldRegion, RegionsX+1)

	for i := 0; i <= int(RegionsX); i++ {
		wj := make([][]*WorldRegion, RegionsY+1)
		for j := 0; j <= int(RegionsY); j++ {
			wz := make([]*WorldRegion, RegionsZ+1)
			for z := 0; z < int(RegionsZ); z++ {
				wz[z] = nil
			}
			wj[j] = wz
		}
		World[i] = wj
	}

}

// GetNeighbors x,y,z - координаты региона
func GetNeighbors(regionX, regionY, regionZ int32) []interfaces.WorldRegioner {
	res := make([]interfaces.WorldRegioner, 0, 27)
	for x := validX(regionX - 1); x <= validX(regionX+1); x++ {
		for y := validY(regionY - 1); y <= validY(regionY+1); y++ {
			for z := validZ(regionZ - 1); z <= validZ(regionZ+1); z++ {
				res = append(res, getRegion(x, y, z))
			}
		}
	}
	return res
}

// GetRegion игровые координаты объекта
func GetRegion(x, y, z int32) *WorldRegion {
	return getRegion(validX(regionX(x)), validY(regionY(y)), validZ(regionZ(z)))
}

// getRegion x,y,z - координаты региона
func getRegion(x, y, z int32) *WorldRegion {
	if World[x][y][z] == nil {
		World[x][y][z] = NewWorldRegion(x, y, z)
	}
	return World[x][y][z]
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

func isNeighbour(x1, y1, z1, x2, y2, z2 int32) bool {
	return (x1 <= x2+1) && (x1 >= x2-1) && (y1 <= y2+1) && (y1 >= y2-1) && (z1 <= z2+1) && (z1 >= z2-1)
}
