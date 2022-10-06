package serverpackets

import "l2gogameserver/packets"

func PartySmallWindowDeleteAll() *packets.Buffer {
	buffer := packets.Get()
	buffer.WriteSingleByte(0x50)
	return buffer
}
