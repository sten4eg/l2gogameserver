package models

import (
	"context"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/dto"
	"strconv"
)

const MaxShortcutsPerBar = 12

func RegisterShortCut(sc dto.ShortCutDTO, client *Client) {
	shorts := client.CurrentChar.ShortCut

	s, exist := shorts[sc.Slot+(sc.Page*MaxShortcutsPerBar)]
	if exist {
		deleteShortCutFromDb(s, client.CurrentChar.CharId, client.CurrentChar.ClassId)
	}
	registerShortCutInDb(sc, client.CurrentChar.CharId, client.CurrentChar.ClassId)
}

func deleteShortCutFromDb(shortCut dto.ShortCutDTO, charId int32, classId int32) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	_, err = dbConn.Exec(context.Background(), "DELETE FROM character_shortcuts WHERE char_id=$1 AND slot=$2 AND page=$3 AND class_index=$4",
		charId,
		shortCut.Slot,
		shortCut.Page,
		classId)
	if err != nil {
		panic(err)
	}
}

func registerShortCutInDb(shortCut dto.ShortCutDTO, charId, classId int32) {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	_, err = dbConn.Exec(context.Background(), "INSERT INTO character_shortcuts (char_id, slot, page, type,shortcut_id, level, class_index) VALUES($1,$2,$3,$4,$5,$6,$7)",
		charId,
		shortCut.Slot,
		shortCut.Page,
		dto.IndexOfShortTypes(shortCut.ShortcutType),
		shortCut.Id,
		strconv.Itoa(int(shortCut.Level)),
		classId)
	if err != nil {
		panic(err)
	}
}

func restoreMe(charId, classId int32) map[int32]dto.ShortCutDTO {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	shorts := make(map[int32]dto.ShortCutDTO)
	rows, err := dbConn.Query(context.Background(), "SELECT slot, page, type, shortcut_id, level FROM character_shortcuts WHERE char_id = $1 AND class_index = $2", charId, classId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var t dto.ShortCutDTO
		var shortType int
		var lvl string
		err = rows.Scan(&t.Slot, &t.Page, &shortType, &t.Id, &lvl)
		if err != nil {
			panic(err)
		}
		t.ShortcutType = dto.ShortTypes[shortType]
		i, err := strconv.Atoi(lvl)
		if err != nil {
			panic(err)
		}
		t.Level = int32(i)
		shorts[t.Slot+(t.Page*MaxShortcutsPerBar)] = t
	}

	//Verify shortcuts todo need release

	return shorts
}

func GetAllShortCuts(charId, classId int32) []dto.ShortCutSimpleDTO {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT type, slot, page,shortcut_id, level FROM character_shortcuts WHERE char_id = $1 AND class_index = $2",
		charId,
		classId)
	if err != nil {
		panic(err)
	}
	var shortCuts []dto.ShortCutSimpleDTO
	for rows.Next() {
		var t dto.ShortCutSimpleDTO
		var lvl string
		err = rows.Scan(&t.ShortcutType, &t.Slot, &t.Page, &t.Id, &lvl)
		if err != nil {
			panic(err)
		}
		ilvl, err := strconv.Atoi(lvl)
		if err != nil {
			panic(err)
		}
		t.Level = int32(ilvl)
		shortCuts = append(shortCuts, t)
	}
	return shortCuts
}

func DeleteShortCut(slot, page int32, client *Client) {
	all := client.CurrentChar.ShortCut
	e, ok := all[slot+(page*MaxShortcutsPerBar)]
	if !ok {
		return
	}
	deleteShortCutFromDb(e, client.CurrentChar.CharId, client.CurrentChar.ClassId)
	// todo Проверка на соски

}
