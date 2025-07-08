package serverpackets

import (
	"l2gogameserver/packets"
)

func ExGetBookMarkInfoPacket() []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0x84)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	return buffer.Bytes()
}
