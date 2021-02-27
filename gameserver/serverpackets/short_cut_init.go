package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewShortCutInit(client *models.Client) {

	client.Buffer.WriteSingleByte(0x45)

	client.Buffer.WriteD(0)

	client.SaveAndCryptDataInBufferToSend(true)
}
