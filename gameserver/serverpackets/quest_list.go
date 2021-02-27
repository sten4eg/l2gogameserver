package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewQuestList(client *models.Client) {

	client.Buffer.WriteSingleByte(0x86)
	client.Buffer.WriteH(0)
	x := make([]byte, 128)
	client.Buffer.WriteSlice(x)

	client.SaveAndCryptDataInBufferToSend(true)
}
