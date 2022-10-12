package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

// SetupGauge полоска над персонажем во время каста скила
func SetupGauge(character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x6b)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(0) // color 0-blue 1-red 2-cyan 3-green

	buffer.WriteD(4132)
	buffer.WriteD(4132)

	return buffer.Bytes()

}
