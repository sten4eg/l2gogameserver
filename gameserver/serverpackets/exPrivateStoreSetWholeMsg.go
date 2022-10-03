package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ExPrivateStoreSetWholeMsg(character interfaces.CharacterI, msg string) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0x80)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteS(msg)

	return buffer
}
