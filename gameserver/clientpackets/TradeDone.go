package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

//Игрок подтвердил сделку
func TradeDone(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	response := packet.ReadInt32() // 1 - пользователь нажал ОК, 0 пользователь отменил трейд

	player2, exchange, ok := trade.FindTrade(client.GetCurrentChar())
	if !ok {
		logger.Info.Println("Обменивающихся не найдено")
		return
	}
	if response == 1 {
		if exchange.Sender.ObjectId == client.GetCurrentChar().GetObjectId() {
			exchange.Sender.Completed = true
			logger.Info.Printf("Игрок %s подтвердил сделку\n", client.GetCurrentChar().GetName())
			serverpackets.TradeOtherDone(player2)
		}
		if exchange.Recipient.ObjectId == client.GetCurrentChar().GetObjectId() {
			exchange.Recipient.Completed = true
			logger.Info.Printf("Игрок %s подтвердил сделку\n", client.GetCurrentChar().GetName())
			serverpackets.TradeOtherDone(client.GetCurrentChar())
		}
		if exchange.Recipient.Completed == true && exchange.Sender.Completed == true {
			logger.Info.Println("Обмен завершен успешно", exchange.Recipient.Completed, exchange.Sender.Completed)
			serverpackets.TradeOK(client.GetCurrentChar(), player2)
			tradeUserInfo := trade.TradeAddInventory(client.GetCurrentChar(), player2, exchange)

			for _, tradeData := range tradeUserInfo {
				//getItem, _ := tradeData.Player.(*models.Character).Inventory.ExistItemID(tradeData.Item.Id)
				ut1 := utils.GetPacketByte()
				ut1.SetData(serverpackets.InventoryUpdate(tradeData.Item, tradeData.UpdateType))
				tradeData.Player.EncryptAndSend(ut1.GetData())
			}

			if ok = trade.UserClear(client.GetCurrentChar()); !ok {
				logger.Info.Println("Трейд не найден")
				return
			}
		}
	} else if response == 0 {
		serverpackets.TradeCancel(client.GetCurrentChar(), player2)
		if ok = trade.UserClear(client.GetCurrentChar()); !ok {
			logger.Info.Println("Трейд не найден")
			return
		}
	}

}
