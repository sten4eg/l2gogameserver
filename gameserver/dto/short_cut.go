package dto

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
