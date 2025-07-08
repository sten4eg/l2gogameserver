package serverpackets

import (
	"l2gogameserver/packets"
)

func SsqInfo() []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x73)
	buffer.WriteH(256)

	return buffer.Bytes()
}
