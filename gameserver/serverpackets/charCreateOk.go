package serverpackets

import (
	"l2gogameserver/packets"
)

func CharCreateOk() *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x0f)
	buffer.WriteD(1)

	return buffer
}
