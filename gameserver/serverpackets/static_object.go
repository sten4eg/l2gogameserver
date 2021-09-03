package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func StaticObject(client *models.Client) {

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
