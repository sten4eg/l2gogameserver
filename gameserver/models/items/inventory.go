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
	Enchant         int16
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
				ni.Enchant = int16(q.Enchant)
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
	if i.Loc == "INVENTORY" {
		return 0
	}
	return 1
}
