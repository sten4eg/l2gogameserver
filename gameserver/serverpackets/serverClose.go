package serverpackets

import "l2gogameserver/packets"

func ServerClose() []byte {
	buffer := packets.Get()
	buffer.WriteSingleByte(0x20)
	return buffer.Bytes()
}
