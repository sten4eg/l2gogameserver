package serverpackets

import "l2gogameserver/packets"

func NewSendMacroList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xE8)
	buffer.WriteD(0)
	buffer.WriteSingleByte(0)
	buffer.WriteSingleByte(0)
	buffer.WriteSingleByte(0)

	return buffer.Bytes()
}
