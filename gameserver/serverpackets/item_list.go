package serverpackets

import "l2gogameserver/packets"

func NewItemList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x11)
	buffer.WriteH(0)
	buffer.WriteH(0)

	return buffer.Bytes()
}
