package serverpackets

import "l2gogameserver/packets"

func NewExVoteSystemInfo() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xC9)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	return buffer.Bytes()
}
