package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewStaticObject(client *models.Client) {

	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)
	client.Buffer.WriteD(0)

	client.SaveAndCryptDataInBufferToSend(true)
}
