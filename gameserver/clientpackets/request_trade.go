package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

//interfaces.ReciverAndSender

type TradeRequestInterface interface {
	EncryptAndSend([]byte)
	GetCurrentChar() interfaces.CharacterI
}

func TradeRequest(data []byte, client TradeRequestInterface) {
	var packet = packets.NewReader(data)
	targetObjectId := packet.ReadInt32()

	target := broadcast.GetCharacterByObjectId(targetObjectId)

	if target == nil {
		pkg := sysmsg.SystemMessage(sysmsg.TargetIsIncorrect)
		client.EncryptAndSend(pkg)
		return
	}
	if target.GetObjectId() == client.GetCurrentChar().GetObjectId() {
		pkg := sysmsg.SystemMessage(sysmsg.TargetIsIncorrect)
		client.EncryptAndSend(pkg)
		return
	}

	if target.IsProcessingTransaction() {
		sm := sysmsg.C1IsBusyTryLater
		sm.AddString(target.GetName())

		pkg := sysmsg.SystemMessage(sm)
		client.EncryptAndSend(pkg)
		return
	}

	if client.GetCurrentChar().IsProcessingTransaction() {
		client.EncryptAndSend(sysmsg.SystemMessage(sysmsg.AlreadyTrading))
		return
	}

	if target.GetTradeRefusal() {
		sysmsg.SendCustomSystemMessage("That person is in trade refusal mode.")
		return
	}

	if client.GetCurrentChar().CalculateDistanceTo(target.GetX(), target.GetY(), target.GetZ(), false, false) > 150 {
		client.EncryptAndSend(sysmsg.SystemMessage(sysmsg.TargetTooFar))
		return
	}

	client.GetCurrentChar().OnTransactionRequest(target)
	pkg := serverpackets.TradeSendRequest(target)
	target.EncryptAndSend(pkg)
	sm := sysmsg.RequestC1ForTrade
	sm.AddString(target.GetName())
	client.EncryptAndSend(sysmsg.SystemMessage(sm))

	logger.Info.Println("Отправлен запрос на трейд к", target.GetName())

	//trade.NewRequestTrade(client.GetCurrentChar(), target)

}
