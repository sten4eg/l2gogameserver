package models

import (
	"context"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"math"
	"sync"
)

const InsertIntoDB = `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', 0, 0, 0, 0, '-1', 0, 0)`

// const UpdateInDB = `UPDATE items SET owner_id=$1, count=$2, loc=$3, loc_data=$4, enchant_level=$5, custom_type1=$6, custom_type2=$7, mana_left=$8, time=$9, agathion_energy=$10 WHERE object_id=$11`
const UpdateInDB = `UPDATE items SET owner_id=$1, count=$2 WHERE object_id=$3`
const RemoveFromDB = `DELETE FROM items WHERE object_id = $1`

type MyItem struct {
	// встроенный "шаблон" предмета
	*items.Item
	ObjectId            int32
	ownerId             int32
	Enchant             int16
	LocData             int32
	Count               int64
	Location            string
	Time                int
	AttackAttributeType attribute.Attribute
	AttackAttributeVal  int16
	Mana                int32
	AttributeDefend     [6]int16
	EnchantedOption     [3]int32
	sync.Mutex
	//UpdateType для обновления инвентаря
	LastChange int16
	//БД
	existsInDb bool
	storedInDb bool

	x int32
	y int32
	z int32
}

func (i *MyItem) SetObjectId(value int32) {
	i.ObjectId = value
}
func (i *MyItem) GetObjectId() int32 {
	return i.ObjectId
}
func (i *MyItem) GetOwnerId() int32 {
	return i.ownerId
}
func (i *MyItem) SetOwnerId(ownerId int32) {
	i.ownerId = ownerId
	i.storedInDb = false
}
func (i *MyItem) IsEquipped() int16 {
	if i.Location == InventoryLoc {
		return 0
	}
	return 1
}
func (i *MyItem) GetAttackElementType() attribute.Attribute {
	el := attribute.Attribute(-2) // none
	if i.IsWeapon() {
		el = i.AttackAttributeType
	}

	if el == attribute.None {
		if i.BaseAttributeAttack.Val > 0 {
			return i.getBaseAttributeElement()
		}
	}

	return el
}
func (i *MyItem) getBaseAttributeElement() attribute.Attribute {
	return i.BaseAttributeAttack.Type
}
func (i *MyItem) GetCount() int64 {
	return i.Count
}
func (i *MyItem) SetCount(count int64) {
	i.Count = count
}
func (i *MyItem) SetEnchant(value int16) {
	i.Enchant = value
}
func (i *MyItem) GetEnchant() int16 {
	return i.Enchant
}
func (i *MyItem) GetAttackElementPower() int16 {
	return i.AttackAttributeVal
}
func (i *MyItem) GetElementDefAttr() [6]int16 {
	return i.AttributeDefend
}
func (i *MyItem) GetEnchantedOption() [3]int32 {
	return i.EnchantedOption
}
func (i *MyItem) GetLocation() string {
	return i.Location
}

func (i *MyItem) GetUpdateType() int16 {
	return i.LastChange
}
func (i *MyItem) SetUpdateType(lastChange int16) {
	i.LastChange = lastChange
}
func (i *MyItem) GetLocData() int32 {
	return i.LocData
}
func (i *MyItem) GetMana() int32 {
	return i.Mana
}
func (i *MyItem) GetDefaultPrice() int {
	return i.Item.DefaultPrice
}
func (i *MyItem) SetPrice(value int64) {
	i.Price = int(value)
}
func (i *MyItem) GetPrice() int64 {
	return int64(i.Item.Price)
}
func (i *MyItem) IsAvailable(character interfaces.CharacterI, allowAdena, allowNonTradable bool) bool {
	return !utils.I2B(i.IsEquipped()) &&
		i.GetItemType2() != items.Quest &&
		(i.GetItemType2() != items.Money || i.GetItemType1() != items.ShieldArmor) &&
		character.GetActiveEnchantItemId() != i.GetObjectId() &&
		(allowAdena || i.GetId() != config.AdenaId)
	//allowNonTradable
}

func (i *MyItem) ChangeCount(count int) {
	if count == 0 {
		return
	}
	//TODO log [old := i.GetCount()]
	var max int
	if i.GetId() == config.AdenaId {
		max = config.MaxAdena
	} else {
		max = math.MaxInt64
	}

	if count > 0 && int(i.GetCount()) > max-count {
		i.SetCount(int64(max))
	} else {
		i.SetCount(i.GetCount() + int64(count))
	}

	if i.GetCount() < 0 {
		i.SetCount(0)
	}

	i.storedInDb = false
	i.SetUpdateType(UpdateTypeModify)
}
func (i *MyItem) UpdateDB() {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()
	if i.existsInDb {
		if i.ownerId == 0 || i.GetCount() == 0 { //TODO добавить проверки для удаления итема из бд
			_, err = dbConn.Exec(context.Background(), RemoveFromDB, i.GetObjectId())
			i.existsInDb = false
			i.storedInDb = false
		} else {
			if !i.storedInDb {
				_, err = dbConn.Exec(context.Background(), UpdateInDB, i.ownerId, i.GetCount(), i.GetObjectId())
				i.storedInDb = true
			}
		}
	} else {
		if i.ownerId == 0 || i.GetCount() == 0 {
			return
		}
		_, err = dbConn.Exec(context.Background(), InsertIntoDB, i.ownerId, i.ObjectId, i.Item.Id, i.Count)
		i.existsInDb = true
		i.storedInDb = true
		//TODO доделать функцию
	}
}

func CreateItem(itemId int, count int) interfaces.MyItemInterface {
	item, ok := items.GetItemInfo(itemId)
	if !ok {
		return nil
	}
	mt := MyItem{
		Item:       item,
		ObjectId:   idfactory.GetNext(),
		Enchant:    0,
		Count:      int64(count),
		Location:   InventoryLoc,
		existsInDb: false,
		storedInDb: false,
	}
	return &mt
}

// TODO додолеть
func DestroyItem(item interfaces.MyItemInterface) {
	item.SetCount(0)
	item.SetOwnerId(0)
	// item.setItemLocation(ItemLocation.VOID); ?
	item.SetUpdateType(UpdateTypeRemove)

	// L2World.getInstance().removeObject(item); ?
	// IdFactory.getInstance().releaseId(item.getObjectId()); ?
}

func (i *MyItem) SetCoordinate(x, y, z int32) {
	i.x = x
	i.y = y
	i.z = z
}

func (i *MyItem) GetCoordinate() (x, y, z int32) {
	return i.x, i.y, i.z
}

func (i *MyItem) WriteItem(buffer *packets.Buffer) {
	buffer.WriteD(i.GetObjectId())
	buffer.WriteD(i.GetId())
	buffer.WriteD(i.GetLocData())
	buffer.WriteQ(i.GetCount())
	buffer.WriteH(i.GetItemType2())
	buffer.WriteH(0)
	buffer.WriteH(i.IsEquipped())
	buffer.WriteD(i.GetBodyPart())
	buffer.WriteH(i.GetEnchant())
	buffer.WriteH(i.GetItemType2())
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	i.writeItemElementalAndEnchant(buffer)
}

func (i *MyItem) writeItemElementalAndEnchant(buffer *packets.Buffer) {
	buffer.WriteH(int16(i.GetAttackElementType()))
	buffer.WriteH(i.GetAttackElementPower())

	for i := 0; i < 6; i++ {
		buffer.WriteH(0)
	}

	for _, op := range i.GetEnchantedOption() {
		buffer.WriteH(int16(op))
	}
}

func (i *MyItem) SetStoreCount(value int64) {
	return
}

func (i *MyItem) GetStoreCount() int64 {
	return 0
}
