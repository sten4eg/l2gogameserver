package serverpackets

import "l2gogameserver/packets"

func ExVoteSystemInfo() []byte {

	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xC9)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	return buffer.Bytes()
}
