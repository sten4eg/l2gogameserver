package clientpackets

import (
	"fmt"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestActionUse(client interfaces.ReciverAndSender, data []byte) {
	packet := packets.NewReader(data)

	actionId := packet.ReadInt32()
	ctrlPressed := packet.ReadInt32() == 1
	shiftPressed := packet.ReadInt32() == 1

	_, _ = ctrlPressed, shiftPressed

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	switch actionId {
	default:
		fmt.Printf("Неопознаный второй опкод %x в RequestActionUse\n", data[0])
	case 0:
		ChangeWaitType(client)
	case 10:
		tryOpenPrivateSellShop(client, false)
	case 12:
		tryBroadcastSocial(character, 2) // Greeting
	case 13:
		tryBroadcastSocial(character, 3) // Victory
	case 14:
		tryBroadcastSocial(character, 4) // Advance
	case 24:
		tryBroadcastSocial(character, 6) // Yes
	case 25:
		tryBroadcastSocial(character, 5) // No
	case 26:
		tryBroadcastSocial(character, 7) // Bow
	case 28:
		tryOpenPrivateBuyStore(client)
	case 29:
		tryBroadcastSocial(character, 8) // Unaware
	case 30:
		tryBroadcastSocial(character, 9) // Social Waiting
	case 31:
		tryBroadcastSocial(character, 10) // Laugh
	case 33:
		tryBroadcastSocial(character, 11) // Applaud
	case 34:
		tryBroadcastSocial(character, 12) // Dance
	case 35:
		tryBroadcastSocial(character, 13) // Sorrow
	case 61:
		tryOpenPrivateSellShop(client, true)
	case 62:
		tryBroadcastSocial(character, 14) // Charm
	case 66:
		tryBroadcastSocial(character, 15) // Shyness
	case 71, 72, 73:
		useCoupleSocial(character, actionId-55)

	}

}

func tryOpenPrivateSellShop(client interfaces.ReciverAndSender, isPackageSale bool) {
	c := client.GetCurrentChar()
	if true { //TODO проверка на возможность создания магазина
		if c.GetPrivateStoreType() == privateStoreType.SELL || c.GetPrivateStoreType() == privateStoreType.SELL_MANAGE || c.GetPrivateStoreType() == privateStoreType.PACKAGE_SELL {
			c.SetPrivateStoreType(privateStoreType.NONE)
		}

		if c.GetPrivateStoreType() == privateStoreType.NONE {
			if c.IsSittings() {
				ChangeWaitType(client)
			}
			c.SetPrivateStoreType(privateStoreType.SELL_MANAGE)
			pkg := serverpackets.PrivateStoreManageListSell(c, isPackageSale)
			client.SendBuf(pkg)
		}

	} else {
		if false { //TODO проверка что персонаж находится в зоне, в которой нельзя торговать
			c.SendSysMsg(sysmsg.NoPrivateStoreHere)
		}
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)

	}
}

func tryOpenPrivateBuyStore(client interfaces.ReciverAndSender) {
	c := client.GetCurrentChar()
	if true { //TODO проверка на возможность создания магазина
		if c.GetPrivateStoreType() == privateStoreType.BUY || c.GetPrivateStoreType() == privateStoreType.BUY_MANAGE {
			c.SetPrivateStoreType(privateStoreType.NONE)
		}
		if c.GetPrivateStoreType() == privateStoreType.NONE {
			if c.IsSittings() {
				ChangeWaitType(client)
			}
			c.SetPrivateStoreType(privateStoreType.BUY_MANAGE)
			pkg := serverpackets.PrivateStoreManageListBuy(c)
			client.SendBuf(pkg)
		}
	} else {
		if false { //TODO проверка что персонаж находится в зоне, в которой нельзя торговать
			c.SendSysMsg(sysmsg.NoPrivateStoreHere)
		}
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)
	}
}

func tryBroadcastSocial(character interfaces.CharacterI, id int32) {
	//TODO isFishing()

	if true { //TODO canMakeSocialAction
		broadcast.BroadCastPkgToAroundPlayer(character, serverpackets.SocialAction(character, id))
	}
}

func useCoupleSocial(character interfaces.CharacterI, id int32) {
	targetObject := getTargetByObjectId(character.GetTarget(), character.GetCurrentRegion())
	if targetObject == nil {
		character.SendSysMsg(sysmsg.IncorrectTarget)
		return
	}

	switch target := targetObject.(type) {
	default:
		character.SendSysMsg(sysmsg.IncorrectTarget)
		return
	case interfaces.CharacterI:
		ox, oy, oz := character.GetXYZ()
		mx, my, mz := target.GetXYZ()
		distance := models.CalculateDistance(ox, oy, oz, mx, my, mz, false, false)

		if distance > 125 || distance < 15 || character.GetObjectId() == target.GetObjectId() {
			character.SendSysMsg(sysmsg.TargetDoNotMeetLocRequirements)
			return
		}

		if character.GetPrivateStoreType() != privateStoreType.NONE { // TODO isInCraftMode()
			msg := sysmsg.C1IsInPrivateShopModeOrInABattleAndCannotBeRequestedForACoupleAction
			msg.AddCharacterName(character.GetName())
			character.SendSysMsg(msg)
			return
		}

		// TODO if (requester.isInCombat() || requester.isInDuel() || AttackStanceTaskManager.getInstance().hasAttackStanceTask(requester))

		// TODO isFissing()

		if character.GetKarma() > 0 {
			msg := sysmsg.C1IsInAChaoticStateAndCannotBeRequestedForACoupleAction
			msg.AddCharacterName(character.GetName())
			character.SendSysMsg(msg)
			return
		}

		// TODO if (requester.isInOlympiadMode())

		// TODO if (requester.isInSiege())

		// TODO if (requester.isInHideoutSiege())

		// TODO if (requester.isMounted() || requester.isFlyingMounted() || requester.isInBoat() || requester.isInAirShip())

		// TODO if (requester.isTransformed())

		// TODO if (requester.isAlikeDead())

		if target.GetPrivateStoreType() != privateStoreType.NONE { // TODO isInCraftMode()
			msg := sysmsg.C1IsInPrivateShopModeOrInABattleAndCannotBeRequestedForACoupleAction
			msg.AddCharacterName(target.GetName())
			character.SendSysMsg(msg)
			return
		}

		// TODO if (partner.isInCombat() || partner.isInDuel() || AttackStanceTaskManager.getInstance().hasAttackStanceTask(partner))

		if target.GetMultiSocialAction() > 0 {
			msg := sysmsg.C1IsAlreadyParticipatingInACoupleActionAndCannotBeRequestedForAnotherCoupleAction
			msg.AddCharacterName(target.GetName())
			character.SendSysMsg(msg)
			return
		}

		// TODO if (partner.isFishing())

		if target.GetKarma() > 0 {
			msg := sysmsg.C1IsInAChaoticStateAndCannotBeRequestedForACoupleAction
			msg.AddCharacterName(target.GetName())
			character.SendSysMsg(msg)
			return
		}

		// TODO if (partner.isInOlympiadMode())
		//
		// TODO if (partner.isInHideoutSiege())
		//
		// TODO if (partner.isInSiege())
		//
		// TODO if (partner.isMounted() || partner.isFlyingMounted() || partner.isInBoat() || partner.isInAirShip())
		//
		// TODO if (partner.isTeleporting())
		//
		// TODO if (partner.isTransformed())
		//
		// TODO if (partner.isAlikeDead())
		//
		// TODO if (requester.isAllSkillsDisabled() || partner.isAllSkillsDisabled())

		character.SetMultiSocialAction(id, target.GetObjectId())
		msg := sysmsg.YouHaveRequestedCoupleActionC1
		msg.AddCharacterName(target.GetName())
		character.SendSysMsg(msg)

		// TODO проверки

		target.SendBuf(serverpackets.ExAskCoupleAction(character.GetObjectId(), id))
	}
}
