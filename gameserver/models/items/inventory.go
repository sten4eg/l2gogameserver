package items

import (
	"context"
	"encoding/json"
	"l2gogameserver/db"
	"log"
	"os"
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
)

func RestoreVisibleInventory(charId int32) [31][3]int32 {
	var paperdoll [31][3]int32

	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT object_id, item, loc_data, enchant_level FROM items WHERE owner_id= $1 AND loc= $2", charId, Paperdoll)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var objId int
		var item int
		var enchantLevel int
		var locData int
		err := rows.Scan(&objId, &item, &locData, &enchantLevel)
		if err != nil {
			log.Println(err)
		}
		paperdoll[int32(locData)][0] = int32(objId)
		paperdoll[int32(locData)][1] = int32(item)
		paperdoll[int32(locData)][2] = int32(enchantLevel)
	}
	return paperdoll
}

type Item struct {
	Id                     int
	ItemType               string
	Name                   string
	Icon                   string
	SlotBitType            ItemType `json:"slot_bit_type"`
	ArmorType              string
	EtcItemType            string
	ItemMultiSkillList     string
	RecipeId               int
	Weight                 int
	ConsumeType            string
	SoulShotCount          int
	SpiritShotCount        int
	DropPeriod             int
	Duration               int
	Period                 int
	DefaultPrice           int
	ItemSkill              string
	CriticalAttackSkill    string
	AttackSkill            string
	MagicSkill             string
	ItemSkillEnchantedFour int
	MaterialType           string
	CrystalType            string
	CrystalCount           int
	IsTrade                bool
	IsDrop                 bool
	IsDestruct             bool
	IsPrivateStore         bool
	KeepType               int
	PhysicalDamage         int
	RandomDamage           int
	WeaponType             string
	Critical               int
	HitModify              int
	AvoidModify            int
	ShieldDefense          int
	ShieldDefenseRate      int
	AttackRange            int
	AttackSpeed            int
	ReuseDelay             int
	MpConsume              int
	MagicalDamage          int
	Durability             int
	PhysicalDefence        int
	MagicalDefence         int
	MpBonus                int
	MagicWeapon            bool
	EnchantEnable          bool
	ElementalEnable        bool
	ForNpc                 bool
	IsOlympiadCanUse       bool
	IsPremium              bool
}

// AllItems - ONLY READ MAP, set in init server
var AllItems map[int]Item

func GetMyItems(charId int32) []Item {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	rows, err := dbConn.Query(context.Background(), "SELECT object_id,item,loc_data,enchant_level,count,loc FROM items WHERE owner_id=$1", charId)
	if err != nil {
		panic(err)
	}

	type tempItemFromDB struct {
		objId   int
		Item    int
		Enchant int
		LocData int
		Count   int
		Loc     string
	}

	var tmp []tempItemFromDB

	for rows.Next() {
		var itm tempItemFromDB
		err := rows.Scan(&itm.objId, &itm.Item, &itm.LocData, &itm.Enchant, &itm.Count, &itm.Loc)
		if err != nil {
			log.Println(err)
		}
		tmp = append(tmp, itm)
	}

	var myItems []Item

	//for _, itemFromDB := range tmp {
	//	_, ok := AllItems[int32(itemFromDB.Item)]
	//	if ok {
	//		myItem := AllItems[int32(itemFromDB.Item)]
	//		myItem.Id = int32(itemFromDB.Item)
	//		myItem.ObjId = int32(itemFromDB.objId)
	//		myItem.LocData = int32(itemFromDB.LocData)
	//		myItem.Count = int64(itemFromDB.Count)
	//		myItem.Loc = itemFromDB.Loc
	//		myItem.Enchant = int16(itemFromDB.Enchant)
	//
	//		myItems = append(myItems, myItem)
	//	}
	//}

	return myItems
}

func LoadItems() {
	SetSlots()
	AllItems = make(map[int]Item)

	//loadWeapons()
	//loadArmors()
	loadItems()
}

//func loadArmors() {
//	file, err := os.Open("./data/stats/items/armor.json")
//	if err != nil {
//		panic("Failed to load config file")
//	}
//
//	decoder := json.NewDecoder(file)
//
//	var armorsJson []armorJson
//
//	err = decoder.Decode(&armorsJson)
//	if err != nil {
//		panic("Failed to decode config file")
//	}
//
//	for _, v := range armorsJson {
//		armor := new(Item)
//		armor.Loc = ""
//		armor.Bodypart = getSlots(v.Bodypart)
//		armor.ItemType = 1 // armor/shield
//		armor.Name = v.Name
//		AllItems[int32(v.Id)] = *armor
//	}
//}
//
//func loadWeapons() {
//	file, err := os.Open("./data/stats/items/weapon.json")
//	if err != nil {
//		panic("Failed to load config file")
//	}
//
//	decoder := json.NewDecoder(file)
//
//	var weaponJson []weaponJson
//
//	err = decoder.Decode(&weaponJson)
//	if err != nil {
//		panic("Failed to decode config file")
//	}
//
//	for _, v := range weaponJson {
//		weapon := new(Item)
//		weapon.Loc = ""
//		weapon.Bodypart = getSlots(v.Bodypart)
//		weapon.ItemType = 0 //weapon
//		weapon.Name = v.Name
//		weapon.Icon = v.Icon
//		weapon.AttackRange = v.AttackRange
//		weapon.CritRate = v.CritRate
//		weapon.DamageRange = v.DamageRange
//		weapon.ImmediateEffect = v.ImmediateEffect
//		weapon.MAtk = v.MAtk
//		weapon.PAtk = v.PAtk
//		AllItems[int32(v.Id)] = *weapon
//	}
//}

func loadItems() {
	file, err := os.Open("./data/stats/items/items.json")
	if err != nil {
		panic("Failed to load config file")
	}

	decoder := json.NewDecoder(file)

	var items []Item

	err = decoder.Decode(&items)
	if err != nil {
		panic("Failed to decode config file")
	}
	//for _, v := range otherJson {
	//	weapon := new(Item)
	//	weapon.Loc = ""
	//	weapon.Bodypart = 0
	//	weapon.ItemType = 05 //item
	//	weapon.Name = v.Name
	//	weapon.Icon = v.Icon
	////	weapon.ImmediateEffect = v.ImmediateEffect
	//	AllItems[int32(v.Id)] = *weapon
	//}

}

var Slots map[string]int32

func SetSlots() {
	slots := make(map[string]int32)
	Slots = slots
	Slots["shirt"] = SlotUnderwear
	Slots["lbracelet"] = SlotLBracelet
	Slots["rbracelet"] = SlotRBracelet
	Slots["talisman"] = SlotDeco
	Slots["chest"] = SlotChest
	Slots["fullarmor"] = SlotFullArmor
	Slots["head"] = SlotHead
	Slots["hair"] = SlotHair
	Slots["hairall"] = SlotHairall
	Slots["underwear"] = SlotUnderwear
	Slots["back"] = SlotBack
	Slots["neck"] = SlotNeck
	Slots["legs"] = SlotLegs
	Slots["feet"] = SlotFeet
	Slots["gloves"] = SlotGloves
	Slots["chest,legs"] = SlotChest | SlotLegs
	Slots["belt"] = SlotBelt
	Slots["rhand"] = SlotRHand
	Slots["lhand"] = SlotLHand
	Slots["lrhand"] = SlotLrHand
	Slots["rear;lear"] = SlotREar | SlotLEar
	Slots["rfinger;lfinger"] = SlotRFinger | SlotLFinger
	Slots["wolf"] = SlotWolf
	Slots["greatwolf"] = SlotGreatwolf
	Slots["hatchling"] = SlotHatchling
	Slots["strider"] = SlotStrider
	Slots["babypet"] = SlotBabypet
	Slots["none"] = SlotNone

	// retail compatibility
	Slots["onepiece"] = SlotFullArmor
	Slots["hair2"] = SlotHair2
	Slots["dhair"] = SlotHairall
	Slots["alldress"] = SlotAlldress
	Slots["deco1"] = SlotDeco
	Slots["waist"] = SlotBelt

}
func getSlots(s string) int32 {
	return Slots[s]
}

func (i *Item) IsEquipped() int16 {
	//if i.Loc == Inventory {
	//	return 0
	//}
	return 1
}

func SaveInventoryInDB(inventory []Item) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	for _, _ = range inventory {
		//TODO sql в цикле надо переделать
		//		_, err = dbConn.Exec(context.Background(), "UPDATE items SET loc_data = $1, loc = $2 WHERE object_id = $3", v.LocData, v.Loc, v.ObjId)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func GetMyItemByObjId(charId int32, objId int32) Item {
	dbConn, err := db.GetConn()
	if err != nil {
		return Item{}
	}
	defer dbConn.Release()

	items := GetMyItems(charId)

	for _, _ = range items {
		//if v.ObjId == objId {
		//	return v
		//}
	}
	return Item{}
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
