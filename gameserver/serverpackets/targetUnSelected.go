package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func TargetUnselected(character interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	character.SetTarget(0)

	x, y, z := character.GetXYZ()

	buffer.WriteSingleByte(0x24)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)
	buffer.WriteD(0)

	return buffer.Bytes()
}
