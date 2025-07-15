package models

import (
	"database/sql"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
	"math"
	"sync"
)

const InsertIntoDB = `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', 0, 0, 0, 0, '-1', 0, 0)`

// const UpdateInDB = `UPDATE items SET owner_id=$1, count=$2, loc=$3, loc_data=$4, enchant_level=$5, custom_type1=$6, custom_type2=$7, mana_left=$8, time=$9, agathion_energy=$10 WHERE object_id=$11`
const UpdateInDB = `UPDATE items SET owner_id=$1, count=$2 WHERE object_id=$3`
const RemoveFromDB = `DELETE FROM items WHERE object_id = $1`

type PlayerItem struct {
	ItemInfo *items.Item

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
	Price               int64
	BonusStats          []items.ItemBonusStat
	sync.Mutex
	//UpdateType для обновления инвентаря
	LastChange int16
	//БД
	existsInDb bool
	storedInDb bool

	x int32
	y int32
	z int32

	// Ссылка на персонажа для методов предмета
	character interfaces.CharacterI
}

func (i *PlayerItem) GetItemByItemId(itemId int) []interfaces.MyItemInterface {
	//TODO implement me
	panic("implement me")
}

func (i *PlayerItem) GetItemInfo() *items.Item {
	return i.ItemInfo
}

func (i *PlayerItem) SetObjectId(value int32) {
	i.ObjectId = value
}
func (i *PlayerItem) GetObjectId() int32 {
	return i.ObjectId
}
func (i *PlayerItem) GetOwnerId() int32 {
	return i.ownerId
}
func (i *PlayerItem) SetOwnerId(ownerId int32) {
	i.ownerId = ownerId
	i.storedInDb = false
}
func (i *PlayerItem) IsEquipped() int16 {
	if i.Location == InventoryLoc {
		return 0
	}
	return 1
}
func (i *PlayerItem) GetAttackElementType() attribute.Attribute {
	el := attribute.Attribute(-2) // none
	if i.IsWeapon() {
		el = i.AttackAttributeType
	}

	if el == attribute.None {
		if i.ItemInfo != nil && i.ItemInfo.BaseAttributeAttack.Val > 0 {
			return i.getBaseAttributeElement()
		}
	}

	return el
}
func (i *PlayerItem) getBaseAttributeElement() attribute.Attribute {
	if i.ItemInfo == nil {
		return attribute.None
	}
	return i.ItemInfo.BaseAttributeAttack.Type
}
func (i *PlayerItem) GetCount() int64 {
	return i.Count
}
func (i *PlayerItem) SetCount(count int64) {
	i.Count = count
}
func (i *PlayerItem) SetEnchant(value int16) {
	i.Enchant = value
}
func (i *PlayerItem) GetEnchant() int16 {
	return i.Enchant
}
func (i *PlayerItem) GetAttackElementPower() int16 {
	return i.AttackAttributeVal
}
func (i *PlayerItem) GetElementDefAttr() [6]int16 {
	return i.AttributeDefend
}
func (i *PlayerItem) GetEnchantedOption() [3]int32 {
	return i.EnchantedOption
}
func (i *PlayerItem) GetLocation() string {
	return i.Location
}

func (i *PlayerItem) GetUpdateType() int16 {
	return i.LastChange
}
func (i *PlayerItem) SetUpdateType(lastChange int16) {
	i.LastChange = lastChange
}
func (i *PlayerItem) GetLocData() int32 {
	return i.LocData
}
func (i *PlayerItem) SetLocData(loc int32) {
	i.LocData = loc
}
func (i *PlayerItem) GetMana() int32 {
	return i.Mana
}
func (i *PlayerItem) GetDefaultPrice() int {
	return i.ItemInfo.DefaultPrice
}
func (i *PlayerItem) SetPrice(value int64) {
	i.Price = value
}
func (i *PlayerItem) GetPrice() int64 {
	return i.ItemInfo.Price
}
func (i *PlayerItem) IsAvailable(character interfaces.CharacterI, allowAdena, allowNonTradable bool) bool {
	return !utils.I2B(i.IsEquipped()) &&
		i.GetItemType2() != items.Quest &&
		(i.GetItemType2() != items.Money || i.GetItemType1() != items.ShieldArmor) &&
		character.GetActiveEnchantItemId() != i.GetObjectId() &&
		(allowAdena || i.GetId() != config.AdenaId)
	//allowNonTradable
}

func (i *PlayerItem) ChangeCount(count int) {
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
func (i *PlayerItem) UpdateDB(db *sql.DB) {

	if i.existsInDb {
		if i.ownerId == 0 || i.GetCount() == 0 { //TODO добавить проверки для удаления итема из бд
			_, err := db.Exec(RemoveFromDB, i.GetObjectId())
			if err != nil {
				log.Println(err)
			}
			i.existsInDb = false
			i.storedInDb = false
		} else {
			if !i.storedInDb {
				_, err := db.Exec(UpdateInDB, i.ownerId, i.GetCount(), i.GetObjectId())
				if err != nil {
					log.Println(err)
				}
				i.storedInDb = true
			}
		}
	} else {
		if i.ownerId == 0 || i.GetCount() == 0 {
			return
		}
		_, err := db.Exec(InsertIntoDB, i.ownerId, i.ObjectId, i.ItemInfo.Id, i.Count)
		if err != nil {
			log.Println(err)
		}
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
	mt := PlayerItem{
		ItemInfo:   item,
		ObjectId:   idfactory.GetNext(),
		Enchant:    0,
		Count:      int64(count),
		Location:   InventoryLoc,
		existsInDb: false,
		storedInDb: false,
	}
	return &mt
}

// CreateItemWithCharacter создает предмет с установленной ссылкой на персонажа
func CreateItemWithCharacter(itemId int, count int, character interfaces.CharacterI) interfaces.MyItemInterface {
	item := CreateItem(itemId, count)
	if item != nil {
		item.SetCharacter(character)
	}
	return item
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

func (i *PlayerItem) SetCoordinate(x, y, z int32) {
	i.x = x
	i.y = y
	i.z = z
}

func (i *PlayerItem) GetCoordinate() (x, y, z int32) {
	return i.x, i.y, i.z
}

func (i *PlayerItem) WriteItem(buffer *packets.Buffer) {
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

func (i *PlayerItem) writeItemElementalAndEnchant(buffer *packets.Buffer) {
	buffer.WriteH(int16(i.GetAttackElementType()))
	buffer.WriteH(i.GetAttackElementPower())

	for i := 0; i < 6; i++ {
		buffer.WriteH(0)
	}

	for _, op := range i.GetEnchantedOption() {
		buffer.WriteH(int16(op))
	}
}

func (i *PlayerItem) SetStoreCount(value int64) {
	return
}

func (i *PlayerItem) GetStoreCount() int64 {
	return 0
}

func (i *PlayerItem) GetId() int32 {
	if i.ItemInfo == nil {
		return 0
	}
	return i.ItemInfo.GetId()
}

func (i *PlayerItem) GetTime() int {
	return i.Time
}

// Методы для реализации BaseItemInterface
func (i *PlayerItem) GetBaseItem() interfaces.BaseItemInterface {
	return i.ItemInfo
}

func (i *PlayerItem) IsEquipable() bool {
	return i.ItemInfo.IsEquipable()
}

func (i *PlayerItem) IsHeavyArmor() bool {
	return i.ItemInfo.IsHeavyArmor()
}

func (i *PlayerItem) IsMagicArmor() bool {
	return i.ItemInfo.IsMagicArmor()
}

func (i *PlayerItem) IsArmor() bool {
	return i.ItemInfo.IsArmor()
}

func (i *PlayerItem) IsOnlyKamaelWeapon() bool {
	return i.ItemInfo.IsOnlyKamaelWeapon()
}

func (i *PlayerItem) IsWeapon() bool {
	return i.ItemInfo.IsWeapon()
}

func (i *PlayerItem) IsWeaponTypeNone() bool {
	return i.ItemInfo.IsWeaponTypeNone()
}

func (i *PlayerItem) IsStackable() bool {
	return i.ItemInfo.IsStackable()
}

func (i *PlayerItem) GetItemType1() int {
	return i.ItemInfo.GetItemType1()
}

func (i *PlayerItem) GetItemType2() int16 {
	return i.ItemInfo.GetItemType2()
}

func (i *PlayerItem) GetWeight() int {
	return i.ItemInfo.GetWeight()
}

func (i *PlayerItem) GetSlotBitType() int32 {
	if i.ItemInfo == nil {
		return 0
	}
	return int32(i.ItemInfo.SlotBitType)
}

func (i *PlayerItem) GetBodyPart() int32 {
	return i.ItemInfo.GetBodyPart()
}

// Добавляем недостающие методы для доступа к полям ItemInfo
func (i *PlayerItem) GetWeaponType() int16 {
	if i.ItemInfo == nil {
		return 0
	}
	return int16(i.ItemInfo.WeaponType)
}

func (i *PlayerItem) GetEtcItemType() int16 {
	if i.ItemInfo == nil {
		return 0
	}
	return int16(i.ItemInfo.EtcItemType)
}

func (i *PlayerItem) GetConsumeType() int32 {
	if i.ItemInfo == nil {
		return 0
	}
	return int32(i.ItemInfo.ConsumeType)
}

// Добавляем недостающие методы для полной реализации интерфейса
func (i *PlayerItem) GetName() string {
	if i.ItemInfo == nil {
		return ""
	}
	return i.ItemInfo.Name
}

// SetCharacter устанавливает ссылку на персонажа
func (i *PlayerItem) SetCharacter(character interfaces.CharacterI) {
	i.character = character
}

// GetCharacter возвращает ссылку на персонажа
func (i *PlayerItem) GetCharacter() interfaces.CharacterI {
	return i.character
}

// UseEquippableItem использует предмет без явной передачи character
// Пример использования:
// client.GetAccount().GetCurrentChar().GetInventory().GetItemByObjectId(384771).UseEquippableItem()
func (i *PlayerItem) UseEquippableItem() {
	if i.character == nil {
		return
	}

	// Получаем Character из интерфейса
	character, ok := i.character.(*Character)
	if !ok {
		return
	}

	// Вызываем существующую функцию
	UseEquippableItem(i, character)
}
