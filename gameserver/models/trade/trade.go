package trade

import (
	"l2gogameserver/gameserver/interfaces"
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
	ObjectId  int32 //objectId персонажа
	Completed bool  //true подтверждение сделки
	Client    *models.Character
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

//NewRequestTrade Добавляем в реестр трейдов
func NewRequestTrade(senderI, recipientI interfaces.CharacterI) {
	sender, ok := senderI.(*models.Character)
	if !ok {
		panic("NewRequestTrade sender not client")
	}
	recipient, ok := recipientI.(*models.Character)
	if !ok {
		panic("NewRequestTrade sender not client")
	}

	u := &Exchange{
		Sender: Action{
			ObjectId: sender.ObjectId,
			Client:   sender,
		},
		Recipient: Action{
			ObjectId: recipient.ObjectId,
			Client:   recipient,
		},
		Status: Wait,
		Time:   time.Now(),
	}
	allTrade = append(allTrade, u)
}

//Answer Когда пользователь отвечает "Да" или "нет" на предложение торговать
func Answer(client interfaces.CharacterI) (*Exchange, bool) {
	for _, exchange := range allTrade {
		if exchange.Recipient.ObjectId == client.GetObjectId() {
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

func AddItemTrade(client interfaces.CharacterI, objectId int32, count int64) (*models.MyItem, interfaces.CharacterI, bool) {
	for _, exchange := range allTrade {
		//Проверяем, есть ли предмет в инвентаре
		if exchange.Sender.ObjectId == client.GetObjectId() {
			if item, ok := models.ExistItemObject(client, objectId, count); ok {
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
		} else if exchange.Recipient.ObjectId == client.GetObjectId() {
			if item, ok := models.ExistItemObject(client, objectId, count); ok {
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

//ExistItemTradeObject Проверяет, есть ли в трайде уже добавленный X предмет
func (e *Exchange) ExistItemTradeObject(client interfaces.CharacterI, objectid int32) bool {
	for _, exchanges := range allTrade {
		if exchanges.Sender.ObjectId == client.GetObjectId() {
			for _, exchange := range exchanges.Sender.Items {
				if exchange.ObjId == objectid {
					return true
				}
			}
		}
		if exchanges.Recipient.ObjectId == client.GetObjectId() {
			for _, exchange := range exchanges.Recipient.Items {
				if exchange.ObjId == objectid {
					return true
				}
			}
		}
	}
	return false
}

//FindTrade Пользователь отменил сделку
//Возращает ID персонажей, которые участвовали в торговле
func FindTrade(client interfaces.CharacterI) (interfaces.CharacterI, *Exchange, bool) {
	var playerTo interfaces.CharacterI
	for _, exchange := range allTrade {
		if exchange.Sender.ObjectId == client.GetObjectId() {
			playerTo = exchange.Recipient.Client
			exchange.ChangeStatusTrade(Cancel)
			return playerTo, exchange, true
		} else if exchange.Recipient.ObjectId == client.GetObjectId() {
			playerTo = exchange.Sender.Client
			exchange.ChangeStatusTrade(Cancel)
			return playerTo, exchange, true
		}

	}
	return nil, nil, false
}

//Подтверждение согласия на торг
func TradeOK(client *models.Client) {

}

//Очистка информации трейде
func TradeUserClear(client interfaces.CharacterI) bool {
	for index, exchange := range allTrade {
		if exchange.Sender.ObjectId == client.GetObjectId() || exchange.Recipient.ObjectId == client.GetObjectId() {
			allTrade = append(allTrade[:index], allTrade[index+1:]...)
			return true
		}
	}
	return false
}

//TradeAddInventory Обмен предметами
func TradeAddInventory(clientI, player2I interfaces.CharacterI, exchange *Exchange) ([]*models.MyItem, []*models.MyItem) {
	var allItemUpdateClient []*models.MyItem
	var allItemUpdatePlayer []*models.MyItem

	client, ok := clientI.(*models.Character)
	if !ok {
		panic("TradeAddInventory clientI not character")
	}
	player2, ok := player2I.(*models.Character)
	if !ok {
		panic("TradeAddInventory clientI not character")
	}
	for _, itm := range exchange.Sender.Items {
		if exchange.Sender.ObjectId == client.GetObjectId() {
			log.Println(itm.Name, itm.Count, player2.CharName)
			client.Inventory.RemoveItem(client, itm, itm.Count)
			mi, ok := models.AddInventoryItem(player2, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdateClient = append(allItemUpdateClient, mi)
		} else {
			log.Println(itm.Name, itm.Count, client.CharName)
			player2.Inventory.RemoveItem(player2, itm, itm.Count)
			mi, ok := models.AddInventoryItem(client, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdatePlayer = append(allItemUpdatePlayer, mi)
		}
	}

	for _, itm := range exchange.Recipient.Items {
		if exchange.Sender.ObjectId == client.ObjectId {
			log.Println(itm.Name, itm.Count, client.CharName)
			player2.Inventory.RemoveItem(player2, itm, itm.Count)
			mi, ok := models.AddInventoryItem(client, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdateClient = append(allItemUpdateClient, mi)
		} else {
			log.Println(itm.Name, itm.Count, player2.CharName)
			client.Inventory.RemoveItem(client, itm, itm.Count)
			mi, ok := models.AddInventoryItem(player2, itm, itm.Count)
			if !ok {
				log.Println("НЕ ОК")
			}
			allItemUpdatePlayer = append(allItemUpdatePlayer, mi)
		}
	}
	return allItemUpdateClient, allItemUpdatePlayer
}
