package serverpackets

import "l2gogameserver/gameserver/models"

func NewCharCreateFail(client *models.Client, reason int32) {
	client.Buffer.WriteH(0) //reserve
	client.Buffer.WriteSingleByte(0x10)
	client.Buffer.WriteD(reason)
}
