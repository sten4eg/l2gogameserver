package serverpackets

import "l2gogameserver/packets"

func CloseBuyList(adena int64, listId int32) *packets.Buffer {

	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteD(0xB7)
	buffer.WriteD(0x00)
	buffer.WriteQ(adena) // current money
	buffer.WriteD(listId)
	buffer.WriteH(0)

	buffer.WriteQ(1)

	return buffer
}
