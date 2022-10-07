package models

import (
	"github.com/alphadose/haxmap"
	"github.com/puzpuzpuz/xsync"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
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

	go DropItemChecker(&newRegion)
	return &newRegion
}

func DropItemChecker(region interfaces.WorldRegioner) {
	for {
		ids := region.DropItemChecker()
		for _, id := range ids {
			buffer := delObj(id)
			pb := utils.GetPacketByte()
			pb.SetData(buffer.Bytes())
			for _, r := range region.GetNeighbors() {
				for _, character := range r.GetCharsInRegion() {
					logger.Info.Println(character.GetName())
					character.EncryptAndSend(pb.GetData())
				}
			}
			pb.Release()
			packets.Put(buffer)
		}
		time.Sleep(time.Second * 5)
	}
}

// delObj копия serverpackets.DeleteObject
// написана тут чтобы небыло цикл зависимости
func delObj(objectId int32) *packets.Buffer {
	buffer := packets.Get()
	buffer.WriteSingleByte(0x08)
	buffer.WriteD(objectId)
	buffer.WriteD(0x00)
	return buffer

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
	return GetNeighbors(w.TileX, w.TileY, w.TileZ)
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
	} else if x > RegionsX {
		x = RegionsX
	}
	return x
}
func validY(y int32) int32 {
	if y < 0 {
		y = 0
	} else if y > RegionsY {
		y = RegionsY
	}
	return y
}
func validZ(z int32) int32 {
	if z < 0 {
		z = 0
	} else if z > RegionsZ {
		z = RegionsZ
	}
	return z
}

func regionX(x int32) int32 {
	return (x >> ShiftBy) + int32(OffsetX)
}
func regionY(y int32) int32 {
	return (y >> ShiftBy) + int32(OffsetY)
}
func regionZ(z int32) int32 {
	return (z >> ShiftByZ) + int32(OffsetZ)
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
func (w *WorldRegion) DropItemChecker() []int32 {
	var result []int32

	if w == nil {
		return result
	}

	now := time.Now().Unix()

	w.ItemsExpireTime.ForEach(func(key int32, value int64) bool {
		if value <= now {
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
