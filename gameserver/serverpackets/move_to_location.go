package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewMoveToLocation(location *models.BackwardToLocation, client *models.Client) []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x2f)

	buffer.WriteD(client.CurrentChar.CharId)

	buffer.WriteD(location.TargetX)
	buffer.WriteD(location.TargetY)
	buffer.WriteD(location.TargetZ)

	buffer.WriteD(location.OriginX)
	buffer.WriteD(location.OriginY)
	buffer.WriteD(location.OriginZ)

	client.CurrentChar.SetXYZ(location.TargetX, location.TargetY, location.TargetZ)
	return buffer.Bytes()
}
