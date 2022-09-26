package serverpackets

import (
	"l2gogameserver/packets"
)

func TargetSelected(objectId, targetId, x, y, z int32) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x23)
	buffer.WriteD(objectId)
	buffer.WriteD(targetId)
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)
	buffer.WriteD(0)

	return buffer.Bytes()
}
