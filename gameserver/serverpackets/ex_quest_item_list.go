package serverpackets

import "l2gogameserver/packets"

func NewExQuestItemList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xC6)
	buffer.WriteH(0)
	buffer.WriteH(0)
	return buffer.Bytes()
}
