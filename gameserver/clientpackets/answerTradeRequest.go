package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

//AnswerTradeRequest Если пользователь отвечает на запрос трейда
func AnswerTradeRequest(data []byte, sender interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	response := packet.ReadInt32() // 0-отказ,1-принял
	if response == 0 {
		logger.Info.Println("Пользователь не захотел торговать")
		return
	}

	partner := sender.GetCurrentChar().GetActiveRequester()
	if partner == nil {
		sender.EncryptAndSend(serverpackets.TradeDone(0))
		sender.EncryptAndSend(sysmsg.SystemMessage(sysmsg.TargetIsNotFoundInTheGame))
		sender.GetCurrentChar().SetActiveRequester(nil)
		return
	}

	if broadcast.GetCharacterByObjectId(partner.GetObjectId()) == nil {
		sender.EncryptAndSend(serverpackets.TradeDone(0))
		sender.EncryptAndSend(sysmsg.SystemMessage(sysmsg.TargetIsNotFoundInTheGame))
		sender.GetCurrentChar().SetActiveRequester(nil)
		return
	}

	if response == 1 && !partner.IsRequestExpired() {
		sender.GetCurrentChar().StartTrade(partner)
		msg := sysmsg.BeginTradeWithC1
		msg.AddString(partner.GetName())
		sender.EncryptAndSend(sysmsg.SystemMessage(msg))

		msg1 := sysmsg.BeginTradeWithC1
		msg1.AddString(sender.GetCurrentChar().GetName())
		partner.EncryptAndSend(sysmsg.SystemMessage(msg1))

		sender.EncryptAndSend(serverpackets.TradeStart(sender.GetCurrentChar()))
		partner.EncryptAndSend(serverpackets.TradeStart(partner))
	} else {
		sm := sysmsg.C1DeniedTradeRequest
		sm.AddString(sender.GetCurrentChar().GetName())
		partner.EncryptAndSend(sysmsg.SystemMessage(sm))
	}

	sender.GetCurrentChar().SetActiveRequester(nil)
	partner.OnTransactionResponse()

}
