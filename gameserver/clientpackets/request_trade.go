package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func TradeRequest(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	targetObjectId := packet.ReadInt32()

	target := broadcast.GetCharacterByObjectId(targetObjectId)

	if target == nil {
		pkg := serverpackets.SystemMessage(sysmsg.TargetIsIncorrect)
		client.EncryptAndSend(pkg)
		return
	}
	if target.GetObjectId() == client.GetCurrentChar().GetObjectId() {
		pkg := serverpackets.SystemMessage(sysmsg.TargetIsIncorrect)
		client.EncryptAndSend(pkg)
		return
	}

	if target.IsProcessingTransaction() {
		sm := sysmsg.C1IsBusyTryLater
		sm.AddString(target.GetName())

		pkg := serverpackets.SystemMessage(sm)
		client.EncryptAndSend(pkg)
		return
	}

	if client.GetCurrentChar().IsProcessingTransaction() {
		client.EncryptAndSend(serverpackets.SystemMessage(sysmsg.AlreadyTrading))
		return
	}

	if target.GetTradeRefusal() {
		serverpackets.SendCustomSystemMessage("That person is in trade refusal mode.")
		return
	}

	if client.GetCurrentChar().CalculateDistanceTo(target, false, false) > 150 {
		client.EncryptAndSend(serverpackets.SystemMessage(sysmsg.TargetTooFar))
		return
	}

	pkg := serverpackets.TradeSendRequest(target)
	target.EncryptAndSend(pkg)
	sm := sysmsg.RequestC1ForTrade
	sm.AddString(target.GetName())
	client.EncryptAndSend(serverpackets.SystemMessage(sm))

	logger.Info.Println("Отправлен запрос на трейд к", target.GetName())

	//trade.NewRequestTrade(client.GetCurrentChar(), target)

}
