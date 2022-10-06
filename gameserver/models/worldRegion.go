package models

import (
	"github.com/alphadose/haxmap"
	"github.com/puzpuzpuz/xsync"
	"l2gogameserver/gameserver/interfaces"
	"math"
	"strconv"
	"time"
)

type WorldRegion struct {
	TileX         int32
	TileY         int32
	TileZ         int32
	CharsInRegion *xsync.MapOf[interfaces.CharacterI]
	NpcInRegion   *xsync.MapOf[interfaces.Npcer]
	ItemsInRegion *xsync.MapOf[interfaces.MyItemInterface]

	ItemsExpireTime *haxmap.Map[int32, int64]
}

func NewWorldRegion(x, y, z int32) *WorldRegion {
	var newRegion WorldRegion
	newRegion.TileX = x
	newRegion.TileY = y
	newRegion.TileZ = z
	newRegion.CharsInRegion = xsync.NewMapOf[interfaces.CharacterI]()
	newRegion.NpcInRegion = xsync.NewMapOf[interfaces.Npcer]()
	newRegion.ItemsInRegion = xsync.NewMapOf[interfaces.MyItemInterface]()
	newRegion.ItemsExpireTime = haxmap.New[int32, int64]()
	return &newRegion
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
	expireTime := time.Now().Add(time.Second * 10).Unix()

	w.ItemsInRegion.Store(key, item)
	w.ItemsExpireTime.Set(item.GetObjectId(), expireTime)
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
	w.ItemsExpireTime.Del(item.GetObjectId())
}

//func (w *WorldRegion) Get

func Contains(regions []interfaces.WorldRegioner, region interfaces.WorldRegioner) bool {
	for _, v := range regions {
		if v == region {
			return true
		}
	}
	return false
}
func GetAroundPlayer(c interfaces.Positionable) []interfaces.CharacterI {
	currentRegion := c.GetCurrentRegion()
	if nil == currentRegion {
		return nil
	}
	result := make([]interfaces.CharacterI, 0, 64)

	for _, v := range currentRegion.GetNeighbors() {
		result = append(result, v.GetCharsInRegion()...)
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

func (w *WorldRegion) GetCharacterInRegions(objectId int32) interfaces.CharacterI {
	for _, region := range w.GetNeighbors() {
		char, ok := region.GetChar(objectId)
		if ok {
			return char
		}
	}
	return nil
}

func GetAroundPlayersInRadius(c interfaces.CharacterI, radius int32) []*Character {
	currentRegion := c.GetCurrentRegion()
	if nil == currentRegion {
		return nil
	}
	result := make([]*Character, 0, 64)

	sqrad := radius * radius

	for _, v := range currentRegion.GetNeighbors() {
		charInRegion := v.GetCharsInRegion()
		for _, vv := range charInRegion {
			char, ok := vv.(*Character)
			if !ok {
				continue
			}
			if char.GetObjectId() == c.GetObjectId() {
				continue
			}
			dx := math.Abs(float64(char.Coordinates.X - c.GetX()))
			if dx > float64(radius) {
				continue
			}

			//toDO тут Y должен быть ? проверить на других сборках
			dy := math.Abs(float64(char.Coordinates.Y - c.GetY()))
			if dy > float64(radius) {
				continue
			}

			if dx*dx+dy*dy > float64(sqrad) {
				continue
			}

			result = append(result, char)

		}
	}
	return result
}

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

func (w *WorldRegion) DropItemChecker() []int32 {
	var result []int32
	w.ItemsExpireTime.ForEach(func(key int32, value int64) bool {
		if value <= time.Now().Unix() {
			key_ := strconv.FormatInt(int64(key), 10)

			//item, _ := w.ItemsInRegion.Load(key_)

			w.ItemsInRegion.Delete(key_)
			w.ItemsExpireTime.Del(key)
			result = append(result, key)
		}
		return true
	})
	return result
}
