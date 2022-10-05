package models

import (
	"github.com/puzpuzpuz/xsync"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/interfaces"
	"math"
	"strconv"
)

const MAP_MIN_X = (config.GeoFirstX - 20) << 15
const MAP_MAX_X = ((config.GeoLastX - 20 + 1) << 15) - 1
const MAP_MIN_Y = (config.GeoFirstY - 18 + 1) << 15
const MAP_MAX_Y = ((config.GeoLastY - 18 + 1) << 15) - 1
const MAP_MIN_Z = -16384
const MAP_MAX_Z = 16384

const WORLD_SIZE_X = config.GeoLastX - config.GeoFirstX + 1
const WORLD_SIZE_Y = config.GeoLastY - config.GeoFirstY + 1

const SHIFT_BY = config.SHIFT_BY
const SHIFT_BY_Z = config.SHIFT_BY_Z

var OFFSET_X = math.Abs(MAP_MIN_X >> SHIFT_BY)
var OFFSET_Y = math.Abs(MAP_MIN_Y >> SHIFT_BY)
var OFFSET_Z = math.Abs(MAP_MIN_Z >> SHIFT_BY_Z)

var REGIONS_X = int32((MAP_MAX_X >> SHIFT_BY) + OFFSET_X)
var REGIONS_Y = int32((MAP_MAX_Y >> SHIFT_BY) + OFFSET_Y)
var REGIONS_Z = int32((MAP_MAX_Z >> SHIFT_BY_Z) + OFFSET_Z)

type WorldRegion struct {
	TileX         int32
	TileY         int32
	TileZ         int32
	CharsInRegion *xsync.MapOf[interfaces.CharacterI]
	NpcInRegion   *xsync.MapOf[interfaces.Npcer]
	ItemsInRegion *xsync.MapOf[interfaces.MyItemInterface]
}

func NewWorldRegion(x, y, z int32) *WorldRegion {
	var newRegion WorldRegion
	newRegion.TileX = x
	newRegion.TileY = y
	newRegion.TileZ = z
	newRegion.CharsInRegion = xsync.NewMapOf[interfaces.CharacterI]()
	newRegion.NpcInRegion = xsync.NewMapOf[interfaces.Npcer]()
	newRegion.ItemsInRegion = xsync.NewMapOf[interfaces.MyItemInterface]()
	return &newRegion
}

func (w *WorldRegion) GetX() int32 {
	return w.TileX
}
func (w *WorldRegion) GetY() int32 {
	return w.TileY
}
func (w *WorldRegion) GetZ() int32 {
	return w.TileZ
}
func (w *WorldRegion) AddVisibleChar(character interfaces.CharacterI) {
	key := strconv.FormatInt(int64(character.GetObjectId()), 10)
	w.CharsInRegion.Store(key, character)
}
func (w *WorldRegion) DeleteVisibleChar(character interfaces.CharacterI) {
	key := strconv.FormatInt(int64(character.GetObjectId()), 10)
	w.CharsInRegion.Delete(key)
}

func (w *WorldRegion) AddVisibleNpc(npc Npc) {
	key := strconv.FormatInt(int64(npc.GetObjectId()), 10)
	w.NpcInRegion.Store(key, &npc)
}

func (w *WorldRegion) GetNeighbors() []interfaces.WorldRegioner {
	return GetNeighbors(int(w.TileX), int(w.TileY), int(w.TileZ), 1, 1)
}

func (w *WorldRegion) GetChar(objectId int32) (interfaces.CharacterI, bool) {
	key := strconv.FormatInt(int64(objectId), 10)
	return w.CharsInRegion.Load(key)
}

func (w *WorldRegion) GetCharsInRegion() []interfaces.CharacterI {
	result := make([]interfaces.CharacterI, 0, 64)
	w.CharsInRegion.Range(func(key string, value interfaces.CharacterI) bool {
		result = append(result, value.(*Character))
		return true
	})

	return result
}

func (w *WorldRegion) GetNpc(objectId int32) (interfaces.Npcer, bool) {
	key := strconv.FormatInt(int64(objectId), 10)
	return w.NpcInRegion.Load(key)
}

func (w *WorldRegion) GetNpcInRegion() []interfaces.Npcer {
	result := make([]interfaces.Npcer, 0, 64)
	w.NpcInRegion.Range(func(key string, value interfaces.Npcer) bool {
		result = append(result, value.(*Npc))
		return true
	})

	return result
}

func (w *WorldRegion) AddVisibleItems(item interfaces.MyItemInterface) {
	key := strconv.FormatInt(int64(item.GetObjectId()), 10)
	w.ItemsInRegion.Store(key, item)
}
func (w *WorldRegion) GetItem(objectId int32) (interfaces.MyItemInterface, bool) {
	key := strconv.FormatInt(int64(objectId), 10)
	return w.ItemsInRegion.Load(key)
}
func (w *WorldRegion) GetItemsInRegion() []interfaces.MyItemInterface {
	result := make([]interfaces.MyItemInterface, 0, 64)
	w.ItemsInRegion.Range(func(key string, value interfaces.MyItemInterface) bool {
		result = append(result, value)
		return true
	})
	return result
}
func (w *WorldRegion) DeleteVisibleItem(item interfaces.MyItemInterface) {
	key := strconv.FormatInt(int64(item.GetObjectId()), 10)
	w.ItemsInRegion.Delete(key)
}

func Contains(regions []interfaces.WorldRegioner, region interfaces.WorldRegioner) bool {
	for index := range regions {
		if regions[index] == region {
			return true
		}
	}
	return false
}

// GetAroundPlayer возвращает персонажей в текущем и всех соседних регионах
func GetAroundPlayer(character interfaces.Positionable) []interfaces.CharacterI {
	currentRegion := character.GetCurrentRegion()
	if nil == currentRegion {
		return nil
	}
	result := make([]interfaces.CharacterI, 0, 64)

	neighbors := currentRegion.GetNeighbors()
	for index := range neighbors {
		result = append(result, neighbors[index].GetCharsInRegion()...)
	}
	return result
}

func (w *WorldRegion) GetCharacterInRegions(objectId int32) interfaces.CharacterI {
	neighbors := w.GetNeighbors()
	for index := range neighbors {
		char, ok := neighbors[index].GetChar(objectId)
		if ok {
			return char
		}
	}
	return nil
}

// GetAroundPlayersInRadius найди пользователей в радиусе radius, с максимальной разницой по Z = height
func GetAroundPlayersInRadius(myCharacter interfaces.CharacterI, radius int32, height float64) []interfaces.CharacterI {
	currentRegion := myCharacter.GetCurrentRegion()
	if nil == currentRegion {
		return nil
	}

	result := make([]interfaces.CharacterI, 0, 64)

	sqrad := radius * radius

	for _, v := range currentRegion.GetNeighbors() {
		charactersInRegion := v.GetCharsInRegion()
		for charInRegionIndex := range charactersInRegion {

			if charactersInRegion[charInRegionIndex].GetObjectId() == myCharacter.GetObjectId() {
				continue
			}
			dx := math.Abs(float64(charactersInRegion[charInRegionIndex].GetX() - myCharacter.GetX()))
			if dx > float64(radius) {
				continue
			}

			dy := math.Abs(float64(charactersInRegion[charInRegionIndex].GetY() - myCharacter.GetY()))
			if dy > float64(radius) {
				continue
			}

			dz := math.Abs(float64(charactersInRegion[charInRegionIndex].GetZ() - myCharacter.GetZ()))
			if dz > height {
				continue
			}

			if dx*dx+dy*dy > float64(sqrad) {
				continue
			}

			result = append(result, charactersInRegion[charInRegionIndex])
		}
	}
	return result
}

func validX(x int32) int32 {
	if x < 0 {
		x = 0
	} else if x > REGIONS_X {
		x = REGIONS_X
	}
	return x
}
func validY(y int32) int32 {
	if y < 0 {
		y = 0
	} else if y > REGIONS_Y {
		y = REGIONS_Y
	}
	return y
}
func validZ(z int32) int32 {
	if z < 0 {
		z = 0
	} else if z > REGIONS_Z {
		z = REGIONS_Z
	}
	return z
}

func regionX(x int32) int32 {
	return (x >> SHIFT_BY) + int32(OFFSET_X)
}
func regionY(y int32) int32 {
	return (y >> SHIFT_BY) + int32(OFFSET_Y)
}
func regionZ(z int32) int32 {
	return (z >> SHIFT_BY_Z) + int32(OFFSET_Z)
}

//todo
//private static WorldRegion getRegion(final int x, final int y, final int z) {

func GetAroundPlayersObjIdInRadius(c *Character, radius int32) []int32 {
	currentRegion := c.CurrentRegion
	if nil == currentRegion {
		return nil
	}
	result := make([]int32, 0, 64)

	sqrad := radius * radius

	for _, v := range currentRegion.GetNeighbors() {
		charInRegion := v.GetCharsInRegion()
		for _, vv := range charInRegion {
			char, ok := vv.(*Character)
			if !ok {
				continue
			}
			if char.ObjectId == c.ObjectId {
				continue
			}
			dx := math.Abs(float64(char.Coordinates.X - c.Coordinates.X))
			if dx > float64(radius) {
				continue
			}

			dy := math.Abs(float64(char.Coordinates.Y - c.Coordinates.Y))
			if dy > float64(radius) {
				continue
			}

			if dx*dx+dy*dy > float64(sqrad) {
				continue
			}

			result = append(result, char.ObjectId)

		}
	}
	return result
}
func GetAroundPlayerObjId(c *Character) []int32 {
	currentRegion := c.GetCurrentRegion()
	if nil == currentRegion {
		return nil
	}
	result := make([]int32, 0, 64)

	for _, v := range currentRegion.GetNeighbors() {
		for _, vv := range v.GetCharsInRegion() {
			result = append(result, vv.GetObjectId())
		}
	}
	return result
}
