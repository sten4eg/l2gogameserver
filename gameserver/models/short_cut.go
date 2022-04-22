package models

import (
	"context"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/dto"
)

const MaxShortcutsPerBar = 12

func RegisterShortCut(sc dto.ShortCutDTO, client *Client) {
	shorts := client.CurrentChar.ShortCut
	//todo пересмотреть, тут есть еще проверки
	s, exist := shorts[sc.Slot+(sc.Page*MaxShortcutsPerBar)]

	if exist {
		deleteShortCutFromDb(s, client.CurrentChar.ObjectId, client.CurrentChar.ClassId)
	}
	registerShortCutInDb(sc, client.CurrentChar.ObjectId, client.CurrentChar.ClassId)
	client.CurrentChar.ShortCut[sc.Slot+(sc.Page*MaxShortcutsPerBar)] = sc
}

func registerShortCutInDb(shortCut dto.ShortCutDTO, charId, classId int32) {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	_, err = dbConn.Exec(context.Background(), "INSERT INTO character_shortcuts (char_id, slot, page, type,shortcut_id, level, class_index) VALUES($1,$2,$3,$4,$5,$6,$7)",
		charId,
		shortCut.Slot,
		shortCut.Page,
		dto.IndexOfShortTypes(shortCut.ShortcutType),
		shortCut.Id,
		shortCut.Level,
		classId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}

func RestoreMe(charId, classId int32) map[int32]dto.ShortCutDTO {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	shorts := make(map[int32]dto.ShortCutDTO)
	rows, err := dbConn.Query(context.Background(), "SELECT slot, page, type, shortcut_id, level FROM character_shortcuts WHERE char_id = $1 AND class_index = $2", charId, classId)
	if err != nil {
		logger.Error.Panicln(err)
	}

	for rows.Next() {
		var t dto.ShortCutDTO
		var shortType int
		err = rows.Scan(&t.Slot, &t.Page, &shortType, &t.Id, &t.Level)
		if err != nil {
			logger.Error.Panicln(err)
		}
		t.ShortcutType = dto.ShortTypes[shortType]
		shorts[t.Slot+(t.Page*MaxShortcutsPerBar)] = t
	}

	//Verify shortcuts todo need release

	return shorts
}

func GetAllShortCuts(charId, classId int32) []dto.ShortCutSimpleDTO {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT type, slot, page,shortcut_id, level FROM character_shortcuts WHERE char_id = $1 AND class_index = $2",
		charId,
		classId)
	if err != nil {
		logger.Error.Panicln(err)
	}
	var shortCuts []dto.ShortCutSimpleDTO
	for rows.Next() {
		var t dto.ShortCutSimpleDTO
		err = rows.Scan(&t.ShortcutType, &t.Slot, &t.Page, &t.Id, &t.Level)
		if err != nil {
			logger.Error.Panicln(err)
		}
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
	deleteShortCutFromDb(e, client.CurrentChar.ObjectId, client.CurrentChar.ClassId)
	// todo Проверка на соски

}

func deleteShortCutFromDb(shortCut dto.ShortCutDTO, charId int32, classId int32) {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	_, err = dbConn.Exec(context.Background(), "DELETE FROM character_shortcuts WHERE char_id=$1 AND slot=$2 AND page=$3 AND class_index=$4",
		charId,
		shortCut.Slot,
		shortCut.Page,
		classId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}
