package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func AskJoinParty(targetName string, partyDistributionType interfaces.PartyDistributionTypeInterface) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x39)
	buffer.WriteS(targetName)
	buffer.WriteD(partyDistributionType.GetId())

	return buffer
}
