package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func DropItem(item interfaces.MyItemInterface, charObjectId int32) *packets.Buffer {
	buf := packets.Get()
	x, y, z := item.GetCoordinate()

	buf.WriteSingleByte(0x16)
	buf.WriteD(charObjectId)
	buf.WriteD(item.GetObjectId())
	buf.WriteD(item.GetId())

	buf.WriteD(x)
	buf.WriteD(y)
	buf.WriteD(z)

	buf.WriteD(utils.BoolToInt32(item.IsStackable()))
	buf.WriteQ(item.GetCount())

	buf.WriteD(0x01)

	return buf
}
