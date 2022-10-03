package serverpackets

import "l2gogameserver/packets"

func StopRotation(objectId, degree, speed int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x61)
	buffer.WriteD(objectId)
	buffer.WriteD(degree)
	buffer.WriteD(speed)
	buffer.WriteD(0)

	return buffer
}
