package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewObservationReturn(user *models.Character, client *models.Client) {

	x, y, z := user.GetXYZ()

	client.Buffer.WriteSingleByte(0xEC)
	client.Buffer.WriteD(x) //x 53
	client.Buffer.WriteD(y) //y 57
	client.Buffer.WriteD(z) //z 61
	client.SaveAndCryptDataInBufferToSend(true)
}
