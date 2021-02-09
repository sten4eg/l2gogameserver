package models

import (
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

var World [][]WorldRegion

/** Gracia border Flying objects not allowed to the east of it. */
var GraciaMaxX = -166168
var GraciaMaxZ = 6105
var GraciaMinZ = -895

/** Biteshift, defines number of regions note, shifting by 15 will result in regions corresponding to map tiles shifting by 12 divides one tile to 8x8 regions. */
var ShiftBy = 12

var TileSize = 32768

/** Map dimensions */
var TileXMin = 11
var TileYMin = 10
var TileXMax = 26
var TileYMax = 26
var TileZeroCoordX = 20
var TileZeroCoordY = 18
var MapMinX = (TileXMin - TileZeroCoordX) * TileSize
var MapMinY = (TileYMin - TileZeroCoordY) * TileSize

var MapMaxX = ((TileXMax - TileZeroCoordX) + 1) * TileSize
var MapMaxY = ((TileYMax - TileZeroCoordY) + 1) * TileSize

/** calculated offset used so top left region is 0,0 */
var OffsetX = math.Abs(float64(MapMinX >> ShiftBy))
var OffsetY = math.Abs(float64(MapMinY >> ShiftBy))

/** number of regions */
var RegionsX = int32((float64(MapMaxX >> ShiftBy)) + OffsetX)
var RegionsY = int32((float64(MapMaxY >> ShiftBy)) + OffsetY)

func NewWorld() {

	var i int32
	var j int32
	World = make([][]WorldRegion, RegionsX+1)
	for i = 0; i <= RegionsX; i++ {
		wj := make([]WorldRegion, RegionsY+1)
		for j = 0; j <= RegionsY; j++ {
			wj[j] = NewWorldRegion(i, j)
		}
		World[i] = wj
	}
	var x, y, a, b int32
	for x = 0; x <= RegionsX; x++ {
		for y = 0; y <= RegionsY; y++ {
			for a = -1; a <= 1; a++ {
				for b = -1; b <= 1; b++ {
					if validRegion(x+a, y+b) {
						World[x+a][y+b].AddSurrounding(&World[x][y])
					}
				}
			}
		}
	}
}

func validRegion(x, y int32) bool {
	return (x >= 0) && (x <= RegionsX) && (y >= 0) && (y <= RegionsY)
}

//func getVisibleObjects(region WorldRegion, radius int32) {
//	sqRadius := radius * radius
//
//for regi := range region.Sur {
//	//Todo если я то надо континью (не надо для самого себя высчитывать)
//	if sqRadius >  {
//		calculateDistance()
//	}
//}
//}

//func calculateDistance(ox, oy, oz, mx, my, mz int32, includeZAxis, squared bool) float64 {
//	var distance float64
//	if includeZAxis {
//		distance = math.Pow(float64(ox-mx), 2) + math.Pow(float64(oy-my), 2) + math.Pow(float64(oz-mz), 2)
//	} else {
//		distance = math.Pow(float64(ox-mx), 2) + math.Pow(float64(oy-my), 2)
//	}
//
//	if squared {
//		return distance
//	}
//
//	return math.Sqrt(distance)
//}

func GetRegion(x, y int32) *WorldRegion {
	return &World[(x>>ShiftBy)+int32(OffsetX)][(y>>ShiftBy)+int32(OffsetY)]
}
