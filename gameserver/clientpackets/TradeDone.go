package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
)

//Игрок подтвердил сделку
func TradeDone(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	response := packet.ReadInt32() // 1 - пользователь нажал ОК, 0 пользователь отменил трейд

	player2, exchange, ok := trade.FindTrade(client.GetCurrentChar())
	if !ok {
		log.Println("Обменивающихся не найдено")
		return
	}
	if response == 1 {
		if exchange.Sender.ObjectId == client.GetCurrentChar().GetObjectId() {
			exchange.Sender.Completed = true
			log.Printf("Игрок %s подтвердил сделку\n", client.GetCurrentChar().GetName())
			serverpackets.TradeOtherDone(player2)
		}
		if exchange.Recipient.ObjectId == client.GetCurrentChar().GetObjectId() {
			exchange.Recipient.Completed = true
			log.Printf("Игрок %s подтвердил сделку\n", client.GetCurrentChar().GetName())
			serverpackets.TradeOtherDone(client.GetCurrentChar())
		}
		if exchange.Recipient.Completed == true && exchange.Sender.Completed == true {
			log.Println("Обмен завершен успешно", exchange.Recipient.Completed, exchange.Sender.Completed)
			serverpackets.TradeOK(client.GetCurrentChar(), player2)
			//Теперь сделаем физическую передачу предметов от персонажа к персонажу
			//cplayer, toplayer := trade.TradeAddInventory(client.GetCurrentChar(), player2, exchange)
			tradeUserInfo := trade.TradeAddInventory(client.GetCurrentChar(), player2, exchange)

			for _, tradeData := range tradeUserInfo {
				if tradeData.Item.Id == 57 {
					log.Println("57===>", tradeData.Count, tradeData.Item.ObjId, tradeData.Player.GetName())
				}
				getItem, _ := tradeData.Player.(*models.Character).Inventory.ExistItemID(tradeData.Item.Id)
				log.Println(getItem.ObjId, getItem.Id, getItem.Count)
				ut1 := utils.GetPacketByte()
				ut1.SetData(serverpackets.InventoryUpdate(*getItem, tradeData.UpdateType))
				tradeData.Player.EncryptAndSend(ut1.GetData())
			}

			if ok = trade.TradeUserClear(client.GetCurrentChar()); !ok {
				log.Println("Трейд не найден")
				return
			}
		}
	} else if response == 0 {
		serverpackets.TradeCancel(client.GetCurrentChar(), player2)
		if ok = trade.TradeUserClear(client.GetCurrentChar()); !ok {
			log.Println("Трейд не найден")
			return
		}
	}

}
