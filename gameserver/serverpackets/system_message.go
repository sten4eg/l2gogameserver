package serverpackets

import "l2gogameserver/packets"

func NewSystemMessage() []byte {
	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x62)
	buffer.WriteD(1129)
	buffer.WriteD(1)
	buffer.WriteD(0)
	buffer.WriteS("fuck")
	return buffer.Bytes()
}
