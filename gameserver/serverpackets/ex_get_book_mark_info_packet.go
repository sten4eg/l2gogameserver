package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func ExGetBookMarkInfoPacket(client *models.Client) {

	client.Buffer.WriteSingleByte(0xFE)
	client.Buffer.WriteH(0x84)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)

	client.SaveAndCryptDataInBufferToSend(true)
}
