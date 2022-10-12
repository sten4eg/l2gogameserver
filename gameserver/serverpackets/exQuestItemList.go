package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ExQuestItemList(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xC6)
	buffer.WriteH(0)
	buffer.WriteH(0)

	return buffer.Bytes()
}
