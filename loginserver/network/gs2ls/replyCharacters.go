package gs2ls

import (
	"context"
	"l2gogameserver/db"
	"l2gogameserver/packets"
	"log"
)

const CharsCount = `SELECT COUNT(object_id) FROM characters WHERE login = $1`

func ReplyCharacters(login string) *packets.Buffer {
	dbConn, err := db.GetConn()
	if err != nil {
		return nil
	}
	defer dbConn.Release()

	var timeToDel []int64

	rows, err := dbConn.Query(context.Background(), `SELECT delete_in FROM characters WHERE login = $1 AND delete_in IS NOT NULL`, login)
	if err != nil {
		log.Println("err", err.Error())
		return nil
	}

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
	err = dbConn.QueryRow(context.Background(), CharsCount, login).Scan(&charCount)
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
