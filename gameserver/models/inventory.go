package models

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"l2gogameserver/db"
	ItemsPkg "l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/armorType"
	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/gameserver/models/items/crystalType"
	"l2gogameserver/gameserver/models/items/etcItemType"
	"l2gogameserver/gameserver/models/items/materialType"
	"l2gogameserver/gameserver/models/items/weaponType"
	"log"
	"os"
	"strconv"
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

	Paperdoll string = "PAPERDOLL"
	Inventory string = "INVENTORY"

	UpdateTypeAdd    int16 = 1
	UpdateTypeModify int16 = 2
	UpdateTypeRemove int16 = 3
)

func RestoreVisibleInventory(charId int32) [26]MyItem {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT object_id, item, loc_data, enchant_level FROM items WHERE owner_id= $1 AND loc= $2", charId, Paperdoll)
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

		item := AllItems[itemId]
		mt := MyItem{
			Item:    item,
			ObjId:   int32(objId),
			Enchant: enchantLevel,
			Count:   1,
			Loc:     Paperdoll,
		}
		mts[int32(locData)] = mt
	}
	return mts
}

type Item struct {
	Id                     int                          `json:"id"`
	ItemType               ItemsPkg.ItemType            `json:"itemType"`
	Name                   string                       `json:"name"`
	Icon                   string                       `json:"icon"`
	SlotBitType            ItemsPkg.SlotBitType         `json:"slot_bit_type"`
	ArmorType              armorType.ArmorType          `json:"armor_type"`
	EtcItemType            etcItemType.EtcItemType      `json:"etcitem_type"`
	ItemMultiSkillList     []string                     `json:"item_multi_skill_list"`
	RecipeId               int                          `json:"recipe_id"`
	Weight                 int                          `json:"weight"`
	ConsumeType            string                       `json:"consume_type"`
	SoulShotCount          int                          `json:"soulshot_count"`
	SpiritShotCount        int                          `json:"spiritshot_count"`
	DropPeriod             int                          `json:"drop_period"`
	DefaultPrice           int                          `json:"default_price"`
	ItemSkill              string                       `json:"item_skill"`
	CriticalAttackSkill    string                       `json:"critical_attack_skill"`
	AttackSkill            string                       `json:"attack_skill"`
	MagicSkill             string                       `json:"magic_skill"`
	ItemSkillEnchantedFour string                       `json:"item_skill_enchanted_four"`
	MaterialType           materialType.MaterialType    `json:"material_type"`
	CrystalType            crystalType.CrystalType      `json:"crystal_type"`
	CrystalCount           int                          `json:"crystal_count"`
	IsTrade                bool                         `json:"is_trade"`
	IsDrop                 bool                         `json:"is_drop"`
	IsDestruct             bool                         `json:"is_destruct"`
	IsPrivateStore         bool                         `json:"is_private_store"`
	KeepType               int                          `json:"keep_type"`
	RandomDamage           int                          `json:"random_damage"`
	WeaponType             weaponType.WeaponType        `json:"weapon_type"`
	HitModify              int                          `json:"hit_modify"`
	AvoidModify            int                          `json:"avoid_modify"`
	ShieldDefense          int                          `json:"shield_defense"`
	ShieldDefenseRate      int                          `json:"shield_defense_rate"`
	AttackRange            int                          `json:"attack_range"`
	ReuseDelay             int                          `json:"reuse_delay"`
	MpConsume              int                          `json:"mp_consume"`
	Durability             int                          `json:"durability"`
	MagicWeapon            bool                         `json:"magic_weapon"`
	EnchantEnable          bool                         `json:"enchant_enable"`
	ElementalEnable        bool                         `json:"elemental_enable"`
	ForNpc                 bool                         `json:"for_npc"`
	IsOlympiadCanUse       bool                         `json:"is_olympiad_can_use"`
	IsPremium              bool                         `json:"is_premium"`
	BonusStats             []ItemsPkg.ItemBonusStat     `json:"stats"`
	DefaultAction          ItemsPkg.DefaultAction       `json:"default_action"`
	InitialCount           int                          `json:"initial_count"`
	ImmediateEffect        int                          `json:"immediate_effect"`
	CapsuledItems          []ItemsPkg.CapsuledItem      `json:"capsuled_items"`
	DualFhitRate           int                          `json:"dual_fhit_rate"`
	DamageRange            int                          `json:"damage_range"`
	Enchanted              int                          `json:"enchanted"`
	BaseAttributeAttack    ItemsPkg.BaseAttributeAttack `json:"base_attribute_attack"`
	BaseAttributeDefend    ItemsPkg.BaseAttributeDefend `json:"base_attribute_defend"`
	UnequipSkill           []string                     `json:"unequip_skill"`
	ItemEquipOption        []string                     `json:"item_equip_option"`
	CanMove                bool                         `json:"can_move"`
	DelayShareGroup        int                          `json:"delay_share_group"`
	Blessed                int                          `json:"blessed"`
	ReducedSoulshot        []string                     `json:"reduced_soulshot"`
	ExImmediateEffect      int                          `json:"ex_immediate_effect"`
	UseSkillDistime        int                          `json:"use_skill_distime"`
	Period                 int                          `json:"period"`
	EquipReuseDelay        int                          `json:"equip_reuse_delay"`
	Price                  int                          `json:"price"`
}

// AllItems - ONLY READ MAP, set in init server
var AllItems map[int]Item

type MyItem struct {
	Item
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

// IsEquipable Можно ли надеть предмет
func (i *MyItem) IsEquipable() bool {
	return !((i.SlotBitType == ItemsPkg.SlotNone) || (i.EtcItemType == etcItemType.ARROW) || (i.EtcItemType == etcItemType.BOLT) || (i.EtcItemType == etcItemType.LURE))
}
func (i *MyItem) IsHeavyArmor() bool {
	return i.ArmorType == armorType.HEAVY
}
func (i *MyItem) IsMagicArmor() bool {
	return i.ArmorType == armorType.MAGIC
}
func (i *MyItem) IsArmor() bool {
	return i.ItemType == ItemsPkg.ShieldOrArmor
}
func (i *MyItem) IsOnlyKamaelWeapon() bool {
	return i.WeaponType == weaponType.RAPIER || i.WeaponType == weaponType.CROSSBOW || i.WeaponType == weaponType.ANCIENTSWORD
}
func (i *MyItem) IsWeapon() bool {
	return i.ItemType == ItemsPkg.Weapon
}
func (i *MyItem) IsWeaponTypeNone() bool {
	return i.WeaponType == weaponType.NONE
}
func GetMyItems(charId int32) []MyItem {
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

	var myItems []MyItem

	for rows.Next() {
		var itm MyItem
		var id int

		err := rows.Scan(&itm.ObjId, &id, &itm.LocData, &itm.Enchant, &itm.Count, &itm.Loc, &itm.Time, &itm.Mana)
		if err != nil {
			panic(err)
		}

		it, ok := AllItems[id]
		if ok {
			itm.Item = it

			if itm.IsWeapon() {

				itm.AttackAttributeType, itm.AttackAttributeVal = getAttributeForWeapon(itm.ObjId)
			} else if itm.IsArmor() {
				itm.AttributeDefend = getAttributeForArmor(itm.ObjId)
			}

			myItems = append(myItems, itm)
		}
	}

	return myItems
}

func getAttributeForWeapon(objId int32) (attribute.Attribute, int) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()
	el := attribute.Attribute(-2) // None

	var elementType, elementValue int
	err = dbConn.QueryRow(context.Background(), "SELECT element_type,element_value FROM item_elementals WHERE object_id = $1", objId).
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

	rows, err := dbConn.Query(context.Background(), "SELECT element_type,element_value FROM item_elementals WHERE object_id = $1", objId)

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

func LoadItems() {
	AllItems = make(map[int]Item)
	loadItems()
}

func loadItems() {
	file, err := os.Open("./data/stats/items/items.json")
	if err != nil {
		panic("Failed to load config file")
	}

	var items []Item

	err = json.NewDecoder(file).Decode(&items)

	if err != nil {
		panic("Ошибка при чтении с файла items.json. " + err.Error())
	}

	for _, v := range items {
		v.removeEmptyStats()
		AllItems[v.Id] = v
	}

}

func (i *MyItem) IsEquipped() int16 {
	if i.Loc == Inventory {
		return 0
	}
	return 1
}

func SaveInventoryInDB(inventory []MyItem) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	for _, v := range inventory {
		//TODO sql в цикле надо переделать
		_, err = dbConn.Exec(context.Background(), "UPDATE items SET loc_data = $1, loc = $2 WHERE object_id = $3", v.LocData, v.Loc, v.ObjId)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (i *Item) removeEmptyStats() {
	var bStat []ItemsPkg.ItemBonusStat
	for _, v := range i.BonusStats {
		if v.Val != 0 {
			bStat = append(bStat, v)
		}
	}
	i.BonusStats = bStat
}

func GetActiveWeapon(inventory []MyItem, paperdoll [26]MyItem) *MyItem {
	q := paperdoll[PAPERDOLL_RHAND]
	for _, v := range inventory {
		if v.ObjId == q.ObjId {
			return &v
		}
	}
	return nil
}

// UseEquippableItem исользовать предмет который можно надеть на персонажа
func UseEquippableItem(selectedItem MyItem, character *Character) {
	//todo надо как то обновлять paperdoll, или возвращать массив или же  вынести это в другой пакет
	if selectedItem.IsEquipped() == 1 {
		unEquipAndRecord(selectedItem, character)
	} else {
		equipItemAndRecord(selectedItem, character)
	}
}

// unEquipAndRecord cнять предмет
func unEquipAndRecord(selectedItem MyItem, character *Character) {
	switch selectedItem.SlotBitType {
	case ItemsPkg.SlotLEar:
		setPaperdollItem(PAPERDOLL_LEAR, nil, character)
	case ItemsPkg.SlotREar:
		setPaperdollItem(PAPERDOLL_REAR, nil, character)
	case ItemsPkg.SlotNeck:
		setPaperdollItem(PAPERDOLL_NECK, nil, character)
	case ItemsPkg.SlotRFinger:
		setPaperdollItem(PAPERDOLL_RFINGER, nil, character)
	case ItemsPkg.SlotLFinger:
		setPaperdollItem(PAPERDOLL_LFINGER, nil, character)
	case ItemsPkg.SlotHair:
		setPaperdollItem(PAPERDOLL_HAIR, nil, character)
	case ItemsPkg.SlotHair2:
		setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
	case ItemsPkg.SlotHairall: //todo Разобраться что тут на l2j
		setPaperdollItem(PAPERDOLL_HAIR, nil, character)
	case ItemsPkg.SlotHead:
		setPaperdollItem(PAPERDOLL_HEAD, nil, character)
	case ItemsPkg.SlotRHand, ItemsPkg.SlotLrHand:
		setPaperdollItem(PAPERDOLL_RHAND, nil, character)
	case ItemsPkg.SlotLHand:
		setPaperdollItem(PAPERDOLL_LHAND, nil, character)
	case ItemsPkg.SlotGloves:
		setPaperdollItem(PAPERDOLL_GLOVES, nil, character)
	case ItemsPkg.SlotChest, ItemsPkg.SlotAlldress, ItemsPkg.SlotFullArmor:
		setPaperdollItem(PAPERDOLL_CHEST, nil, character)
	case ItemsPkg.SlotLegs:
		setPaperdollItem(PAPERDOLL_LEGS, nil, character)
	case ItemsPkg.SlotBack:
		setPaperdollItem(PAPERDOLL_CLOAK, nil, character)
	case ItemsPkg.SlotFeet:
		setPaperdollItem(PAPERDOLL_FEET, nil, character)
	case ItemsPkg.SlotUnderwear:
		setPaperdollItem(PAPERDOLL_UNDER, nil, character)
	case ItemsPkg.SlotLBracelet:
		setPaperdollItem(PAPERDOLL_LBRACELET, nil, character)
	case ItemsPkg.SlotRBracelet:
		setPaperdollItem(PAPERDOLL_RBRACELET, nil, character)
	case ItemsPkg.SlotDeco:
		setPaperdollItem(PAPERDOLL_DECO1, nil, character)
	case ItemsPkg.SlotBelt:
		setPaperdollItem(PAPERDOLL_BELT, nil, character)
	}
}

// equipItemAndRecord одеть предмет
func equipItemAndRecord(selectedItem MyItem, character *Character) {
	//todo проверка на приват Store, надо будет передавать character?
	// еще проверка на ITEM_CONDITIONS

	formal := character.Paperdoll[PAPERDOLL_CHEST]
	// Проверка надето ли офф. одежда и предмет не является букетом(id=21163)
	if (selectedItem.Id != 21163) && (formal.ObjId != 0) && (formal.SlotBitType == ItemsPkg.SlotAlldress) {
		// только chest можно
		switch selectedItem.SlotBitType {
		case ItemsPkg.SlotLrHand, ItemsPkg.SlotLHand, ItemsPkg.SlotRHand, ItemsPkg.SlotLegs, ItemsPkg.SlotFeet, ItemsPkg.SlotGloves, ItemsPkg.SlotHead:
			return
		}
	}

	paperdoll := character.Paperdoll

	switch selectedItem.SlotBitType {
	case ItemsPkg.SlotLrHand:
		setPaperdollItem(PAPERDOLL_LHAND, nil, character)
		setPaperdollItem(PAPERDOLL_RHAND, &selectedItem, character)
	case ItemsPkg.SlotLEar, ItemsPkg.SlotREar, ItemsPkg.SlotLrEar:
		if paperdoll[PAPERDOLL_LEAR].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_LEAR, &selectedItem, character)
		} else if paperdoll[PAPERDOLL_REAR].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_REAR, &selectedItem, character)
		} else {
			setPaperdollItem(PAPERDOLL_LEAR, &selectedItem, character)
		}

	case ItemsPkg.SlotNeck:
		setPaperdollItem(PAPERDOLL_NECK, &selectedItem, character)
	case ItemsPkg.SlotRFinger, ItemsPkg.SlotLFinger, ItemsPkg.SlotLrFinger:
		if paperdoll[PAPERDOLL_LFINGER].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_LFINGER, &selectedItem, character)
		} else if paperdoll[PAPERDOLL_RFINGER].ObjId == 0 {
			setPaperdollItem(PAPERDOLL_RFINGER, &selectedItem, character)
		} else {
			setPaperdollItem(PAPERDOLL_LFINGER, &selectedItem, character)
		}

	case ItemsPkg.SlotHair:
		hair := paperdoll[PAPERDOLL_HAIR]
		if hair.ObjId != 0 && hair.SlotBitType == ItemsPkg.SlotHairall {
			setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
		} else {
			setPaperdollItem(PAPERDOLL_HAIR, nil, character)
		}
		setPaperdollItem(PAPERDOLL_HAIR, &selectedItem, character)
	case ItemsPkg.SlotHair2:
		hair2 := paperdoll[PAPERDOLL_HAIR]
		if hair2.ObjId != 0 && hair2.SlotBitType == ItemsPkg.SlotHairall {
			setPaperdollItem(PAPERDOLL_HAIR, nil, character)
		} else {
			setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
		}
		setPaperdollItem(PAPERDOLL_HAIR2, &selectedItem, character)
	case ItemsPkg.SlotHairall:
		setPaperdollItem(PAPERDOLL_HAIR2, nil, character)
		setPaperdollItem(PAPERDOLL_HAIR, &selectedItem, character)
	case ItemsPkg.SlotHead:
		setPaperdollItem(PAPERDOLL_HEAD, &selectedItem, character)
	case ItemsPkg.SlotRHand:
		//todo снять стрелы
		setPaperdollItem(PAPERDOLL_RHAND, &selectedItem, character)
	case ItemsPkg.SlotLHand:
		rh := paperdoll[PAPERDOLL_RHAND]
		if (rh.ObjId != 0) && (rh.SlotBitType == ItemsPkg.SlotLrHand) && !(((rh.WeaponType == weaponType.BOW) && (selectedItem.EtcItemType == etcItemType.ARROW)) || ((rh.WeaponType == weaponType.CROSSBOW) && (selectedItem.EtcItemType == etcItemType.BOLT)) || ((rh.WeaponType == weaponType.FISHINGROD) && (selectedItem.EtcItemType == etcItemType.LURE))) {
			setPaperdollItem(PAPERDOLL_RHAND, nil, character)
		}
		setPaperdollItem(PAPERDOLL_LHAND, &selectedItem, character)
	case ItemsPkg.SlotGloves:
		setPaperdollItem(PAPERDOLL_GLOVES, &selectedItem, character)
	case ItemsPkg.SlotChest:
		setPaperdollItem(PAPERDOLL_CHEST, &selectedItem, character)
	case ItemsPkg.SlotLegs:
		chest := paperdoll[PAPERDOLL_CHEST]
		if chest.ObjId != 0 && chest.SlotBitType == ItemsPkg.SlotFullArmor {
			setPaperdollItem(PAPERDOLL_CHEST, nil, character)
		}
		setPaperdollItem(PAPERDOLL_LEGS, &selectedItem, character)
	case ItemsPkg.SlotBack:
		setPaperdollItem(PAPERDOLL_CLOAK, &selectedItem, character)
	case ItemsPkg.SlotFeet:
		setPaperdollItem(PAPERDOLL_FEET, &selectedItem, character)
	case ItemsPkg.SlotUnderwear:
		setPaperdollItem(PAPERDOLL_UNDER, &selectedItem, character)
	case ItemsPkg.SlotLBracelet:
		setPaperdollItem(PAPERDOLL_LBRACELET, &selectedItem, character)
	case ItemsPkg.SlotRBracelet:
		setPaperdollItem(PAPERDOLL_RBRACELET, &selectedItem, character)
	case ItemsPkg.SlotDeco:
		setPaperdollItem(PAPERDOLL_DECO1, &selectedItem, character)
	case ItemsPkg.SlotBelt:
		setPaperdollItem(PAPERDOLL_BELT, &selectedItem, character)
	case ItemsPkg.SlotFullArmor:
		setPaperdollItem(PAPERDOLL_LEGS, nil, character)
		setPaperdollItem(PAPERDOLL_CHEST, &selectedItem, character)
	case ItemsPkg.SlotAlldress:
		setPaperdollItem(PAPERDOLL_LEGS, nil, character)
		setPaperdollItem(PAPERDOLL_LHAND, nil, character)
		setPaperdollItem(PAPERDOLL_RHAND, nil, character)
		setPaperdollItem(PAPERDOLL_HEAD, nil, character)
		setPaperdollItem(PAPERDOLL_FEET, nil, character)
		setPaperdollItem(PAPERDOLL_GLOVES, nil, character)
		setPaperdollItem(PAPERDOLL_CHEST, &selectedItem, character)
	default:
		panic("Не определен Slot для itemId: " + strconv.Itoa(selectedItem.Id))
	}
}

func setPaperdollItem(slot uint8, selectedItem *MyItem, character *Character) {
	// eсли selectedItem nil, то ищем предмет которых находиться в slot
	// переносим его в инвентарь, убираем бонусы этого итема у персонажа
	if selectedItem == nil {
		for i, itemInInventory := range character.Inventory {
			if itemInInventory.LocData == int32(slot) {
				itemInInventory.LocData = getFirstEmptySlot(character.Inventory)
				itemInInventory.Loc = Inventory
				character.Inventory[i] = itemInInventory

				character.RemoveBonusStat(itemInInventory.BonusStats)
				break
			}
		}
		return
	}

	var oldItemInSelectedSlot MyItem
	var inventoryKeyOldItemInSelectedSlot int
	var keyCurrentItem int

	for i, v := range character.Inventory {
		// находим предмет, который стоял на нужном слоте раннее
		// его необходимо переместить в инвентарь
		if v.LocData == int32(slot) && v.Loc == Paperdoll {
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
	if oldItemInSelectedSlot.Id != 0 {
		oldItemInSelectedSlot.Loc = Inventory
		oldItemInSelectedSlot.LocData = selectedItem.LocData
		character.Inventory[inventoryKeyOldItemInSelectedSlot] = oldItemInSelectedSlot
		selectedItem.LocData = int32(slot)
		selectedItem.Loc = Paperdoll

		character.RemoveBonusStat(oldItemInSelectedSlot.BonusStats)
	} else {
		selectedItem.LocData = int32(slot)
		selectedItem.Loc = Paperdoll
	}
	// добавить бонусы предмета персонажу
	character.AddBonusStat(selectedItem.BonusStats)
	character.Inventory[keyCurrentItem] = *selectedItem

}

func getFirstEmptySlot(myItems []MyItem) int32 {
	var max int32
	for _, v := range myItems {
		if v.LocData > max {
			max = v.LocData
		}
	}

	var i int32
	for i = 0; i < max; i++ {
		flag := false
		for _, q := range myItems {
			if q.LocData == i && q.Loc != Paperdoll {
				flag = true
			}
		}

		if !flag {
			return i
		}
	}

	return max + 1
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
func DeleteItem(selectedItem MyItem, character *Character) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}

	if selectedItem.Loc == Paperdoll {
		character.Paperdoll[selectedItem.LocData] = MyItem{}
	}
	var itm []MyItem
	for _, v := range character.Inventory {
		if v.ObjId != selectedItem.ObjId {
			itm = append(itm, v)
		}
	}
	character.Inventory = itm

	_, _ = dbConn.Exec(context.Background(), "DELETE FROM items WHERE object_id = $1", selectedItem.ObjId)
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
