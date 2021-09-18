package serverpackets

import "l2gogameserver/packets"

func ExShowContactList() []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xD3)
	buffer.WriteD(0)

	return buffer.Bytes()
}
