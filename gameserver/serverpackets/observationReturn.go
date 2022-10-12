package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ObservationReturn(character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	x, y, z := character.GetXYZ()

	buffer.WriteSingleByte(0xEC)
	buffer.WriteD(x) //x 53
	buffer.WriteD(y) //y 57
	buffer.WriteD(z) //z 61

	return buffer.Bytes()
}
