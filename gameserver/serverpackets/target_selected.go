package serverpackets

import "l2gogameserver/gameserver/models"

func NewTargetSelected(objectId, targetId, x, y, z int32, client *models.Client) {

	client.Buffer.WriteSingleByte(0x23)
	client.Buffer.WriteD(objectId)
	client.Buffer.WriteD(targetId)
	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)
	client.Buffer.WriteD(0)
	client.SaveAndCryptDataInBufferToSend(true)
}
