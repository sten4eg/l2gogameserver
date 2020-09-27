package serverpackets

import (
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/packets"
)

func NewMoveToLocation(location *clientpackets.Location) []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x2f)

	buffer.WriteD(1)

	buffer.WriteD(location.TargetX)
	buffer.WriteD(location.TargetY)
	buffer.WriteD(location.TargetZ)

	buffer.WriteD(location.OriginX)
	buffer.WriteD(location.OriginY)
	buffer.WriteD(location.OriginZ)

	return buffer.Bytes()
}
