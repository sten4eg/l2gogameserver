package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func SendMacroList(client *models.Client) {

	client.Buffer.WriteSingleByte(0xE8)
	client.Buffer.WriteD(0)
	client.Buffer.WriteSingleByte(0)
	client.Buffer.WriteSingleByte(0)
	client.Buffer.WriteSingleByte(0)

	client.SaveAndCryptDataInBufferToSend(true)
}
