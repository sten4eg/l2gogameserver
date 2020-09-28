package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewExSendManorList(client *models.Client) {

	client.Buffer.WriteH(0)
	client.Buffer.WriteSingleByte(0xFE)
	client.Buffer.WriteH(0x22)
	client.Buffer.WriteD(0)
}
