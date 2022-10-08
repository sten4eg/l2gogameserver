package dto

import (
	"context"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
)

type ShortCutDTO struct {
	// Slot from 0 to 11
	Slot int32

	// Page from 0 to 9
	Page int32

	// ShortcutType: item, skill, action, macro, recipe, bookmark.
	// Values from ShortTypes
	ShortcutType string

	// Id shortcut
	Id int32

	// Level shortcut level (skills)
	Level int32

	// CharacterType: 1 player, 2 summon
	CharacterType int32

	// SharedReuseGroup ???
	SharedReuseGroup int32
}

type ShortCutSimpleDTO struct {
	Slot         int32
	Page         int32
	ShortcutType int32
	Id           int32
	Level        int32
}

func GetShortCutDTO(slot, page, id, level, characterType int32, shortType string) ShortCutDTO {
	var shDto ShortCutDTO

	shDto.Slot = slot
	shDto.Page = page
	shDto.Id = id
	shDto.Level = level
	shDto.CharacterType = characterType
	shDto.ShortcutType = shortType

	shDto.SharedReuseGroup = -1

	return shDto
}

func IndexOfShortTypes(shortType string) int32 {
	for k, v := range ShortTypes {
		if shortType == v {
			return int32(k)
		}
	}
	logger.Error.Panicln("ShortType: " + shortType + " не найден")
	return 0
}

var ShortTypes = [7]string{
	"NONE",
	"ITEM",
	"SKILL",
	"ACTION",
	"MACRO",
	"RECIPE",
	"BOOKMARK",
}

func GetAllShortCuts(charId, classId int32) []ShortCutSimpleDTO {
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
	var shortCuts []ShortCutSimpleDTO
	for rows.Next() {
		var t ShortCutSimpleDTO
		err = rows.Scan(&t.ShortcutType, &t.Slot, &t.Page, &t.Id, &t.Level)
		if err != nil {
			logger.Error.Panicln(err)
		}
		shortCuts = append(shortCuts, t)
	}
	return shortCuts
}
