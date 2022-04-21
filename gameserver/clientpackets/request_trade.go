package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
)

func TradeRequest(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	targetObjectId := packet.ReadInt32()

	target := broadcast.GetCharacterByObjectId(targetObjectId)
	if target != nil {
		log.Println("TradeRequest target not found")
		return
	}

	pkg := serverpackets.TradeSendRequest(target)
	target.EncryptAndSend(pkg)
	log.Println("Отправлен запрос на трейд к", target.GetName())

	trade.NewRequestTrade(client.GetCurrentChar(), target)

}
