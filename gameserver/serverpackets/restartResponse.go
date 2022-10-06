package serverpackets

import "l2gogameserver/packets"

func RestartResponse(result int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x71)
	buffer.WriteD(result)

	return buffer
}
