package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewExBasicActionList(client *models.Client) {

	client.Buffer.WriteSingleByte(0xfe)
	client.Buffer.WriteH(0x5f)

	client.Buffer.WriteD(0)

	client.SaveAndCryptDataInBufferToSend(true)
}
