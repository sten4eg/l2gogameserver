package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func TradeSendRequest(target interfaces.CharacterI) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x70)
	buffer.WriteD(target.GetObjectId())

	return buffer.Bytes()
}
