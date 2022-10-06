package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestAnswerJoinParty(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	respose := reader.ReadInt32()

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	requester := character.GetActiveRequester()
	if requester == nil {
		return
	}

	buffer := serverpackets.JoinParty(respose)
	requester.SendBuf(buffer)

	switch respose {
	case -1:
		msg := sysmsg.C1IsSetToRefusePartyRequest
		msg.AddCharacterName(character.GetName())
		requester.SendSysMsg(msg)
	case 0:
		break
	case 1:
		if requester.IsinParty() {
			if requester.GetParty().GetMemberCount() >= 9 {
				character.SendSysMsg(sysmsg.PartyFull)
				requester.SendSysMsg(sysmsg.PartyFull)
				return
			}
			success := character.JoinParty(requester.GetParty())
			if success {
				onJoinParty(character, requester.GetParty())
			}
		} else {
			party := models.NewParty(requester, requester.GetPartyDistributionType())
			requester.SetParty(party)
			success := character.JoinParty(requester.GetParty())
			if success {
				onJoinParty(character, requester.GetParty())
			}
		}

		//TODO какие то проверки и настройки

		character.SetActiveRequester(nil)

	}
}

// Отпавка пакетов для обновления UI при успешном присоединений к группе
func onJoinParty(character interfaces.CharacterI, party interfaces.PartyInterface) {
	buffer := serverpackets.PartySmallWindowAll(character, party)
	character.SendBuf(buffer)

	for _, member := range party.GetMembers() {
		if false { //TODO hasSummon
			//TODO Пакет
			_ = member
		}
	}

	msg := sysmsg.YouJoinedS1Party
	msg.AddString(party.GetLeader().GetName())
	character.SendSysMsg(msg)

	msg = sysmsg.C1JoinedParty
	msg.AddString(character.GetName())
	//TODO сделать функцию бродкаст
	for _, member := range party.GetMembers() {
		member.SendSysMsg(msg)
	}

	buffer2 := serverpackets.PartySmallWindowAdd(character, party)
	broadcast.BroadCastBufferToAroundPlayersWithoutSelf(character, buffer2)

	if false { //TODO hasSummon

	}

	//L2Summon summon;
	//for (L2PcInstance member : getMembers()) {
	//	if (member != null) {
	//		member.updateEffectIcons(true); // update party icons only
	//		summon = member.getSummon();
	//		member.broadcastUserInfo();
	//		if (summon != null) {
	//			summon.updateEffectIcons();
	//		}
	//	}
	//}

	//if (isInCommandChannel()) {
	//	player.sendPacket(ExOpenMPCC.STATIC_PACKET);
	//}

}
