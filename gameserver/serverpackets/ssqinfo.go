package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewSSQInfo(client *models.Client) {

	client.Buffer.WriteSingleByte(0x73)
	client.Buffer.WriteH(256)

	client.SaveAndCryptDataInBufferToSend(true)

}
