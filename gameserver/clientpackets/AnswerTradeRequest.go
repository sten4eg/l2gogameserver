package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
)

//Если пользователь отвечает на запрос трейда
func AnswerTradeRequest(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)
	response := packet.ReadInt32()
	if response == 0 {
		log.Println("Пользователь не захотел торговать")
		return
	}
	if exchange, ok := trade.TradeAnswer(client, response); ok {
		buffer := packets.Get()
		defer packets.Put(buffer)

		ut1 := utils.GetPacketByte()
		ut1.SetData(serverpackets.TradeStart(exchange.Sender.Client))
		exchange.Sender.Client.EncryptAndSend(ut1.GetData())

		ut := utils.GetPacketByte()
		ut.SetData(serverpackets.TradeStart(exchange.Recipient.Client))
		exchange.Recipient.Client.EncryptAndSend(ut.GetData())
	}

}
