package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ExBasicActionList(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xfe)
	buffer.WriteH(0x5f)

	buffer.WriteD(0)

	return buffer.Bytes()
}
