package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func PrivateStoreMsgBuy(character interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xBF)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteS(character.GetBuyList().GetTitle())

	return buffer
}
