package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func PrivateStoreMsgSell(character interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()
	msg := character.GetSellList().GetTitle()

	buffer.WriteSingleByte(0xA2)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteS(msg)

	return buffer
}
