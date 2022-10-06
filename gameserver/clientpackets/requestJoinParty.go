package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/party"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestJoinParty(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	name := reader.ReadString()
	partyDistributionTypeId := reader.ReadInt32()

	character := client.GetCurrentChar()
	target := getCharacterByName(character.GetCurrentRegion(), name)

	if character == nil {
		return
	}

	if target == nil {
		character.SendSysMsg(sysmsg.FirstSelectUserToInviteToParty)
		return
	}

	//TODO if ((target.getClient() == null) || target.getClient().isDetached())

	//TODO if (requestor.isPartyBanned())

	//TODO if (target.isPartyBanned())

	//TODO if (!target.isVisibleFor(requestor))

	if target.IsinParty() {
		msg := sysmsg.C1IsAlreadyInParty
		msg.AddString(target.GetName())
		character.SendSysMsg(msg)
		return
	}

	//TODO if (BlockList.isBlocked(target, requestor))

	if target == character {
		character.SendSysMsg(sysmsg.YouHaveInvitedTheWrongTarget)
		return
	}

	//TODO if (target.isCursedWeaponEquipped() || requestor.isCursedWeaponEquipped())

	//TODO if (target.isJailed() || requestor.isJailed())

	//TODO if (target.isInOlympiadMode() || requestor.isInOlympiadMode())

	msg := sysmsg.C1InvitedToParty
	msg.AddCharacterName(target.GetName())
	character.SendSysMsg(msg)

	if !character.IsinParty() {
		createNewParty(character, target, partyDistributionTypeId)
	} else {
		if false { //TODO isInDimensionalRift()
			// requestor.sendMessage("You cannot invite a player when you are in the Dimensional Rift.")
		} else {
			addTargetToParty(target, character)
		}
	}
}

func getCharacterByName(region interfaces.WorldRegioner, name string) interfaces.CharacterI {
	for _, r := range region.GetNeighbors() {
		for _, targer := range r.GetCharsInRegion() {
			if targer.GetName() == name {
				return targer
			}
		}
	}
	return nil
}

func createNewParty(leader, target interfaces.CharacterI, partyDistributionTypeId int32) {
	partyDistributionType, _ := party.GetPartyDistributionTypeById(partyDistributionTypeId)
	if partyDistributionType == nil {
		return
	}

	if true { //TODO (!target.isProcessingRequest())
		buffer := serverpackets.AskJoinParty(leader.GetName(), partyDistributionType)
		target.SendBuf(buffer)
		target.SetActiveRequester(leader)
		//TODO
		leader.SetPartyDistributionType(partyDistributionType)
	} else {
		return
	}
}

func addTargetToParty(target, requester interfaces.CharacterI) {
	party := requester.GetParty()

	if party.GetLeader().GetObjectId() != requester.GetObjectId() {
		requester.SendSysMsg(sysmsg.OnlyLeaderCanInvite)
		return
	}

	if len(party.GetMembers()) >= 9 {
		requester.SendSysMsg(sysmsg.PartyFull)
		return
	}

	//TODO проверка что нет активного приглашения в группу

	//TODO проверка что у таргета нет активного приглашения в другую группу
	if true {
		requester.OnTransactionRequest(target)
		target.SendBuf(serverpackets.AskJoinParty(requester.GetName(), party.GetDistributionType()))
		//TODO party.setPendingInvitation(true)
	} else {
		msg := sysmsg.C1IsBusyTryLater
		msg.AddString(target.GetName())
		requester.SendSysMsg(msg)
	}
}
