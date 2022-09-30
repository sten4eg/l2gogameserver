package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func GetItem(item interfaces.MyItemInterface, playerId int32) *packets.Buffer {
	buf := packets.Get()
	x, y, z := item.GetCoordinate()

	buf.WriteSingleByte(0x17)
	buf.WriteD(playerId)
	buf.WriteD(item.GetObjectId())

	buf.WriteD(x)
	buf.WriteD(y)
	buf.WriteD(z)

	return buf

}
