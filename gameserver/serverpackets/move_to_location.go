package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func MoveToLocation(location *models.BackwardToLocation, client *models.Client) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x2f)

	buffer.WriteD(client.CurrentChar.ObjectId)

	buffer.WriteD(location.TargetX)
	buffer.WriteD(location.TargetY)
	buffer.WriteD(location.TargetZ)

	buffer.WriteD(location.OriginX)
	buffer.WriteD(location.OriginY)
	buffer.WriteD(location.OriginZ)

	client.CurrentChar.SetXYZ(location.TargetX, location.TargetY, location.TargetZ)
	return buffer.Bytes()
}
