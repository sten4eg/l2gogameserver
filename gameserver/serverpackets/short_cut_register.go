package serverpackets

import (
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ShortCutRegister(shortCut dto.ShortCutDTO, client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x44)
	buffer.WriteD(dto.IndexOfShortTypes(shortCut.ShortcutType))

	buffer.WriteD(shortCut.Slot + (shortCut.Page * 12)) //C4 Client ?????
	buffer.WriteD(shortCut.Id)

	switch shortCut.ShortcutType {
	case "ITEM":
		buffer.WriteD(shortCut.CharacterType)
		buffer.WriteD(shortCut.SharedReuseGroup)
		buffer.WriteD(0) // unknown
		buffer.WriteD(0) // unknown
		buffer.WriteD(0) // item augment id
	case "SKILL":
		buffer.WriteD(shortCut.Level)
		buffer.WriteSingleByte(0)
		buffer.WriteD(shortCut.CharacterType)
	case "ACTION", "MACRO", "RECIPE", "BOOKMARK":
		buffer.WriteD(shortCut.CharacterType)
	}

	return buffer.Bytes()
}
