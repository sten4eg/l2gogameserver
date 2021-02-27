package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewObservationReturn(user *models.Character, client *models.Client) {

	client.Buffer.WriteSingleByte(0xEC)
	client.Buffer.WriteD(user.Coordinates.X) //x 53
	client.Buffer.WriteD(user.Coordinates.Y) //y 57
	client.Buffer.WriteD(user.Coordinates.Z) //z 61
	client.SaveAndCryptDataInBufferToSend(true)
}
