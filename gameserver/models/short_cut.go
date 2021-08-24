package models

import (
	"context"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/dto"
	"strconv"
)

const MaxShortcutsPerBar = 12

func RegisterShortCut(sc dto.ShortCutDTO, charId, classId int32) {
	registerShortCutInDb(sc, charId, classId)
}

func registerShortCutInDb(shortCut dto.ShortCutDTO, charId, classId int32) {
	dbConn, err := db.GetConn()
	if err != nil {
		return
	}
	defer dbConn.Release()

	_, err = dbConn.Exec(context.Background(), "INSERT INTO character_shortcuts (char_id, slot, page, type, level, class_index) VALUES($1,$2,$3,$4,$5,$6)",
		charId,
		shortCut.Slot,
		shortCut.Page,
		dto.IndexOfShortTypes(shortCut.ShortcutType),
		strconv.Itoa(int(shortCut.Level)),
		classId)
	if err != nil {
		return
	}
}
