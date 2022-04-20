package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ExGetBookMarkInfoPacket(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0x84)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
