package items

import (
	"context"
	"encoding/json"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/models/items/armorType"
	"l2gogameserver/gameserver/models/items/weaponType"
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

	UpdateTypeAdd    int16 = 1
	UpdateTypeModify int16 = 2
	UpdateTypeRemove int16 = 3
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
	ItemType               ItemType
	Name                   string
	Icon                   string
	SlotBitType            SlotBitType         `json:"slot_bit_type"`
	ArmorType              armorType.ArmorType `json:"armor_type"`
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
	WeaponType             weaponType.WeaponType `json:"weapon_type"`
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

type MyItem struct {
	Item
	ObjId   int32
	Enchant int
	LocData int32
	Count   int64
	Loc     string
}

// IsEquipable Можно ли надеть предмет
func (i *MyItem) IsEquipable() bool {
	//todo сделать 'enum' для EtcItemType
	return !((i.SlotBitType == SlotNone) || (i.EtcItemType == "ARROW") || (i.EtcItemType == "BOLT") || (i.EtcItemType == "LURE"))
}

func GetMyItems(charId int32) []MyItem {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT object_id,item,loc_data,enchant_level,count,loc FROM items WHERE owner_id=$1", charId)
	if err != nil {
		panic(err)
	}

	var myItems []MyItem

	for rows.Next() {
		var itm MyItem
		var id int
		err := rows.Scan(&itm.ObjId, &id, &itm.LocData, &itm.Enchant, &itm.Count, &itm.Loc)
		if err != nil {
			log.Println(err)
		}
		it, ok := AllItems[id]
		if ok {
			itm.Item = it
			myItems = append(myItems, itm)
		}

	}

	return myItems
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

//func GetMyItemByObjId(character *models.Character, objId int32) MyItem {
//	items := character.Inventory
//
//	for _, v := range items {
//		if v.ObjId == objId {
//			return v
//		}
//	}
//	return MyItem{}
//}

func GetActiveWeapon(inventory []MyItem, paperdoll [31][3]int32) *MyItem {
	q := paperdoll[PAPERDOLL_RHAND][0]
	for _, v := range inventory {
		if v.ObjId == q {
			return &v
		}
	}
	return nil
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
