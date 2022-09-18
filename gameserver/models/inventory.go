package models

import (
	"context"
	"github.com/jackc/pgx/v4"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/gameserver/models/items/consumeType"
	"l2gogameserver/gameserver/models/items/etcItemType"
	"l2gogameserver/gameserver/models/items/weaponType"
	"l2gogameserver/utils"
	"strconv"
	"strings"
	"sync"
)

const (
	PAPERDOLL_UNDER      uint8 = 0
	PAPERDOLL_HEAD       uint8 = 1
	PAPERDOLL_HAIR       uint8 = 2
	PAPERDOLL_HAIR2      uint8 = 3
	PAPERDOLL_NECK       uint8 = 4
	PAPERDOLL_RHAND      uint8 = 5
	PAPERDOLL_CHEST      uint8 = 6
	PAPERDOLL_LHAND      uint8 = 7
	PAPERDOLL_REAR       uint8 = 8
	PAPERDOLL_LEAR       uint8 = 9
	PAPERDOLL_GLOVES     uint8 = 10
	PAPERDOLL_LEGS       uint8 = 11
	PAPERDOLL_FEET       uint8 = 12
	PAPERDOLL_RFINGER    uint8 = 13
	PAPERDOLL_LFINGER    uint8 = 14
	PAPERDOLL_LBRACELET  uint8 = 15
	PAPERDOLL_RBRACELET  uint8 = 16
	PAPERDOLL_DECO1      uint8 = 17
	PAPERDOLL_DECO2      uint8 = 18
	PAPERDOLL_DECO3      uint8 = 19
	PAPERDOLL_DECO4      uint8 = 20
	PAPERDOLL_DECO5      uint8 = 21
	PAPERDOLL_DECO6      uint8 = 22
	PAPERDOLL_CLOAK      uint8 = 23
	PAPERDOLL_BELT       uint8 = 24
	PAPERDOLL_TOTALSLOTS uint8 = 25

	PaperdollLoc string = "PAPERDOLL"
	InventoryLoc string = "INVENTORY"

	UpdateTypeUnchanged int16 = 0
	UpdateTypeAdd       int16 = 1
	UpdateTypeModify    int16 = 2
	UpdateTypeRemove    int16 = 3

	NoBlockMode = -1 // no block
	//TODO остальные типы блокировок не понятно для чего
	//Block0 - block items from _invItems, allow usage of other items
	//BlockMode1 - allow usage of items from _invItems, block other items

)

// Inventory реализует InventoryInterface
type Inventory struct {
	Items []MyItem
	//BlockItems содержит не objectId, а id предметов
	BlockItems  []int32
	BlockMode   int32
	TotalWeight int32 //todo где заполнять
	mu          sync.Mutex
}

func NewInventory() Inventory {
	return Inventory{
		BlockMode: -1,
	}
}
func (i *Inventory) GetItemByObjectId(id int32) interfaces.MyItemInterface {
	for _, item := range i.Items {
		if item.ObjectId == id {
			return &item
		}
	}
	return nil
}
func (i *Inventory) GetItemByItemId(itemId int) interfaces.MyItemInterface {
	for _, item := range i.Items {
		if item.Id == itemId {
			return &item
		}
	}
	return nil
}
func (i *Inventory) CanManipulateWithItemId(id int32) bool {
	return (i.BlockMode != 0 || !utils.Contains(i.BlockItems, id)) && i.BlockMode != 1 || utils.Contains(i.BlockItems, id)
}
func (i *Inventory) GetItemsWithUpdatedType() []interfaces.MyItemInterface {
	var res []interfaces.MyItemInterface
	for key := range i.Items {
		if i.Items[key].LastChange == UpdateTypeModify {
			res = append(res, &i.Items[key])
		}
	}
	return res
}
func (i *Inventory) Lock() {
	i.mu.Lock()
}
func (i *Inventory) Unlock() {
	i.mu.Unlock()
}

func (i *Inventory) SetAllItemsUpdatedTypeNone() {
	for _, v := range i.Items {
		v.LastChange = UpdateTypeUnchanged
	}
}
func (i *Inventory) ValidateWeight(weight int) bool {
	return int(i.TotalWeight)+weight < 69000 //TODO заменить на реальный допустимый вес
}
func (i *Inventory) ValidateCapacity(slots int, owner interfaces.CharacterI) bool {
	return len(i.Items)+slots <= int(owner.GetInventoryLimit())
}
func (i *Inventory) TransferItem(objectId int32, count int, target interfaces.InventoryInterface, actor interfaces.CharacterI) interfaces.MyItemInterface {
	if target == nil {
		return nil
	}

	sourceItem := i.GetItemByObjectId(objectId)
	if sourceItem == nil {
		return nil
	}

	var targetItem interfaces.MyItemInterface
	if sourceItem.IsStackable() {
		targetItem = target.GetItemByItemId(int(sourceItem.GetId()))
	} else {
		targetItem = nil
	}

	sourceItem.Lock()
	defer sourceItem.Unlock()

	//if i.GetItemByObjectId(objectId) != sourceItem { TODO объект одинаковый, но мьютекс имеет разные состояния
	//	return nil
	//}

	if count > int(sourceItem.GetCount()) {
		count = int(sourceItem.GetCount())
	}

	if int(sourceItem.GetCount()) == count && targetItem == nil {
		i.RemoveItem(sourceItem)
		target.AddItem(sourceItem, actor)
		targetItem = sourceItem
	} else {
		if int(sourceItem.GetCount()) > count {
			sourceItem.ChangeCount(-count)
		} else {
			i.RemoveItem(sourceItem)
		}

		if targetItem != nil {
			targetItem.ChangeCount(count)
		} else {
			targetItem = target.AddItem(sourceItem, actor)
		}
	}

	sourceItem.UpdateDB(actor.GetObjectId())
	if targetItem != sourceItem && targetItem != nil {
		targetItem.UpdateDB(actor.GetObjectId())
	}
	//TODO проверка isAugmented
	i.RefreshWeight()
	target.RefreshWeight()

	return sourceItem
}
func (i *Inventory) RemoveItem(removeItem interfaces.MyItemInterface) bool {
	for index, item := range i.Items {
		if item.GetId() == removeItem.GetId() {
			i.Items = append(i.Items[:index], i.Items[index+1:]...)
			return true
		}
	}
	return false
}
func (i *Inventory) AddItem(item interfaces.MyItemInterface, actor interfaces.CharacterI) interfaces.MyItemInterface {
	oldItem := i.GetItemByItemId(int(item.GetId()))

	if oldItem != nil && oldItem.IsStackable() {
		count := int(item.GetCount())
		oldItem.ChangeCount(count)
		oldItem.SetUpdateType(2) //TODO заменить на константу

		//TODO destroyItem()
		item.UpdateDB(actor.GetObjectId())
		item = oldItem
		//TODO добавить обновление адены
	} else {
		item.SetUpdateType(1) //TODO добавить константу Added
		i.Items = append(i.Items, *item.(*MyItem))
		item.GetObjectId()
		item.UpdateDB(actor.GetObjectId())
	}
	i.RefreshWeight()

	//TODO Манипуляции с аденой

	return item
}
func (i *Inventory) RefreshWeight() {
	weight := 0
	for _, item := range i.Items {
		weight += item.GetWeight() * int(item.GetCount())
	}
	i.TotalWeight = int32(weight)
}
func RestoreVisibleInventory(charId int32) [26]MyItem {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT object_id, item, loc_data, enchant_level FROM items WHERE owner_id= $1 AND loc= $2", charId, PaperdollLoc)
	if err != nil {
		logger.Error.Panicln(err)
	}

	var mts [26]MyItem

	for rows.Next() {
		var objId int
		var itemId int
		var enchantLevel int16
		var locData int
		err = rows.Scan(&objId, &itemId, &locData, &enchantLevel)
		if err != nil {
			logger.Info.Println(err)
		}

		item, ok := items.GetItemFromStorage(itemId)
		if !ok {
			logger.Error.Panicln("Предмет не найден")
		}
		mt := MyItem{
			Item:     item,
			ObjectId: int32(objId),
			Enchant:  enchantLevel,
			Count:    1,
			Location: PaperdollLoc,
		}
		mts[int32(locData)] = mt
	}
	return mts
}

func GetMyItems(charId int32) []MyItem {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	sqlString := "SELECT items.object_id, item, loc_data, enchant_level, count, loc, time, mana_left FROM items WHERE owner_id = $1"
	rows, err := dbConn.Query(context.Background(), sqlString, charId)
	if err != nil {
		logger.Error.Panicln(err)
	}

	itemsInInventory := make([]MyItem, 0, 80)
	for rows.Next() {
		var itm MyItem
		var id int

		err := rows.Scan(&itm.ObjectId, &id, &itm.LocData, &itm.Enchant, &itm.Count, &itm.Location, &itm.Time, &itm.Mana)
		if err != nil {
			logger.Error.Panicln(err)
		}

		it, ok := items.GetItemFromStorage(id)
		if ok {
			itm.Item = it

			if itm.IsWeapon() {
				itm.AttackAttributeType, itm.AttackAttributeVal = getAttributeForWeapon(itm.ObjectId)
			} else if itm.IsArmor() {
				itm.AttributeDefend = getAttributeForArmor(itm.ObjectId)
			}

			itemsInInventory = append(itemsInInventory, itm)
		}
	}

	return itemsInInventory
}

func getAttributeForWeapon(objId int32) (attribute.Attribute, int16) {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()
	elementType := attribute.Attribute(-2) // None

	var elementValue int16
	err = dbConn.QueryRow(context.Background(), "SELECT element_type,element_value FROM item_elementals WHERE item_id = $1", objId).
		Scan(&elementType, &elementValue)

	if err == pgx.ErrNoRows {
		return elementType, 0
	} else if err != nil {
		logger.Error.Panicln(err)
	}

	return elementType, elementValue
}

func getAttributeForArmor(objId int32) [6]int16 {
	var att [6]int16
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT element_type,element_value FROM item_elementals WHERE item_id = $1", objId)

	if err == pgx.ErrNoRows {
		return att
	} else if err != nil {
		logger.Error.Panicln(err)
	}

	for rows.Next() {
		var atType, atVal int
		err = rows.Scan(&atType, &atVal)
		if err != nil {
			logger.Error.Panicln(err)
		}
		att[atType] = int16(atVal)
	}

	return att
}

func SaveInventoryInDB(inventory []MyItem) {
	if len(inventory) == 0 {
		return
	}

	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	var sb strings.Builder
	sb.WriteString("UPDATE items SET loc_data = mylocdata, loc = myloc FROM ( VALUES ")

	for i := range inventory {
		v := &inventory[i]

		sb.WriteString("(" + strconv.Itoa(int(v.LocData)) + ",'" + v.Location + "'," + strconv.Itoa(int(v.ObjectId)) + ")")

		if len(inventory)-1 != i {
			sb.WriteString(",")
		}
	}
	sb.WriteString(") as myval (mylocdata,myloc,myobjid) WHERE items.object_id = myval.myobjid")
	_, err = dbConn.Exec(context.Background(), sb.String())
	if err != nil {
		logger.Info.Println(err.Error())
	}
}

func GetActiveWeapon(inventory []MyItem, paperdoll [26]MyItem) *MyItem {
	q := paperdoll[PAPERDOLL_RHAND]
	for i := range inventory {
		v := &inventory[i]
		if v.ObjectId == q.ObjectId {
			return v
		}
	}
	return nil
}

// UseEquippableItem исользовать предмет который можно надеть на персонажа
func UseEquippableItem(selectedItem *MyItem, character *Character) {
	//todo надо как то обновлять paperdoll, или возвращать массив или же  вынести это в другой пакет
	logger.Info.Println(selectedItem.ObjectId, " and equiped = ", selectedItem.IsEquipped())
	if selectedItem.IsEquipped() == 1 {
		unEquipAndRecord(selectedItem, character)
	} else {
		equipItemAndRecord(selectedItem, character)
	}
}

// unEquipAndRecord cнять предмет
func unEquipAndRecord(selectedItem *MyItem, character *Character) {
	switch selectedItem.SlotBitType {
	case items.SlotLEar:
		setPaperdollItem(PAPERDOLL_LEAR, nil, character)
	case items.SlotREar:
		setPaperdollItem(PAPERDOLL_REAR, nil, character)
	case items.SlotNeck:
		setPaperdollItem(PAPERDOLL_NECK, nil, character)
	case items.SlotRFinger:
		setPaperdollItem(PAPERDOLL_RFINGER, nil, character)
	case items.SlotLFinger:
		setPaperdollItem(PAPERDOLL_LFINGER, nil, character)
	case items.SlotHair:
		setPaperdollItem(PAPERDOLL_HAIR, nil, character)
	case items.SlotHair2:
		setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
	case items.SlotHairall: //todo Разобраться что тут на l2j
		setPaperdollItem(PAPERDOLL_HAIR, nil, character)
	case items.SlotHead:
		setPaperdollItem(PAPERDOLL_HEAD, nil, character)
	case items.SlotRHand, items.SlotLrHand:
		setPaperdollItem(PAPERDOLL_RHAND, nil, character)
	case items.SlotLHand:
		setPaperdollItem(PAPERDOLL_LHAND, nil, character)
	case items.SlotGloves:
		setPaperdollItem(PAPERDOLL_GLOVES, nil, character)
	case items.SlotChest, items.SlotAlldress, items.SlotFullArmor:
		setPaperdollItem(PAPERDOLL_CHEST, nil, character)
	case items.SlotLegs:
		setPaperdollItem(PAPERDOLL_LEGS, nil, character)
	case items.SlotBack:
		setPaperdollItem(PAPERDOLL_CLOAK, nil, character)
	case items.SlotFeet:
		setPaperdollItem(PAPERDOLL_FEET, nil, character)
	case items.SlotUnderwear:
		setPaperdollItem(PAPERDOLL_UNDER, nil, character)
	case items.SlotLBracelet:
		setPaperdollItem(PAPERDOLL_LBRACELET, nil, character)
	case items.SlotRBracelet:
		setPaperdollItem(PAPERDOLL_RBRACELET, nil, character)
	case items.SlotDeco:
		setPaperdollItem(PAPERDOLL_DECO1, nil, character)
	case items.SlotBelt:
		setPaperdollItem(PAPERDOLL_BELT, nil, character)
	}
}

// equipItemAndRecord одеть предмет
func equipItemAndRecord(selectedItem *MyItem, character *Character) {
	//todo проверка на приват Store, надо будет передавать character?
	// еще проверка на ITEM_CONDITIONS

	formal := character.Paperdoll[PAPERDOLL_CHEST]
	// Проверка надето ли офф. одежда и предмет не является букетом(id=21163)
	if (selectedItem.Id != 21163) && (formal.ObjectId != 0) && (formal.SlotBitType == items.SlotAlldress) {
		// только chest можно
		switch selectedItem.SlotBitType {
		case items.SlotLrHand, items.SlotLHand, items.SlotRHand, items.SlotLegs, items.SlotFeet, items.SlotGloves, items.SlotHead:
			return
		}
	}

	paperdoll := character.Paperdoll

	switch selectedItem.SlotBitType {
	case items.SlotLrHand:
		setPaperdollItem(PAPERDOLL_LHAND, nil, character)
		setPaperdollItem(PAPERDOLL_RHAND, selectedItem, character)
	case items.SlotLEar, items.SlotREar, items.SlotLrEar:
		if paperdoll[PAPERDOLL_LEAR].ObjectId == 0 {
			setPaperdollItem(PAPERDOLL_LEAR, selectedItem, character)
		} else if paperdoll[PAPERDOLL_REAR].ObjectId == 0 {
			setPaperdollItem(PAPERDOLL_REAR, selectedItem, character)
		} else {
			setPaperdollItem(PAPERDOLL_LEAR, selectedItem, character)
		}

	case items.SlotNeck:
		setPaperdollItem(PAPERDOLL_NECK, selectedItem, character)
	case items.SlotRFinger, items.SlotLFinger, items.SlotLrFinger:
		if paperdoll[PAPERDOLL_LFINGER].ObjectId == 0 {
			setPaperdollItem(PAPERDOLL_LFINGER, selectedItem, character)
		} else if paperdoll[PAPERDOLL_RFINGER].ObjectId == 0 {
			setPaperdollItem(PAPERDOLL_RFINGER, selectedItem, character)
		} else {
			setPaperdollItem(PAPERDOLL_LFINGER, selectedItem, character)
		}

	case items.SlotHair:
		hair := paperdoll[PAPERDOLL_HAIR]
		if hair.ObjectId != 0 && hair.SlotBitType == items.SlotHairall {
			setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
		} else {
			setPaperdollItem(PAPERDOLL_HAIR, nil, character)
		}
		setPaperdollItem(PAPERDOLL_HAIR, selectedItem, character)
	case items.SlotHair2:
		hair2 := paperdoll[PAPERDOLL_HAIR]
		if hair2.ObjectId != 0 && hair2.SlotBitType == items.SlotHairall {
			setPaperdollItem(PAPERDOLL_HAIR, nil, character)
		} else {
			setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
		}
		setPaperdollItem(PAPERDOLL_HAIR2, selectedItem, character)
	case items.SlotHairall:
		setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
		setPaperdollItem(PAPERDOLL_HAIR, selectedItem, character)
	case items.SlotHead:
		setPaperdollItem(PAPERDOLL_HEAD, selectedItem, character)
	case items.SlotRHand:
		//todo снять стрелы
		setPaperdollItem(PAPERDOLL_RHAND, selectedItem, character)
	case items.SlotLHand:
		rh := paperdoll[PAPERDOLL_RHAND]
		if (rh.ObjectId != 0) && (rh.SlotBitType == items.SlotLrHand) && !(((rh.WeaponType == weaponType.BOW) && (selectedItem.EtcItemType == etcItemType.ARROW)) || ((rh.WeaponType == weaponType.CROSSBOW) && (selectedItem.EtcItemType == etcItemType.BOLT)) || ((rh.WeaponType == weaponType.FISHINGROD) && (selectedItem.EtcItemType == etcItemType.LURE))) {
			setPaperdollItem(PAPERDOLL_RHAND, nil, character)
		}
		setPaperdollItem(PAPERDOLL_LHAND, selectedItem, character)
	case items.SlotGloves:
		setPaperdollItem(PAPERDOLL_GLOVES, selectedItem, character)
	case items.SlotChest:
		setPaperdollItem(PAPERDOLL_CHEST, selectedItem, character)
	case items.SlotLegs:
		chest := paperdoll[PAPERDOLL_CHEST]
		if chest.ObjectId != 0 && chest.SlotBitType == items.SlotFullArmor {
			setPaperdollItem(PAPERDOLL_CHEST, nil, character)
		}
		setPaperdollItem(PAPERDOLL_LEGS, selectedItem, character)
	case items.SlotBack:
		setPaperdollItem(PAPERDOLL_CLOAK, selectedItem, character)
	case items.SlotFeet:
		setPaperdollItem(PAPERDOLL_FEET, selectedItem, character)
	case items.SlotUnderwear:
		setPaperdollItem(PAPERDOLL_UNDER, selectedItem, character)
	case items.SlotLBracelet:
		setPaperdollItem(PAPERDOLL_LBRACELET, selectedItem, character)
	case items.SlotRBracelet:
		setPaperdollItem(PAPERDOLL_RBRACELET, selectedItem, character)
	case items.SlotDeco:
		setPaperdollItem(PAPERDOLL_DECO1, selectedItem, character)
	case items.SlotBelt:
		setPaperdollItem(PAPERDOLL_BELT, selectedItem, character)
	case items.SlotFullArmor:
		setPaperdollItem(PAPERDOLL_LEGS, nil, character)
		setPaperdollItem(PAPERDOLL_CHEST, selectedItem, character)
	case items.SlotAlldress:
		setPaperdollItem(PAPERDOLL_LEGS, nil, character)
		setPaperdollItem(PAPERDOLL_LHAND, nil, character)
		setPaperdollItem(PAPERDOLL_RHAND, nil, character)
		setPaperdollItem(PAPERDOLL_HEAD, nil, character)
		setPaperdollItem(PAPERDOLL_FEET, nil, character)
		setPaperdollItem(PAPERDOLL_GLOVES, nil, character)
		setPaperdollItem(PAPERDOLL_CHEST, selectedItem, character)
	default:
		logger.Error.Panicln("Не определен Slot для itemId: " + strconv.Itoa(selectedItem.Id))
	}
}

func setPaperdollItem(slot uint8, selectedItem *MyItem, character *Character) {
	// eсли selectedItem nil, то ищем предмет которых находиться в slot
	// переносим его в инвентарь, убираем бонусы этого итема у персонажа

	if selectedItem == nil {
		for i := range character.Inventory.Items {
			itemInInventory := &character.Inventory.Items[i]
			if itemInInventory.LocData == int32(slot) && itemInInventory.Location == PaperdollLoc {
				itemInInventory.LocData = getFirstEmptySlot(character.Inventory.Items)
				itemInInventory.Location = InventoryLoc
				character.Inventory.Items[i] = *itemInInventory
				itemInInventory.LastChange = UpdateTypeModify
				logger.Info.Println(itemInInventory.Location, itemInInventory.LocData)
				character.RemoveBonusStat(itemInInventory.BonusStats)
				break
			}
		}
		return
	}

	var oldItemInSelectedSlot *MyItem
	var inventoryKeyOldItemInSelectedSlot int
	var keyCurrentItem int

	for i := range character.Inventory.Items {
		v := &character.Inventory.Items[i]
		// находим предмет, который стоял на нужном слоте раннее
		// его необходимо переместить в инвентарь
		if v.LocData == int32(slot) && v.Location == PaperdollLoc {
			inventoryKeyOldItemInSelectedSlot = i
			oldItemInSelectedSlot = v
		}
		// находим ключ предмет в инвентаре, который нужно поставить в слот
		if v.ObjectId == selectedItem.ObjectId {
			keyCurrentItem = i
		}

	}
	// если на нужном слоте был итем его нужно снять и положить в инвентарь
	// и убрать у персонажа бонусы которые он давал
	if oldItemInSelectedSlot != nil && oldItemInSelectedSlot.Id != 0 {
		oldItemInSelectedSlot.Location = InventoryLoc
		oldItemInSelectedSlot.LocData = selectedItem.LocData
		character.Inventory.Items[inventoryKeyOldItemInSelectedSlot] = *oldItemInSelectedSlot
		character.Inventory.Items[inventoryKeyOldItemInSelectedSlot].LastChange = UpdateTypeModify
		selectedItem.LocData = int32(slot)
		selectedItem.Location = PaperdollLoc

		character.RemoveBonusStat(oldItemInSelectedSlot.BonusStats)
	} else {
		selectedItem.LocData = int32(slot)
		selectedItem.Location = PaperdollLoc
	}
	// добавить бонусы предмета персонажу
	character.AddBonusStat(selectedItem.BonusStats)
	character.Inventory.Items[keyCurrentItem] = *selectedItem
	character.Inventory.Items[keyCurrentItem].LastChange = UpdateTypeModify
}

func getFirstEmptySlot(myItems []MyItem) int32 {
	limit := int32(80) // todo дефолтно 80 , но может быть больше

	i := int32(0)
	for ; i < limit; i++ {
		flag := false
		for j := range myItems {
			v := &myItems[j]
			if v.Location == InventoryLoc && v.LocData == i {
				flag = true
				break
			}
		}
		if !flag {
			return i
		}
	}
	logger.Error.Panicln("не нашёл куда складывать итем")
	return 0
}

func DeleteItem(selectedItem *MyItem, character *Character) {
	//TODO переделать, не надо создавать новый inventiry
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	if selectedItem.Location == PaperdollLoc {
		character.Paperdoll[selectedItem.LocData] = MyItem{}
	}
	var inventory Inventory
	for _, v := range character.Inventory.Items {
		if v.ObjectId != selectedItem.ObjectId {
			inventory.Items = append(inventory.Items, v)
		}
	}
	character.Inventory = inventory

	_, err = dbConn.Exec(context.Background(), "DELETE FROM items WHERE object_id = $1", selectedItem.ObjectId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}
func GetPaperdollOrder() []uint8 {
	return []uint8{
		PAPERDOLL_UNDER,
		PAPERDOLL_REAR,
		PAPERDOLL_LEAR,
		PAPERDOLL_NECK,
		PAPERDOLL_RFINGER,
		PAPERDOLL_LFINGER,
		PAPERDOLL_HEAD,
		PAPERDOLL_RHAND,
		PAPERDOLL_LHAND,
		PAPERDOLL_GLOVES,
		PAPERDOLL_CHEST,
		PAPERDOLL_LEGS,
		PAPERDOLL_FEET,
		PAPERDOLL_CLOAK,
		PAPERDOLL_RHAND,
		PAPERDOLL_HAIR,
		PAPERDOLL_HAIR2,
		PAPERDOLL_RBRACELET,
		PAPERDOLL_LBRACELET,
		PAPERDOLL_DECO1,
		PAPERDOLL_DECO2,
		PAPERDOLL_DECO3,
		PAPERDOLL_DECO4,
		PAPERDOLL_DECO5,
		PAPERDOLL_DECO6,
		PAPERDOLL_BELT,
	}
}

// AddItem Добавление предмета
//func AddItem(selectedItem MyItem, character *Character) Inventory {
//	//Прежде чем просто добавить, необходимо проверить на существование предмета в инвентаре
//	//Если он есть, тогда просто добавим к имеющимся предмету.
//	//TODO: Однако, есть предметы (кроме оружия, брони, бижи), которые не стакуются, к примеру 7832
//	//TODO: потом нужно определить тип предметов которые не стыкуются.
//	for i := range character.Inventory.Items {
//		itemInventory := &character.Inventory.Items[i]
//		if selectedItem.Item.Id == itemInventory.Item.Id {
//			character.Inventory.Items[i].Count = itemInventory.Count + character.Inventory.Items[i].Count
//			return character.Inventory
//		}
//	}
//
//	dbConn, err := db.GetConn()
//	if err != nil {
//		logger.Error.Panicln(err)
//	}
//	defer dbConn.Release()
//
//	nitem := MyItem{
//		Item:                selectedItem.Item,
//		ObjectId:            selectedItem.ObjectId,
//		Enchant:             selectedItem.Enchant,
//		LocData:             selectedItem.LocData,
//		Count:               selectedItem.Count,
//		Location:            "",
//		Time:                selectedItem.Time,
//		AttackAttributeType: selectedItem.AttackAttributeType,
//		AttackAttributeVal:  selectedItem.AttackAttributeVal,
//		Mana:                selectedItem.Mana,
//		AttributeDefend:     [6]int16{},
//	}
//	character.Inventory.Items = append(character.Inventory.Items, nitem)
//
//	_, err = dbConn.Exec(context.Background(), `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', 0, 0, 0, 0, '-1', 0, 0)`, character.ObjectId, selectedItem.ObjectId, selectedItem.Item.Id, selectedItem.Count)
//	if err != nil {
//		logger.Error.Panicln(err)
//	}
//
//	return character.Inventory
//}

// RemoveItemCharacter Удаление предмета из инвентаря персонажа
// count - сколько надо удалить
func RemoveItemCharacter(character *Character, item *MyItem, count int64) {
	logger.Info.Println("Удаление предмета из инвентаря")
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	if item.Count < count || item.Count == 0 || count == 0 {
		logger.Info.Println("Неверное количество предметов для удаления")
	}

	if item.Count == count {
		DeleteItem(item, character)
		item = nil
	} else {
		newCount := item.Count - count
		_, err = dbConn.Exec(context.Background(),
			`UPDATE "items" SET "count" = $1 WHERE "owner_id" = $2 AND "object_id" = $3 AND "item" = $4`,
			newCount, character.ObjectId, item.ObjectId, item.Id)
		if err != nil {
			logger.Error.Panicln(err)
		}
		item.Count = newCount
	}
}

func ExistItemObject(characterI interfaces.CharacterI, objectId int32, count int64) (*MyItem, bool) {
	character, ok := characterI.(*Character)
	if !ok {
		logger.Error.Panicln("ExistItemObject not character")
	}
	for _, item := range character.Inventory.Items {
		if item.ObjectId == objectId && item.Count >= count {
			return &item, true
		}
	}
	return nil, false
}

// AddInventoryItem Добавление предмета в инвентарь пользователю
// Возращаемые параметры
// 1.Ссылка на предмет
// 2.Количество
// 3.Тип обновления/удаления/добавления
// 4.True если предмет найден
func AddInventoryItem(character *Character, item MyItem, count int64) (MyItem, int64, int16, bool) {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	for index, inv := range character.Inventory.Items {
		if inv.Item.Id == item.Id {
			if inv.IsEquipable() {
				logger.Info.Println("Нельзя передавать надетый предмет")
				return MyItem{}, 0, UpdateTypeUnchanged, false
			}
			//Если предмет стакуемый, тогда изменим его значение
			if inv.ConsumeType == consumeType.Stackable || inv.ConsumeType == consumeType.Asset {
				inv.Count = inv.Count + count
				_, err = dbConn.Exec(context.Background(), `UPDATE "items" SET "count" = $1, "loc" = 'INVENTORY' WHERE "owner_id" = $2 AND "item" = $3`, inv.Count, character.ObjectId, inv.Item.Id)
				if err != nil {
					logger.Error.Panicln(err)
				}
				character.Inventory.Items[index].Count = inv.Count
				inv.Location = "INVENTORY"
				return inv, inv.Count, UpdateTypeModify, true
			} else { //Если предмет не стакуемый, тогда добавим новое значение
				item.ObjectId = idfactory.GetNext()
				item.Count = count
				item.LocData = getFirstEmptySlot(character.Inventory.Items)
				character.Inventory.Items = append(character.Inventory.Items, item)
				_, err = dbConn.Exec(context.Background(), `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', 0, 0, 0, 0, '-1', 0, 0)`, character.ObjectId, item.ObjectId, item.Item.Id, item.Count)
				if err != nil {
					logger.Error.Panicln(err)
				}
				return item, count, UpdateTypeAdd, true
			}
		}
	}
	item.ObjectId = idfactory.GetNext()
	item.Count = count
	item.LocData = getFirstEmptySlot(character.Inventory.Items)
	character.Inventory.Items = append(character.Inventory.Items, item)
	_, err = dbConn.Exec(context.Background(), `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', $5, 0, 0, 0, '-1', 0, 0)`, character.ObjectId, item.ObjectId, item.Item.Id, item.Count, item.LocData)
	if err != nil {
		logger.Error.Panicln(err)
	}
	return item, count, UpdateTypeAdd, true
}

// RemoveItem Удаление предмета игрока
// 1.Возвращаемые параметры ссылка на предмет
// 2.Оставшейся кол-во предметов после удаления
// 3.tType удаления (Remove/Update)
// 4.Возращаемт False если предмет не был найден в инвентаре
func RemoveItem(character *Character, item *MyItem, count int64) (MyItem, int64, int16, bool) {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	for index, itm := range character.Inventory.Items {
		if itm.Id == item.Id {
			if itm.ConsumeType == consumeType.Stackable || itm.ConsumeType == consumeType.Asset {
				itm.Count -= count
				if itm.Count <= 0 {
					_, err = dbConn.Exec(context.Background(), `DELETE FROM "items" WHERE "owner_id" = $1 AND "object_id" = $2 AND "item" = $3`, character.ObjectId, itm.ObjectId, itm.Id)
					if err != nil {
						logger.Error.Panicln(err)
					}
					character.Inventory.Items = append(character.Inventory.Items[:index], character.Inventory.Items[index+1:]...)
					return MyItem{}, itm.Count, UpdateTypeRemove, true
				} else {
					_, err = dbConn.Exec(context.Background(), `UPDATE "items" SET "count" = $1 WHERE "owner_id" = $2 AND "object_id" = $3 AND "item" = $4`, itm.Count, character.ObjectId, itm.ObjectId, itm.Id)
					if err != nil {
						logger.Error.Panicln(err)
					}
					character.Inventory.Items[index].Count = itm.Count
					return character.Inventory.Items[index], itm.Count, UpdateTypeModify, true
				}
			} else {
				_, err = dbConn.Exec(context.Background(), `DELETE FROM "items" WHERE "owner_id" = $1 AND "object_id" = $2 AND "item" = $3`, character.ObjectId, itm.ObjectId, itm.Item.Id)
				if err != nil {
					logger.Error.Panicln(err)
				}
				character.Inventory.Items = append(character.Inventory.Items[:index], character.Inventory.Items[index+1:]...)
				return MyItem{}, 0, UpdateTypeRemove, true
			}
		}
	}
	return MyItem{}, 0, UpdateTypeModify, false
}
