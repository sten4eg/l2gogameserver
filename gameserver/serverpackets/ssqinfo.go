package serverpackets

import "l2gogameserver/packets"

func NewSSQInfo() []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x73)
	buffer.WriteH(256)

	return buffer.Bytes()
}
