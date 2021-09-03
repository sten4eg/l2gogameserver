package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func ExBrExtraUserInfo(client *models.Client) {

	client.Buffer.WriteSingleByte(0xFE)
	client.Buffer.WriteH(0xDA)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.SaveAndCryptDataInBufferToSend(true)
}
