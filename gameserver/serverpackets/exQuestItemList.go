package serverpackets

import (
	"l2gogameserver/packets"
)

func ExQuestItemList() []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xC6)
	buffer.WriteH(0)
	buffer.WriteH(0)

	return buffer.Bytes()
}
