package clientpackets

import (
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestRestart(clientI interfaces.ReciverAndSender) {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return
	}

	client.SaveUser()

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	if character.GetActiveEnchantItemId() != -1 { //TODO Заменить на константу и дописать проверку
		client.SendBuf(serverpackets.RestartResponse(0))
		return
	}

	//TODO character.isLocker нужна?

	if character.GetPrivateStoreType() != privateStoreType.NONE {
		msg := sysmsg.S1
		msg.AddString("Cannot restart while trading")
		client.SendSysMsg(msg)
		client.SendBuf(serverpackets.RestartResponse(0))
		return
	}

	//TODO if (AttackStanceTaskManager.getInstance().hasAttackStanceTask(player) && !(player.isGM() && general().gmRestartFighting()))

	//TODO if (player.isFestivalParticipant())

	//TODO if (player.isBlockedFromExit())

	//TODO удалить персонажа из заны с боссом

	broadcast.BroadCastBufferToAroundPlayersWithoutSelf(client, serverpackets.DeleteObject(character.GetObjectId()))
	gameserver.CharOffline(client)
	client.SendBuf(serverpackets.RestartResponse(1))
	client.SendBuf(serverpackets.CharSelectionInfo(client))
	client.SetState(clientStates.Authed)

}
