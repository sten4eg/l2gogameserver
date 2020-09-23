package serverpackets

import "l2gogameserver/packets"

func NewExSendManorList() []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0x22)
	buffer.WriteD(0)
	return buffer.Bytes()
}
