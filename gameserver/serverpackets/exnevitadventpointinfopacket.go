package serverpackets

import "l2gogameserver/packets"

func NewExAdventPointInfoPacket() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xDF)
	buffer.WriteD(0)

	return buffer.Bytes()
}
