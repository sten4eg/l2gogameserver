package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewShowMiniMap(client *models.Client) {

	client.Buffer.WriteH(0)
	client.Buffer.WriteSingleByte(0xa3)
	client.Buffer.WriteD(1665)
	client.Buffer.WriteSingleByte(2)
}
