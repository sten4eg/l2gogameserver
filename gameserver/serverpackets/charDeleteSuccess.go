package serverpackets

import "l2gogameserver/packets"

func CharDeleteSuccess() *packets.Buffer {
	buffer := packets.Get()
	buffer.WriteSingleByte(0x1D)
	return buffer
}
