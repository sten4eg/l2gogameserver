package serverpackets

import "l2gogameserver/gameserver/models"

func RestartResponse(client *models.Client) {

	client.Buffer.WriteSingleByte(0x71)
	client.Buffer.WriteD(1)
	client.SaveAndCryptDataInBufferToSend(true)
}
