package gs2ls

import (
	"database/sql"
	"l2gogameserver/packets"
	"log"
)

const CharsCount = `SELECT COUNT(object_id) FROM characters WHERE login = $1`

func ReplyCharacters(login string, db *sql.DB) *packets.Buffer {
	var timeToDel []int64

	rows, err := db.Query(`SELECT delete_in FROM characters WHERE login = $1 AND delete_in IS NOT NULL`, login)
	if err != nil {
		log.Println("err", err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var i int64
		err = rows.Scan(&i)
		if err != nil {
			log.Println(err)
			return nil
		}
		timeToDel = append(timeToDel, i)
	}

	var charCount byte
	err = db.QueryRow(CharsCount, login).Scan(&charCount)
	if err != nil {
		log.Println("err", err.Error())
		return nil
	}

	buf := packets.Get()
	buf.WriteSingleByte(0x08)
	buf.WriteS(login)
	buf.WriteSingleByte(charCount)
	buf.WriteSingleByte(byte(len(timeToDel)))
	for i := range timeToDel {
		buf.WriteQ(timeToDel[i])
	}

	return buf
}
