package trade

import (
	"l2gogameserver/gameserver/models"
	"log"
	"time"
)

type StatusTrade int

const (
	Wait        StatusTrade = 0 //Ожидание запроса
	During      StatusTrade = 1 //Идет обмен
	Completed   StatusTrade = 2 //Завершен обмен
	Interrupted StatusTrade = 3 //Ошибка
	Cancel      StatusTrade = 4 //Отмена обмена
)

type Action struct {
	UserID    int32 //ID персонажа
	Completed bool  //true подтверждение сделки
	Client    *models.Client
	Items     []*models.MyItem
}

//Обмен
type Exchange struct {
	Sender    Action //Отправитель
	Recipient Action //Получатель
	Status    StatusTrade
	Time      time.Time
}

var allTrade []*Exchange

//Добавляем в реестр трейдов
func NewRequestTrade(sender, recipient *models.Client) {
	u := &Exchange{
		Sender: Action{
			UserID: sender.CurrentChar.ObjectId,
			Client: sender,
		},
		Recipient: Action{
			UserID: recipient.CurrentChar.ObjectId,
			Client: recipient,
		},
		Status: Wait,
		Time:   time.Now(),
	}
	allTrade = append(allTrade, u)
}

//Когда пользователь отвечает "Да" или "нет" на предложение торговать
func TradeAnswer(client *models.Client, response int32) (*Exchange, bool) {
	for _, exchange := range allTrade {
		if exchange.Recipient.UserID == client.CurrentChar.ObjectId {
			exchange.ChangeStatusTrade(During)
			//Теперь отправляем пакет на "открытие" окна обмена
			return exchange, true
		}
	}
	return nil, false
}

func (e *Exchange) ChangeStatusTrade(st StatusTrade) {
	e.Status = st
}

func AddItemTrade(client *models.Client, objectId int32, count int64) (*models.MyItem, *models.Client, bool) {
	for _, exchange := range allTrade {
		//Проверяем, есть ли предмет в инвентаре
		if exchange.Sender.UserID == client.CurrentChar.ObjectId {
			if item, ok := models.ExistItemObject(client.CurrentChar, objectId, count); ok {
				//Проверяем, уже добавили предмет в трейд на обмене
				if exchange.ExistItemTradeObject(client, objectId) {
					log.Println("Вы уже добавили этот предмет в инвентарь")
					return item, exchange.Recipient.Client, false
				}
				item.Count = count
				exchange.Sender.Items = append(exchange.Sender.Items, item)
				return item, exchange.Recipient.Client, true
			} else {
				log.Println("Не найден предмет или неверное количество предметов")
			}
		} else if exchange.Recipient.UserID == client.CurrentChar.ObjectId {
			if item, ok := models.ExistItemObject(client.CurrentChar, objectId, count); ok {
				if exchange.ExistItemTradeObject(client, objectId) {
					log.Println("Вы уже добавили этот предмет в инвентарь")
					return item, exchange.Sender.Client, false
				}
				item.Count = count
				exchange.Recipient.Items = append(exchange.Recipient.Items, item)
				return item, exchange.Sender.Client, true
			} else {
				log.Println("Не найден предмет или неверное количество предметов")
			}
		}
	}
	return &models.MyItem{}, nil, false
}

//Проверяет, есть ли в трайде уже добавленный X предмет
func (e *Exchange) ExistItemTradeObject(client *models.Client, objectid int32) bool {
	for _, exchanges := range allTrade {
		if exchanges.Sender.UserID == client.CurrentChar.ObjectId {
			for _, exchange := range exchanges.Sender.Items {
				if exchange.ObjId == objectid {
					return true
				}
			}
		}
		if exchanges.Recipient.UserID == client.CurrentChar.ObjectId {
			for _, exchange := range exchanges.Recipient.Items {
				if exchange.ObjId == objectid {
					return true
				}
			}
		}
	}
	return false
}

//Пользователь отменил сделку
//Возращает ID персонажей, которые участвовали в торговле
func FindTrade(client *models.Client) (*models.Client, *Exchange, bool) {
	var playerTo *models.Client
	for _, exchange := range allTrade {
		if exchange.Sender.UserID == client.CurrentChar.ObjectId {
			playerTo = exchange.Recipient.Client
		} else if exchange.Recipient.UserID == client.CurrentChar.ObjectId {
			playerTo = exchange.Sender.Client
		}
		exchange.ChangeStatusTrade(Cancel)
		return playerTo, exchange, true
	}
	return nil, nil, false
}

//Подтверждение согласия на торг
func TradeOK(client *models.Client) {

}

//Очистка информации трейде
func TradeUserClear(client *models.Client) bool {
	for index, exchange := range allTrade {
		if exchange.Sender.UserID == client.CurrentChar.ObjectId || exchange.Recipient.UserID == client.CurrentChar.ObjectId {
			allTrade = append(allTrade[:index], allTrade[index+1:]...)
			return true
		}
	}
	return false
}

//Обмен предметами
func TradeAddInventory(client, player2 *models.Client, exchange *Exchange) ([]*models.MyItem, []*models.MyItem) {
	var allItemUpdateClient []*models.MyItem
	var allItemUpdatePlayer []*models.MyItem

	for _, itm := range exchange.Sender.Items {
		if exchange.Sender.UserID == client.CurrentChar.ObjectId {
			log.Println(itm.Name, itm.Count, player2.CurrentChar.CharName)
			client.CurrentChar.Inventory.RemoveItem(client, itm, itm.Count)
			mi, ok := models.AddInventoryItem(player2, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdateClient = append(allItemUpdateClient, mi)
		} else {
			log.Println(itm.Name, itm.Count, client.CurrentChar.CharName)
			player2.CurrentChar.Inventory.RemoveItem(player2, itm, itm.Count)
			mi, ok := models.AddInventoryItem(client, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdatePlayer = append(allItemUpdatePlayer, mi)
		}
	}

	for _, itm := range exchange.Recipient.Items {
		if exchange.Sender.UserID == client.CurrentChar.ObjectId {
			log.Println(itm.Name, itm.Count, client.CurrentChar.CharName)
			player2.CurrentChar.Inventory.RemoveItem(player2, itm, itm.Count)
			mi, ok := models.AddInventoryItem(client, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdateClient = append(allItemUpdateClient, mi)
		} else {
			log.Println(itm.Name, itm.Count, player2.CurrentChar.CharName)
			client.CurrentChar.Inventory.RemoveItem(client, itm, itm.Count)
			mi, ok := models.AddInventoryItem(player2, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdatePlayer = append(allItemUpdatePlayer, mi)
		}
	}
	return allItemUpdateClient, allItemUpdatePlayer
}
