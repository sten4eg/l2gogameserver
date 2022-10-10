package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ChangeWaitType(character interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	locx, locy, locz := character.GetXYZ()

	buffer.WriteSingleByte(0x29)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(character.SetSitStandPose())
	buffer.WriteD(locx)
	buffer.WriteD(locy)
	buffer.WriteD(locz)

	return buffer.Bytes()
}
