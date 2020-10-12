package items

import (
	"database/sql"
	"github.com/jackc/pgx"
	"log"
)

const (
	PAPERDOLL_UNDER     int32 = 0
	PAPERDOLL_REAR      int32 = 1
	PAPERDOLL_LEAR      int32 = 2
	PAPERDOLL_NECK      int32 = 3
	PAPERDOLL_RFINGER   int32 = 4
	PAPERDOLL_LFINGER   int32 = 5
	PAPERDOLL_HEAD            = 6
	PAPERDOLL_RHAND           = 7
	PAPERDOLL_LHAND           = 8
	PAPERDOLL_GLOVES          = 9
	PAPERDOLL_CHEST           = 10
	PAPERDOLL_LEGS            = 11
	PAPERDOLL_FEET            = 12
	PAPERDOLL_BACK            = 13
	PAPERDOLL_LRHAND          = 14
	PAPERDOLL_HAIR            = 15
	PAPERDOLL_DHAIR           = 16
	PAPERDOLL_RBRACELET       = 17
	PAPERDOLL_LBRACELET       = 18
	PAPERDOLL_DECO1           = 19
	PAPERDOLL_DECO2           = 20
	PAPERDOLL_DECO3           = 21
	PAPERDOLL_DECO4           = 22
	PAPERDOLL_DECO5           = 23
	PAPERDOLL_DECO6           = 24
	PAPERDOLL_BELT            = 25
	PaperdollMax        int32 = 26
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
