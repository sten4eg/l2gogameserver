package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/party/messageType"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/utils"
)

func RequestWithDrawalParty(client interfaces.NewClientCtxInterface, gs GsInterfNew) {
	character := client.GetAccount().GetCurrentChar()
	if character == nil {
		return
	}

	party := character.GetParty()
	if party != nil {
		if false { //TOPO party.isInDimensionalRift() && !party.getDimensionalRift().getRevivedAtWaitingRoom().contains(player)
			return
		} else {
			RemovePartyMember(character, party, messageType.Left, gs)

		}

		//TODO character.IsinPartyMatchRoom()
	}
}

func RemovePartyMember(character interfaces.CharacterI, party interfaces.PartyInterface, reason messageType.MessageType, gs GsInterfNew) {
	if party.IsMemberInParty(character) {
		isLeader := party.IsLeader(character)
		if !party.IsDisbanding() {
			if len(party.GetMembers()) == 2 {
				disbandParty(party, gs)
			}
		}

		party.RemoveMember(character)
		//TODO recalculatePartyLevel()

		//TODO if (player.isFestivalParticipant())

		//TODO Channeling a player!

		if reason == messageType.Expelled {
			character.SendSysMsg(sysmsg.HaveBeenExpelledFromParty)
			msg := sysmsg.C1WasExpelledFromParty
			msg.AddString(character.GetName())
			party.BroadcastParty(sysmsg.SystemMessage(msg))
		} else if reason == messageType.Left || reason == messageType.Disconnected {
			character.SendSysMsg(sysmsg.YouLeftParty)
			msg := sysmsg.C1LeftParty
			msg.AddString(character.GetName())
			party.BroadcastParty(sysmsg.SystemMessage(msg))
		}

		gs.GetClientByLogin(character.GetAccountLogin()).SendBuf(serverpackets.PartySmallWindowDeleteAll())
		character.SetParty(nil)
		party.BroadcastParty(serverpackets.PartySmallWindowDelete(character).Bytes())
		if false { //TODO hasSummon
			return
		}

		if false { //TODO (isInDimensionalRift())
			return
		}

		if false { //TODO (isInCommandChannel())
			return
		}

		if isLeader && len(party.GetMembers()) > 1 && reason == messageType.Disconnected {
			msg := sysmsg.C1HasBecomeAPartyLeader
			msg.AddString(party.GetLeader().GetName())
			party.BroadcastParty(sysmsg.SystemMessage(msg))
			broadcastToPartyMembersNewLeader(party, gs)
		} else if len(party.GetMembers()) == 1 {
			if false { //TODO isInCommandChannel()
				return
			}

			if party.GetLeader() != nil {
				party.GetLeader().SetParty(nil)
			}
			party.SetMembers(nil)

		}
	}
}

func disbandParty(party interfaces.PartyInterface, gs GsInterfNew) {
	party.SetDisbanding(true)
	msg := sysmsg.PartyDispersed
	pb := utils.GetPacketByte()
	pb.SetData(sysmsg.SystemMessage(msg))

	for _, member := range party.GetMembers() {
		if member != nil {
			RemovePartyMember(member, party, messageType.None, gs)
			member.EncryptAndSend(pb.GetData())
		}
	}

	pb.Release()
}

func broadcastToPartyMembersNewLeader(p interfaces.PartyInterface, gs GsInterfNew) {
	for _, member := range p.GetMembers() {
		if member != nil {
			clientMemer := gs.GetClientByLogin(member.GetAccountLogin())

			clientMemer.SendBuf(serverpackets.PartySmallWindowDeleteAll())
			clientMemer.SendBuf(serverpackets.PartySmallWindowAll(member, p))
			//TODO надо передавать что туда	broadcast.BroadcastUserInfo(member)
		}
	}
}
