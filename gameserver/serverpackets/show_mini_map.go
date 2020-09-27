package serverpackets

import "l2gogameserver/packets"

func NewShowMiniMap() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xa3)
	buffer.WriteD(1665)
	buffer.WriteSingleByte(2)
	return buffer.Bytes()
}
