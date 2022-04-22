package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func TradeRequest(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	targetObjectId := packet.ReadInt32()

	target := broadcast.GetCharacterByObjectId(targetObjectId)
	if target == nil {
		logger.Info.Println("TradeRequest target not found")
		return
	}

	pkg := serverpackets.TradeSendRequest(target)
	target.EncryptAndSend(pkg)
	logger.Info.Println("Отправлен запрос на трейд к", target.GetName())

	trade.NewRequestTrade(client.GetCurrentChar(), target)

}
