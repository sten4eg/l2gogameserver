package serverpackets

import (
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models"
)

func ShortCutRegister(shortCut dto.ShortCutDTO, client *models.Client) {

	client.Buffer.WriteSingleByte(0x44)
	client.Buffer.WriteD(dto.IndexOfShortTypes(shortCut.ShortcutType))

	client.Buffer.WriteD(shortCut.Slot + (shortCut.Page * 12)) //C4 Client ?????
	client.Buffer.WriteD(shortCut.Id)

	switch shortCut.ShortcutType {
	case "ITEM":
		client.Buffer.WriteD(shortCut.CharacterType)
		client.Buffer.WriteD(shortCut.SharedReuseGroup)
		client.Buffer.WriteD(0) // unknown
		client.Buffer.WriteD(0) // unknown
		client.Buffer.WriteD(0) // item augment id
	case "SKILL":
		client.Buffer.WriteD(shortCut.Level)
		client.Buffer.WriteSingleByte(0)
		client.Buffer.WriteD(shortCut.CharacterType)
	case "ACTION", "MACRO", "RECIPE", "BOOKMARK":
		client.Buffer.WriteD(shortCut.CharacterType)
	}

	client.SaveAndCryptDataInBufferToSend(true)
}
