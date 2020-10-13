package items

import (
	"database/sql"
	"github.com/jackc/pgx"
	"log"
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

type Item struct {
	ObjectId     int32
	Item         int32
	EnchantLevel int32
	LocData      sql.NullInt32
}

func RestoreVisibleInventory(charId int32, db *pgx.Conn) [PaperdollMax][3]int32 {
	var kek [PaperdollMax][3]int32

	rows, err := db.Query("SELECT object_id,item,loc_data,enchant_level FROM items WHERE owner_id=$1 AND loc=$2", charId, "PAPERDOLL")
	if err != nil {
		log.Fatal(err)
	}
	type Items []Item
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