package serverpackets

import "l2gogameserver/gameserver/models"

func NewRestartResponse(client *models.Client) {

	client.Buffer.WriteH(0) //reserve
	client.Buffer.WriteSingleByte(0x71)
	client.Buffer.WriteD(1)
}
