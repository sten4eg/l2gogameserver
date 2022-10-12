package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func PledgeInfo(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x89)
	buffer.WriteD(0)
	buffer.WriteS("")
	buffer.WriteS("")

	return buffer.Bytes()
}
