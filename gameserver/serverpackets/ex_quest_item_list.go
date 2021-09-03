package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func ExQuestItemList(client *models.Client) {

	client.Buffer.WriteSingleByte(0xFE)
	client.Buffer.WriteH(0xC6)
	client.Buffer.WriteH(0)
	client.Buffer.WriteH(0)
	client.SaveAndCryptDataInBufferToSend(true)
}
