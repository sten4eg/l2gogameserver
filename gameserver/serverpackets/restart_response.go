package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func RestartResponse(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x71)
	buffer.WriteD(1)

	//todo останавливать всё у клиента
	defer packets.Put(buffer)
	return buffer.Bytes()
}
