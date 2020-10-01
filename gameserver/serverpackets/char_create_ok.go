package serverpackets

import "l2gogameserver/gameserver/models"

func NewCharCreateOk(client *models.Client) {

	client.Buffer.WriteH(0)
	client.Buffer.WriteSingleByte(0x0f)
	client.Buffer.WriteD(1)
}
