package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func PartySmallWindowAdd(member interfaces.CharacterI, party interfaces.PartyInterface) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x4F)
	buffer.WriteD(party.GetLeaderObjectId())
	buffer.WriteD(party.GetDistributionType().GetId())
	buffer.WriteD(member.GetObjectId())
	buffer.WriteS(member.GetName())
	buffer.WriteD(0) // CurrentCp
	buffer.WriteD(0) // MaxCp
	buffer.WriteD(member.GetCurrentHp())
	buffer.WriteD(member.GetMaxHp())
	buffer.WriteD(member.GetCurrentMp())
	buffer.WriteD(member.GetMaxMp())
	buffer.WriteD(1) // Level
	buffer.WriteD(member.GetClassId())
	buffer.WriteD(0x00) // ?
	buffer.WriteD(0x00) // ?

	return buffer
}
