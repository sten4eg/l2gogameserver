package serverpackets

import (
	"l2gogameserver/packets"
)

func DeleteObject(objectId int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x08)
	buffer.WriteD(objectId)
	buffer.WriteD(0x00)

	return buffer
}
