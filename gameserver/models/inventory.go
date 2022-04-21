package models

import (
	"context"
	"github.com/jackc/pgx/v4"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/armorType"
	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/gameserver/models/items/consumeType"
	"l2gogameserver/gameserver/models/items/etcItemType"
	"l2gogameserver/gameserver/models/items/weaponType"
	"log"
	"strconv"
	"strings"
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
)

type MyItem struct {
	items.Item
	ObjId               int32
	Enchant             int
	LocData             int32
	Count               int64
	Loc                 string
	Time                int
	AttackAttributeType attribute.Attribute
	AttackAttributeVal  int
	Mana                int
	AttributeDefend     [6]int16
}

type Inventory struct {
	Items []MyItem
}

func RestoreVisibleInventory(charId int32) [26]MyItem {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT object_id, item, loc_data, enchant_level FROM items WHERE owner_id= $1 AND loc= $2", charId, PaperdollLoc)
	if err != nil {
		panic(err)
	}

	var mts [26]MyItem

	for rows.Next() {
		var objId int
		var itemId int
		var enchantLevel int
		var locData int
		err = rows.Scan(&objId, &itemId, &locData, &enchantLevel)
		if err != nil {
			log.Println(err)
		}

		item, ok := items.GetItemFromStorage(itemId)
		if !ok {
			panic("Предмет не найден")
		}
		mt := MyItem{
			Item:    item,
			ObjId:   int32(objId),
			Enchant: enchantLevel,
			Count:   1,
			Loc:     PaperdollLoc,
		}
		mts[int32(locData)] = mt
	}
	return mts
}

// IsEquipable Можно ли надеть предмет
func (i *MyItem) IsEquipable() bool {
	return !((i.SlotBitType == items.SlotNone) || (i.EtcItemType == etcItemType.ARROW) || (i.EtcItemType == etcItemType.BOLT) || (i.EtcItemType == etcItemType.LURE))
}
func (i *MyItem) IsHeavyArmor() bool {
	return i.ArmorType == armorType.HEAVY
}
func (i *MyItem) IsMagicArmor() bool {
	return i.ArmorType == armorType.MAGIC
}
func (i *MyItem) IsArmor() bool {
	return i.ItemType == items.ShieldOrArmor
}
func (i *MyItem) IsOnlyKamaelWeapon() bool {
	return i.WeaponType == weaponType.RAPIER || i.WeaponType == weaponType.CROSSBOW || i.WeaponType == weaponType.ANCIENTSWORD
}
func (i *MyItem) IsWeapon() bool {
	return i.ItemType == items.Weapon
}
func (i *MyItem) IsWeaponTypeNone() bool {
	return i.WeaponType == weaponType.NONE
}
func GetMyItems(charId int32) Inventory {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	sqlString := "SELECT items.object_id, item, loc_data, enchant_level, count, loc, time, mana_left FROM items WHERE owner_id = $1"
	rows, err := dbConn.Query(context.Background(), sqlString, charId)
	if err != nil {
		panic(err)
	}

	var inventory Inventory

	for rows.Next() {
		var itm MyItem
		var id int

		err := rows.Scan(&itm.ObjId, &id, &itm.LocData, &itm.Enchant, &itm.Count, &itm.Loc, &itm.Time, &itm.Mana)
		if err != nil {
			panic(err)
		}

		it, ok := items.GetItemFromStorage(id)
		if ok {
			itm.Item = it

			if itm.IsWeapon() {

				itm.AttackAttributeType, itm.AttackAttributeVal = getAttributeForWeapon(itm.ObjId)
			} else if itm.IsArmor() {
				itm.AttributeDefend = getAttributeForArmor(itm.ObjId)
			}

			inventory.Items = append(inventory.Items, itm)
		}
	}

	return inventory
}

func getAttributeForWeapon(objId int32) (attribute.Attribute, int) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()
	el := attribute.Attribute(-2) // None

	var elementType, elementValue int
	err = dbConn.QueryRow(context.Background(), "SELECT element_type,element_value FROM item_elementals WHERE item_id = $1", objId).
		Scan(&elementType, &elementValue)

	if err == pgx.ErrNoRows {
		return el, 0
	} else if err != nil {
		panic(err)
	}

	el = attribute.Attribute(elementType)

	return el, elementValue
}

func getAttributeForArmor(objId int32) [6]int16 {
	var att [6]int16
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT element_type,element_value FROM item_elementals WHERE item_id = $1", objId)

	if err == pgx.ErrNoRows {
		return att
	} else if err != nil {
		panic(err)
	}

	for rows.Next() {
		var atType, atVal int
		err = rows.Scan(&atType, &atVal)
		if err != nil {
			panic(err)
		}
		att[atType] = int16(atVal)
	}

	return att
}

func (i *MyItem) IsEquipped() int16 {
	if i.Loc == InventoryLoc {
		return 0
	}
	return 1
}

func SaveInventoryInDB(inventory []MyItem) {
	if len(inventory) == 0 {
		return
	}

	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	var sb strings.Builder
	sb.WriteString("UPDATE items SET loc_data = mylocdata, loc = myloc FROM ( VALUES ")

	for i := range inventory {
		v := &inventory[i]

		sb.WriteString("(" + strconv.Itoa(int(v.LocData)) + ",'" + v.Loc + "'," + strconv.Itoa(int(v.ObjId)) + ")")

		if len(inventory)-1 != i {
			sb.WriteString(",")
		}
	}
	sb.WriteString(") as myval (mylocdata,myloc,myobjid) WHERE items.object_id = myval.myobjid")
	_, err = dbConn.Exec(context.Background(), sb.String())
	if err != nil {
		log.Println(err.Error())
	}
}

func GetActiveWeapon(inventory []MyItem, paperdoll [26]MyItem) *MyItem {
	q := paperdoll[PAPERDOLL_RHAND]
	for i := range inventory {
		v := &inventory[i]
		if v.ObjId == q.ObjId {
			return v
		}
	}
	return nil
}

// UseEquippableItem исользовать предмет который можно надеть на персонажа
func UseEquippableItem(selectedItem *MyItem, character *Character) {
	//todo надо как то обновлять paperdoll, или возвращать массив или же  вынести это в другой пакет
	log.Println(selectedItem.ObjId, " and equiped = ", selectedItem.IsEquipped())
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
	if (selectedItem.Id != 21163) && (formal.ObjId != 0) && (formal.SlotBitType == items.SlotAlldress) {
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
		if paperdoll[PAPERDOLL_LEAR].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_LEAR, selectedItem, character)
		} else if paperdoll[PAPERDOLL_REAR].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_REAR, selectedItem, character)
		} else {
			setPaperdollItem(PAPERDOLL_LEAR, selectedItem, character)
		}

	case items.SlotNeck:
		setPaperdollItem(PAPERDOLL_NECK, selectedItem, character)
	case items.SlotRFinger, items.SlotLFinger, items.SlotLrFinger:
		if paperdoll[PAPERDOLL_LFINGER].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_LFINGER, selectedItem, character)
		} else if paperdoll[PAPERDOLL_RFINGER].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_RFINGER, selectedItem, character)
		} else {
			setPaperdollItem(PAPERDOLL_LFINGER, selectedItem, character)
		}

	case items.SlotHair:
		hair := paperdoll[PAPERDOLL_HAIR]
		if hair.ObjId != 0 && hair.SlotBitType == items.SlotHairall {
			setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
		} else {
			setPaperdollItem(PAPERDOLL_HAIR, nil, character)
		}
		setPaperdollItem(PAPERDOLL_HAIR, selectedItem, character)
	case items.SlotHair2:
		hair2 := paperdoll[PAPERDOLL_HAIR]
		if hair2.ObjId != 0 && hair2.SlotBitType == items.SlotHairall {
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
		if (rh.ObjId != 0) && (rh.SlotBitType == items.SlotLrHand) && !(((rh.WeaponType == weaponType.BOW) && (selectedItem.EtcItemType == etcItemType.ARROW)) || ((rh.WeaponType == weaponType.CROSSBOW) && (selectedItem.EtcItemType == etcItemType.BOLT)) || ((rh.WeaponType == weaponType.FISHINGROD) && (selectedItem.EtcItemType == etcItemType.LURE))) {
			setPaperdollItem(PAPERDOLL_RHAND, nil, character)
		}
		setPaperdollItem(PAPERDOLL_LHAND, selectedItem, character)
	case items.SlotGloves:
		setPaperdollItem(PAPERDOLL_GLOVES, selectedItem, character)
	case items.SlotChest:
		setPaperdollItem(PAPERDOLL_CHEST, selectedItem, character)
	case items.SlotLegs:
		chest := paperdoll[PAPERDOLL_CHEST]
		if chest.ObjId != 0 && chest.SlotBitType == items.SlotFullArmor {
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
		panic("Не определен Slot для itemId: " + strconv.Itoa(selectedItem.Id))
	}
}

func setPaperdollItem(slot uint8, selectedItem *MyItem, character *Character) {
	// eсли selectedItem nil, то ищем предмет которых находиться в slot
	// переносим его в инвентарь, убираем бонусы этого итема у персонажа
	if selectedItem == nil {
		for i := range character.Inventory.Items {
			itemInInventory := &character.Inventory.Items[i]
			if itemInInventory.LocData == int32(slot) && itemInInventory.Loc == PaperdollLoc {
				itemInInventory.LocData = getFirstEmptySlot(character.Inventory.Items)
				itemInInventory.Loc = InventoryLoc
				character.Inventory.Items[i] = *itemInInventory
				log.Println(itemInInventory.Loc, itemInInventory.LocData)
				character.RemoveBonusStat(itemInInventory.BonusStats)
				break
			}
		}
		return
	}
	log.Println("не должен дойти")
	var oldItemInSelectedSlot *MyItem
	var inventoryKeyOldItemInSelectedSlot int
	var keyCurrentItem int

	for i := range character.Inventory.Items {
		v := &character.Inventory.Items[i]
		// находим предмет, который стоял на нужном слоте раннее
		// его необходимо переместить в инвентарь
		if v.LocData == int32(slot) && v.Loc == PaperdollLoc {
			inventoryKeyOldItemInSelectedSlot = i
			oldItemInSelectedSlot = v
		}
		// находим ключ предмет в инвентаре, который нужно поставить в слот
		if v.ObjId == selectedItem.ObjId {
			keyCurrentItem = i
		}

	}
	// если на нужном слоте был итем его нужно снять и положить в инвентарь
	// и убрать у персонажа бонусы которые он давал
	if oldItemInSelectedSlot != nil && oldItemInSelectedSlot.Id != 0 {
		oldItemInSelectedSlot.Loc = InventoryLoc
		oldItemInSelectedSlot.LocData = selectedItem.LocData
		character.Inventory.Items[inventoryKeyOldItemInSelectedSlot] = *oldItemInSelectedSlot
		selectedItem.LocData = int32(slot)
		selectedItem.Loc = PaperdollLoc

		character.RemoveBonusStat(oldItemInSelectedSlot.BonusStats)
	} else {
		selectedItem.LocData = int32(slot)
		selectedItem.Loc = PaperdollLoc
	}
	// добавить бонусы предмета персонажу
	character.AddBonusStat(selectedItem.BonusStats)
	character.Inventory.Items[keyCurrentItem] = *selectedItem

}

func getFirstEmptySlot(myItems []MyItem) int32 {
	limit := int32(80) // todo дефолтно 80 , но может быть больше

	for i := int32(0); i < limit; i++ {
		flag := false
		for j := range myItems {
			v := &myItems[j]
			if v.Loc == InventoryLoc && v.LocData == i {
				flag = true
				break
			}
		}
		if !flag {
			return i
		}
	}
	panic("не нашёл куда складывать итем")

}

func (i *MyItem) GetAttackElement() attribute.Attribute {
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
func DeleteItem(selectedItem *MyItem, character *Character) {
	//TODO переделать, не надо создавать новый inventiry
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	if selectedItem.Loc == PaperdollLoc {
		character.Paperdoll[selectedItem.LocData] = MyItem{}
	}
	var inventory Inventory
	for _, v := range character.Inventory.Items {
		if v.ObjId != selectedItem.ObjId {
			inventory.Items = append(inventory.Items, v)
		}
	}
	character.Inventory = inventory

	_, err = dbConn.Exec(context.Background(), "DELETE FROM items WHERE object_id = $1", selectedItem.ObjId)
	if err != nil {
		panic(err)
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
func AddItem(selectedItem MyItem, character *Character) Inventory {
	//Прежде чем просто добавить, необходимо проверить на существование предмета в инвентаре
	//Если он есть, тогда просто добавим к имеющимся предмету.
	//TODO: Однако, есть предметы (кроме оружия, брони, бижи), которые не стакуются, к примеру 7832
	//TODO: потом нужно определить тип предметов которые не стыкуются.
	for i := range character.Inventory.Items {
		itemInventory := &character.Inventory.Items[i]
		if selectedItem.Item.Id == itemInventory.Item.Id {
			character.Inventory.Items[i].Count = itemInventory.Count + character.Inventory.Items[i].Count
			return character.Inventory
		}
	}

	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	nitem := MyItem{
		Item:                selectedItem.Item,
		ObjId:               selectedItem.ObjId,
		Enchant:             selectedItem.Enchant,
		LocData:             selectedItem.LocData,
		Count:               selectedItem.Count,
		Loc:                 "",
		Time:                selectedItem.Time,
		AttackAttributeType: selectedItem.AttackAttributeType,
		AttackAttributeVal:  selectedItem.AttackAttributeVal,
		Mana:                selectedItem.Mana,
		AttributeDefend:     [6]int16{},
	}
	character.Inventory.Items = append(character.Inventory.Items, nitem)

	_, err = dbConn.Exec(context.Background(), `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', 0, 0, 0, 0, '-1', 0, 0)`, character.ObjectId, selectedItem.ObjId, selectedItem.Item.Id, selectedItem.Count)
	if err != nil {
		panic(err)
	}

	return character.Inventory
}

//RemoveItemCharacter Удаление предмета из инвентаря персонажа
// count - сколько надо удалить
func RemoveItemCharacter(character *Character, item *MyItem, count int64) {
	log.Println("Удаление предмета из инвентаря")
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	if item.Count < count || item.Count == 0 || count == 0 {
		log.Println("Неверное количество предметов для удаления")
	}

	if item.Count == count {
		DeleteItem(item, character)
		item = nil
	} else {
		newCount := item.Count - count
		_, err = dbConn.Exec(context.Background(),
			`UPDATE "items" SET "count" = $1 WHERE "owner_id" = $2 AND "object_id" = $3 AND "item" = $4`,
			newCount, character.ObjectId, item.ObjId, item.Id)
		if err != nil {
			panic(err)
		}
		item.Count = newCount
	}
}

func ExistItemObject(characterI interfaces.CharacterI, objectId int32, count int64) (*MyItem, bool) {
	character, ok := characterI.(*Character)
	if !ok {
		panic("ExistItemObject not character")
	}
	for _, item := range character.Inventory.Items {
		if item.ObjId == objectId && item.Count >= count {
			return &item, true
		}
	}
	return nil, false
}

func AddInventoryItem(character *Character, item *MyItem, count int64) (*MyItem, bool) {
	log.Println("Добавление предмета ", item.Name, "count : ", count)

	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	for _, inv := range character.Inventory.Items {
		if inv.Item.Id == item.Id {
			//Если предмет стакуемый, тогда изменим его значение
			log.Println("Предмет найден в инвентаре у ", character.CharName)
			if inv.ConsumeType == consumeType.Stackable || inv.ConsumeType == consumeType.Asset {
				inv.Count += count
				log.Println("Будет произведено обновление предмета и увеличино", inv.Count)
				_, err = dbConn.Exec(context.Background(), `UPDATE "items" SET "count" = $1 WHERE "owner_id" = $2 AND "item" = $3`, inv.Count, character.ObjectId, inv.Item.Id)
				if err != nil {
					panic(err)
				}
				return &inv, true
			} else { //Если предмет не стакуемый, тогда добавим новое значение
				log.Println("Предмет не стыкуемый, будем делать новую запись для", character.CharName)
				item.ObjId = idfactory.GetNext()
				item.Count = count
				item.LocData = getFirstEmptySlot(character.Inventory.Items)
				character.Inventory.Items = append(character.Inventory.Items, *item)
				_, err = dbConn.Exec(context.Background(), `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', 0, 0, 0, 0, '-1', 0, 0)`, character.ObjectId, item.ObjId, item.Item.Id, item.Count)
				if err != nil {
					panic(err)
				}
				return item, true
			}
		}
	}

	log.Println("Предмет не найден, будем добавлять новый предмет для", character.CharName)
	item.ObjId = idfactory.GetNext()
	item.Count = count
	item.LocData = getFirstEmptySlot(character.Inventory.Items)
	character.Inventory.Items = append(character.Inventory.Items, *item)
	_, err = dbConn.Exec(context.Background(), `INSERT INTO "items" ("owner_id", "object_id", "item", "count", "enchant_level", "loc", "loc_data", "time_of_use", "custom_type1", "custom_type2", "mana_left", "time", "agathion_energy") VALUES ($1, $2, $3, $4, 0, 'INVENTORY', $5, 0, 0, 0, '-1', 0, 0)`, character.ObjectId, item.ObjId, item.Item.Id, item.Count, item.LocData)
	if err != nil {
		panic(err)
	}

	return item, false
}

func (i *Inventory) RemoveItem(character *Character, item *MyItem, count int64) (*MyItem, int16, bool) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	for index, itm := range i.Items {
		if itm.Id == item.Id {
			log.Println("Удаление найдено", item.Name, count, character.CharName)
			if itm.ConsumeType == consumeType.Stackable || itm.ConsumeType == consumeType.Asset {
				itm.Count -= count
				if itm.Count <= 0 {
					log.Println("Ожидаюсь тут..2.")
					_, err = dbConn.Exec(context.Background(), `DELETE FROM "items" WHERE "owner_id" = $1 AND "object_id" = $2 AND "item" = $3`, character.ObjectId, itm.ObjId, itm.Id)
					if err != nil {
						panic(err)
					}
					i.Items = append(i.Items[:index], i.Items[index+1:]...)
					return nil, UpdateTypeRemove, true
				} else {
					log.Println("Ожидаюсь тут..1.")
					_, err = dbConn.Exec(context.Background(), `UPDATE "items" SET "count" = $1 WHERE "owner_id" = $2 AND "object_id" = $3 AND "item" = $4`, itm.Count, character.ObjectId, itm.ObjId, itm.Id)
					if err != nil {
						panic(err)
					}
					i.Items[index].Count = itm.Count
					log.Println(i.Items[index].Id, i.Items[index].Count)
					return &i.Items[index], UpdateTypeModify, true
				}
			} else {
				log.Println("Ожидаюсь тут..3.")
				_, err = dbConn.Exec(context.Background(), `DELETE FROM "items" WHERE "owner_id" = $1 AND "object_id" = $2 AND "item" = $3`, character.ObjectId, itm.ObjId, itm.Item.Id)
				if err != nil {
					panic(err)
				}
				i.Items = append(i.Items[:index], i.Items[index+1:]...)
				return nil, UpdateTypeRemove, true
			}
		}
	}
	log.Println("Удаление не найдено")
	return &MyItem{}, UpdateTypeModify, false
}

func (i *Inventory) ExistItemID(id int) (*MyItem, bool) {
	for _, item := range i.Items {
		if item.Id == id {
			return &item, true
		}
	}
	return &MyItem{}, false
}
