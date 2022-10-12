package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ShortBuffStatusUpdate(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()
	buffer.WriteSingleByte(0xfa)
	buffer.WriteD(1242)
	buffer.WriteD(1)
	buffer.WriteD(20)

	return buffer.Bytes()
}
