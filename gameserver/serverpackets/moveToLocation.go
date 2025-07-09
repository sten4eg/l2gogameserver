package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/world"
	"l2gogameserver/packets"
)

func MoveToLocation(location *world.BackwardToLocation, character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x2f)

	buffer.WriteD(character.GetObjectId())

	buffer.WriteD(location.TargetX)
	buffer.WriteD(location.TargetY)
	buffer.WriteD(location.TargetZ)

	buffer.WriteD(location.OriginX)
	buffer.WriteD(location.OriginY)
	buffer.WriteD(location.OriginZ)

	character.SetXYZ(location.TargetX, location.TargetY, location.TargetZ)
	return buffer.Bytes()
}
