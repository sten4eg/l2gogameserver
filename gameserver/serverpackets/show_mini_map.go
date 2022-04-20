package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ShowMiniMap(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xa3)
	buffer.WriteD(1665)
	buffer.WriteSingleByte(2) //todo currentPeriod

	return buffer.Bytes()
}
