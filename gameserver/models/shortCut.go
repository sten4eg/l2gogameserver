package models

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/dto"
)

const MaxShortcutsPerBar = 12

func RegisterShortCut(sc dto.ShortCutDTO, client *ClientCtx) {
	shorts := client.CurrentChar.ShortCut
	//todo пересмотреть, тут есть еще проверки
	s, exist := shorts[sc.Slot+(sc.Page*MaxShortcutsPerBar)]

	if exist {
		deleteShortCutFromDb(s, client.CurrentChar.ObjectId, client.CurrentChar.ClassId, client.db)
	}
	registerShortCutInDb(sc, client.CurrentChar.ObjectId, client.CurrentChar.ClassId, client.db)
	client.CurrentChar.ShortCut[sc.Slot+(sc.Page*MaxShortcutsPerBar)] = sc
}

func registerShortCutInDb(shortCut dto.ShortCutDTO, charId, classId int32, db *sql.DB) {

	_, err := db.Exec("INSERT INTO character_shortcuts (char_id, slot, page, type,shortcut_id, level, class_index) VALUES($1,$2,$3,$4,$5,$6,$7)",
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

func RestoreMe(charId, classId int32, db *sql.DB) map[int32]dto.ShortCutDTO {

	shorts := make(map[int32]dto.ShortCutDTO)
	rows, err := db.Query("SELECT slot, page, type, shortcut_id, level FROM character_shortcuts WHERE char_id = $1 AND class_index = $2", charId, classId)
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer rows.Close()
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

func DeleteShortCut(slot, page int32, client *ClientCtx) {
	all := client.CurrentChar.ShortCut
	e, ok := all[slot+(page*MaxShortcutsPerBar)]
	if !ok {
		return
	}
	deleteShortCutFromDb(e, client.CurrentChar.ObjectId, client.CurrentChar.ClassId, client.db)
	// todo Проверка на соски

}

func deleteShortCutFromDb(shortCut dto.ShortCutDTO, charId int32, classId int32, db *sql.DB) {
	_, err := db.Exec("DELETE FROM character_shortcuts WHERE char_id=$1 AND slot=$2 AND page=$3 AND class_index=$4",
		charId,
		shortCut.Slot,
		shortCut.Page,
		classId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}
