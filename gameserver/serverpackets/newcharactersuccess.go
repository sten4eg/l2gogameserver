package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func CharacterSuccess(client *models.Client) {

	client.Buffer.WriteSingleByte(0x0D)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0x46)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0x0A)
	client.Buffer.WriteD(0x46)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0x0A)
	client.Buffer.WriteD(0x46)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0x0A)
	client.Buffer.WriteD(0x46)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0x0A)
	client.Buffer.WriteD(0x46)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0x0A)
	client.Buffer.WriteD(0x46)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(0x0A)
	client.SaveAndCryptDataInBufferToSend(true)

}
