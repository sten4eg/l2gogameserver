package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func LogoutToClient(data []byte, client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x84)
	return buffer.Bytes()
}
