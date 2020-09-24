package serverpackets

import "l2gogameserver/packets"

func NewExBasicActionList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xfe)
	buffer.WriteH(0x5f)

	buffer.WriteD(0)

	return buffer.Bytes()
}
