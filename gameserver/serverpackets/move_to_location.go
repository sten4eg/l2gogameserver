package serverpackets

import (
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewMoveToLocation(location *clientpackets.Location, client *models.Client) []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x2f)

	buffer.WriteD(client.CC.CharId)

	buffer.WriteD(location.TargetX)
	buffer.WriteD(location.TargetY)
	buffer.WriteD(location.TargetZ)

	buffer.WriteD(location.OriginX)
	buffer.WriteD(location.OriginY)
	buffer.WriteD(location.OriginZ)

	client.CC.SetXYZ(location.TargetX, location.TargetY, location.TargetZ)
	return buffer.Bytes()
}
