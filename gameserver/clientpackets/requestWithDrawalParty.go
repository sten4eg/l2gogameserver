package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/party/messageType"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/utils"
)

func RequestWithDrawalParty(client interfaces.ReciverAndSender) {
	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	party := character.GetParty()
	if party != nil {
		if false { //TOPO party.isInDimensionalRift() && !party.getDimensionalRift().getRevivedAtWaitingRoom().contains(player)
			return
		} else {
			RemovePartyMember(character, party, messageType.Left)

		}

		//TODO character.IsinPartyMatchRoom()
	}
}

func RemovePartyMember(character interfaces.CharacterI, p interfaces.PartyInterface, reason messageType.MessageType) {
	if p.IsMemberInParty(character) {
		isLeader := p.IsLeader(character)
		if !p.IsDisbanding() {
			if len(p.GetMembers()) == 2 {
				disbandParty(p)
			}
		}

		index := p.GetMemberIndex(character)
		members := append(p.GetMembers()[:index], p.GetMembers()[index:]...)
		p.SetMembers(members)
		//TODO recalculatePartyLevel()

		//TODO if (player.isFestivalParticipant())

		//TODO Channeling a player!

		if reason == messageType.Expelled {
			character.SendSysMsg(sysmsg.HaveBeenExpelledFromParty)
			msg := sysmsg.C1WasExpelledFromParty
			msg.AddString(character.GetName())
			p.BroadcastParty(sysmsg.SystemMessage(msg))
		} else if reason == messageType.Left || reason == messageType.Disconnected {
			character.SendSysMsg(sysmsg.YouLeftParty)
			msg := sysmsg.C1LeftParty
			msg.AddString(character.GetName())
			p.BroadcastParty(sysmsg.SystemMessage(msg))
		}

		character.SendBuf(serverpackets.PartySmallWindowDeleteAll())
		character.SetParty(nil)
		p.BroadcastParty(serverpackets.PartySmallWindowDelete(character).Bytes())
		if false { //TODO hasSummon
			return
		}

		if false { //TODO (isInDimensionalRift())
			return
		}

		if false { //TODO (isInCommandChannel())
			return
		}

		if isLeader && len(p.GetMembers()) > 1 && reason == messageType.Disconnected {
			msg := sysmsg.C1HasBecomeAPartyLeader
			msg.AddString(p.GetLeader().GetName())
			p.BroadcastParty(sysmsg.SystemMessage(msg))
			broadcastToPartyMembersNewLeader(p)
		} else if len(p.GetMembers()) == 1 {
			if false { //TODO isInCommandChannel()
				return
			}

			if p.GetLeader() != nil {
				p.GetLeader().SetParty(nil)
			}
			p.SetMembers(nil)

		}
	}
}

func disbandParty(party interfaces.PartyInterface) {
	party.SetDisbanding(true)
	msg := sysmsg.PartyDispersed
	pb := utils.GetPacketByte()
	pb.SetData(sysmsg.SystemMessage(msg))

	for _, member := range party.GetMembers() {
		if member != nil {
			RemovePartyMember(member, party, messageType.None)
			member.EncryptAndSend(pb.GetData())
		}
	}

	pb.Release()
}

func broadcastToPartyMembersNewLeader(p interfaces.PartyInterface) {
	for _, member := range p.GetMembers() {
		if member != nil {
			member.SendBuf(serverpackets.PartySmallWindowDeleteAll())
			member.SendBuf(serverpackets.PartySmallWindowAll(member, p))
			broadcast.BroadcastUserInfo(member)
		}
	}
}
