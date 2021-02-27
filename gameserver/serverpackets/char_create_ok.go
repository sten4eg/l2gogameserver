package serverpackets

import "l2gogameserver/gameserver/models"

func NewCharCreateOk(client *models.Client) {

	client.Buffer.WriteSingleByte(0x0f)
	client.Buffer.WriteD(1)
	client.SaveAndCryptDataInBufferToSend(true)
}
