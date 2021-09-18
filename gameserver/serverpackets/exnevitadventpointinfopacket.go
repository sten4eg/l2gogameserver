package serverpackets

import "l2gogameserver/packets"

func ExAdventPointInfoPacket() []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xDF)
	buffer.WriteD(0)

	return buffer.Bytes()
}
