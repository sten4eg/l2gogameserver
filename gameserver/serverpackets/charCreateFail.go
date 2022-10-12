package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func CharCreateFail(client interfaces.ReciverAndSender, reason int32) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x10)
	buffer.WriteD(reason)

	return buffer.Bytes()
}
