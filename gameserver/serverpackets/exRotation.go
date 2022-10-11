package serverpackets

import "l2gogameserver/packets"

func ExRotation(charId, heading int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xC1)
	buffer.WriteD(charId)
	buffer.WriteD(heading)

	return buffer
}
