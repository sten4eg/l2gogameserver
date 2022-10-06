package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func PartySmallWindowDelete(member interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x51)
	buffer.WriteD(member.GetObjectId())
	buffer.WriteS(member.GetName())

	return buffer
}
