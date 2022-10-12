package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func QuestList(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x86)
	buffer.WriteH(0)
	x := make([]byte, 128)
	buffer.WriteSlice(x)

	return buffer.Bytes()
}
