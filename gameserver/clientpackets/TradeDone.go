package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"log"
)

//Игрок подтвердил сделку
func TradeDone(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)
	response := packet.ReadInt32() // 1 - пользователь нажал ОК, 0 пользователь отменил трейд

	player2, exchange, ok := trade.FindTrade(client)
	if !ok {
		log.Println("Обменивающихся не найдено")
		return
	}
	if response == 1 {
		if exchange.Sender.UserID == client.CurrentChar.ObjectId {
			exchange.Sender.Completed = true
			log.Printf("Игрок %s подтвердил сделку\n", client.CurrentChar.CharName)
			serverpackets.TradeOtherDone(player2)
		}
		if exchange.Recipient.UserID == client.CurrentChar.ObjectId {
			exchange.Recipient.Completed = true
			log.Printf("Игрок %s подтвердил сделку\n", client.CurrentChar.CharName)
			serverpackets.TradeOtherDone(client)
		}
		if exchange.Recipient.Completed == exchange.Sender.Completed {
			log.Println("Обмен завершен успешно")
			serverpackets.TradeOK(client, player2)
			//Теперь сделаем физическую передачу предметов от персонажа к персонажу
			cplayer, toplayer := trade.TradeAddInventory(client, player2, exchange)

			buffer := packets.Get()
			defer packets.Put(buffer)

			for _, item := range cplayer {
				log.Println(item)
				ut1 := utils.GetPacketByte()
				ut1.SetData(serverpackets.InventoryUpdate(client, item, models.UpdateTypeModify))
				client.EncryptAndSend(ut1.GetData())
			}

			for _, item := range toplayer {
				log.Println(item)

				ut1 := utils.GetPacketByte()
				ut1.SetData(serverpackets.InventoryUpdate(player2, item, models.UpdateTypeModify))
				player2.EncryptAndSend(ut1.GetData())
			}

			if ok = trade.TradeUserClear(client); !ok {
				log.Println("Трейд не найден")
				return
			}
		}
	} else if response == 0 {
		serverpackets.TradeCancel(client, player2)
		if ok = trade.TradeUserClear(client); !ok {
			log.Println("Трейд не найден")
			return
		}
	}

}
