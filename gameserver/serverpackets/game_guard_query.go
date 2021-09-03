package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func GameGuardQuery(client *models.Client) {

	client.Buffer.WriteSingleByte(0x74)
	client.Buffer.WriteD(0x27533DD9)
	client.Buffer.WriteD(0x2E72A51D)
	client.Buffer.WriteD(0x2017038B)
	client.Buffer.WriteDU(0xC35B1EA3)
	client.SaveAndCryptDataInBufferToSend(true)
}
