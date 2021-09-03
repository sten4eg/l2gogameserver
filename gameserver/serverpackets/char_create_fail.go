package serverpackets

import "l2gogameserver/gameserver/models"

func CharCreateFail(client *models.Client, reason int32) {

	client.Buffer.WriteSingleByte(0x10)
	client.Buffer.WriteD(reason)
	client.SaveAndCryptDataInBufferToSend(true)
}
