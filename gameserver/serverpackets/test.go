package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

// MagicSkillLaunched
func NewTest(character interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x54)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(1216)
	buffer.WriteD(1)
	buffer.WriteD(1)
	buffer.WriteD(character.GetObjectId())

	return buffer.Bytes()
}
