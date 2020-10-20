package items

import (
	"database/sql"
	"encoding/json"
	"github.com/jackc/pgx"
	"log"
	"os"
)

const (
	PAPERDOLL_UNDER     uint8 = 0
	PAPERDOLL_REAR      uint8 = 1
	PAPERDOLL_LEAR      uint8 = 2
	PAPERDOLL_NECK      uint8 = 3
	PAPERDOLL_RFINGER   uint8 = 4
	PAPERDOLL_LFINGER   uint8 = 5
	PAPERDOLL_HEAD      uint8 = 6
	PAPERDOLL_RHAND     uint8 = 7
	PAPERDOLL_LHAND     uint8 = 8
	PAPERDOLL_GLOVES    uint8 = 9
	PAPERDOLL_CHEST     uint8 = 10
	PAPERDOLL_LEGS      uint8 = 11
	PAPERDOLL_FEET      uint8 = 12
	PAPERDOLL_BACK      uint8 = 13
	PAPERDOLL_LRHAND    uint8 = 14
	PAPERDOLL_HAIR      uint8 = 15
	PAPERDOLL_DHAIR     uint8 = 16
	PAPERDOLL_RBRACELET uint8 = 17
	PAPERDOLL_LBRACELET uint8 = 18
	PAPERDOLL_DECO1     uint8 = 19
	PAPERDOLL_DECO2     uint8 = 20
	PAPERDOLL_DECO3     uint8 = 21
	PAPERDOLL_DECO4     uint8 = 22
	PAPERDOLL_DECO5     uint8 = 23
	PAPERDOLL_DECO6     uint8 = 24
	PAPERDOLL_BELT      uint8 = 25
	PaperdollMax        uint8 = 26
)

type ItemFromDb struct {
	ObjectId     int32
	Item         int32
	EnchantLevel int32
	LocData      sql.NullInt32
}

type Inventory struct {
}

func RestoreVisibleInventory(charId int32, db *pgx.Conn) [PaperdollMax][3]int32 {
	var kek [PaperdollMax][3]int32

	rows, err := db.Query("SELECT object_id,item,loc_data,enchant_level FROM items WHERE owner_id=$1 AND loc=$2", charId, "PAPERDOLL")
	if err != nil {
		log.Fatal(err)
	}
	type Items []ItemFromDb
	for rows.Next() {
		var objId int
		var Item int
		var EnchantLevel int
		var LocData int
		err := rows.Scan(&objId, &Item, &LocData, &EnchantLevel)
		if err != nil {
			log.Println(err)
		}
		kek[int32(LocData)][0] = int32(objId)
		kek[int32(LocData)][1] = int32(Item)
		kek[int32(LocData)][2] = int32(EnchantLevel)
	}
	return kek
}

type CrystalType struct {
	Id                        int
	CrystalId                 int
	CrystalEnchantBonusArmor  int
	CrystalEnchantBonusWeapon int32
}
type itemsJson struct {
	Id              int
	ObjId           int32
	Loc             int32
	Count           int64
	Name            string
	Icon            string
	Type            string
	WeaponType      string `json:"weapon_type"`
	Bodypart        string
	AttackRange     int    `json:"attack_range"`
	DamageRange     string `json:"damage_range"`
	ImmediateEffect bool   `json:"immediate_effect"`
	Weight          int
	Material        string
	Price           int
	Soulshots       int
	Spiritshots     int
	PAtk            int
	MAtk            int
	CritRate        int
	PAtkSpd         int
}

type Item struct {
	Id              int32
	ObjId           int32
	Loc             string
	LocData         int32
	Count           int64
	Name            string
	Icon            string
	Type            string
	WeaponType      string
	Bodypart        int32
	ItemType        int16
	AttackRange     int
	DamageRange     string
	ImmediateEffect bool
	Weight          int
	Material        string
	Price           int
	Soulshots       int
	Spiritshots     int
	PAtk            int
	MAtk            int
	CritRate        int
	PAtkSpd         int
}

var AllJsonItems []itemsJson

var AllItems []Item

type MyItem struct {
	Id      int32
	ObjId   int32
	Name    int32
	Loc     int32
	Count   int32
	Enchant int32
	Mana    int32
	Time    int32
}

func GetMyItems(charId int32, db *pgx.Conn) []Item {
	rows, err := db.Query("SELECT object_id,item,loc_data,enchant_level,count,loc FROM items WHERE owner_id=$1", charId)
	if err != nil {
		log.Fatal(err)
	}

	type f struct {
		objId   int
		Item    int
		Enchant int
		LocData int
		Count   int
		Loc     string
	}
	var t []f
	for rows.Next() {
		var q f
		err := rows.Scan(&q.objId, &q.Item, &q.LocData, &q.Enchant, &q.Count, &q.Loc)
		if err != nil {
			log.Println(err)
		}
		t = append(t, q)
	}

	var myItems []Item
	for _, w := range AllItems {
		for _, q := range t {
			if w.Id == int32(q.Item) {
				ni := new(Item)
				ni = &w
				ni.ObjId = int32(q.objId)
				ni.LocData = int32(q.LocData)
				ni.Count = int64(q.Count)
				ni.Loc = q.Loc
				myItems = append(myItems, *ni)

			}
		}
	}
	e := AllItems
	_ = e
	return myItems
}

func LoadItems() {
	file, err := os.Open("./data/stats/items/weapon.json")
	if err != nil {
		log.Fatal("Failed to load config file")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AllJsonItems)
	if err != nil {
		log.Fatal("Failed to decode config file")
	}
	SetSlots()
	for _, v := range AllJsonItems {
		weapon := new(Item)
		weapon.Id = int32(v.Id)
		weapon.Loc = ""
		weapon.Bodypart = getSlots(v.Bodypart)
		weapon.ItemType = 0 //weapon
		weapon.Name = v.Name
		weapon.Icon = v.Icon
		weapon.AttackRange = v.AttackRange
		weapon.CritRate = v.CritRate
		weapon.DamageRange = v.DamageRange
		weapon.ImmediateEffect = v.ImmediateEffect
		weapon.MAtk = v.MAtk
		weapon.PAtk = v.PAtk
		AllItems = append(AllItems, *weapon)
	}

}

var Slots map[string]int32

func SetSlots() {
	slots := make(map[string]int32)
	Slots = slots
	Slots["shirt"] = 0x0001
	Slots["lbracelet"] = 0x200000
	Slots["rbracelet"] = 0x100000
	Slots["talisman"] = 0x400000
	Slots["chest"] = 0x0400
	Slots["fullarmor"] = 0x8000
	Slots["head"] = 0x0040
	Slots["hair"] = 0x010000
	Slots["hairall"] = 0x080000
	Slots["underwear"] = 0x0001
	Slots["back"] = 0x2000
	Slots["neck"] = 0x0008
	Slots["legs"] = 0x0800
	Slots["feet"] = 0x1000
	Slots["gloves"] = 0x0200
	Slots["chest,legs"] = 0x0400 | 0x0800
	Slots["belt"] = 0x10000000
	Slots["rhand"] = 0x0080
	Slots["lhand"] = 0x0100
	Slots["lrhand"] = 0x4000
	Slots["rear;lear"] = 0x0002 | 0x0004
	Slots["rfinger;lfinger"] = 0x0010 | 0x0020
	Slots["wolf"] = -100
	Slots["greatwolf"] = -104
	Slots["hatchling"] = -101
	Slots["strider"] = -102
	Slots["babypet"] = -103
	Slots["none"] = 0x0000

	// retail compatibility
	Slots["onepiece"] = 0x8000
	Slots["hair2"] = 0x040000
	Slots["dhair"] = 0x080000
	Slots["alldress"] = 0x020000
	Slots["deco1"] = 0x400000
	Slots["waist"] = 0x10000000

}
func getSlots(s string) int32 {
	return Slots[s]
}

func (i *Item) IsEquipped() int16 {
	if i.Loc == "INVENTORY" {
		return 0
	}
	return 1
}
