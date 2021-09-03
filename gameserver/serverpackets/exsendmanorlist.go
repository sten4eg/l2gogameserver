package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func ExSendManorList(client *models.Client) {

	client.Buffer.WriteSingleByte(0xFE)
	client.Buffer.WriteH(0x22)
	client.Buffer.WriteD(0)
	client.SaveAndCryptDataInBufferToSend(true)
}
