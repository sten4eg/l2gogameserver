package serverpackets

import (
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models"
)

func NewShortCutInit(client *models.Client) {

	client.Buffer.WriteSingleByte(0x45)

	shortCuts := models.GetAllShortCuts(client.CurrentChar.CharId, client.CurrentChar.ClassId)
	client.Buffer.WriteD(int32(len(shortCuts)))

	for _, v := range shortCuts {
		client.Buffer.WriteD(v.ShortcutType)
		client.Buffer.WriteD(v.Slot + (v.Page * models.MaxShortcutsPerBar))

		client.Buffer.WriteD(v.Id)

		shortCutsType := dto.ShortTypes[v.ShortcutType]
		switch shortCutsType {
		case "ITEM":
			client.Buffer.WriteD(0x01)
			client.Buffer.WriteD(0) //sc.getSharedReuseGroup()
			client.Buffer.WriteD(0)
			client.Buffer.WriteD(0)
			client.Buffer.WriteH(0)
			client.Buffer.WriteH(0)
		case "SKILL":
			client.Buffer.WriteD(v.Level)
			client.Buffer.WriteSingleByte(0) // C5
			client.Buffer.WriteD(0x01)       // C6
		case "ACTION", "MACRO", "RECIPE", "BOOKMARK":
			client.Buffer.WriteD(0x01) // C6
		}
	}
	client.SaveAndCryptDataInBufferToSend(true)
}
