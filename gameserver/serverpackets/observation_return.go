package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewObservationReturn(user *models.Character, client *models.Client) {

	client.Buffer.WriteH(0)
	client.Buffer.WriteSingleByte(0xEC)
	client.Buffer.WriteD(user.X) //x 53
	client.Buffer.WriteD(user.Y) //y 57
	client.Buffer.WriteD(user.Z) //z 61

}
