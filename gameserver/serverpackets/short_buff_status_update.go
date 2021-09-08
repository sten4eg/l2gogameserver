package serverpackets

import "l2gogameserver/gameserver/models"

func ShortBuffStatusUpdate(client *models.Client) {
	client.Buffer.WriteSingleByte(0xfa)
	client.Buffer.WriteD(1242)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(20)
	client.SaveAndCryptDataInBufferToSend(true)
}
