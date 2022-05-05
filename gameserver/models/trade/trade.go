package trade

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
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
		logger.Error.Panicln("NewRequestTrade sender not client")
	}
	recipient, ok := recipientI.(*models.Character)
	if !ok {
		logger.Error.Panicln("NewRequestTrade sender not client")
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
					logger.Info.Println("Вы уже добавили этот предмет в инвентарь")
					return item, exchange.Recipient.Client, false
				}
				item.Count = count
				exchange.Sender.Items = append(exchange.Sender.Items, item)
				return item, exchange.Recipient.Client, true
			} else {
				logger.Info.Println("Не найден предмет или неверное количество предметов")
			}
		} else if exchange.Recipient.ObjectId == client.GetObjectId() {
			if item, ok := models.ExistItemObject(client, objectId, count); ok {
				if exchange.ExistItemTradeObject(client, objectId) {
					logger.Info.Println("Вы уже добавили этот предмет в инвентарь")
					return item, exchange.Sender.Client, false
				}
				item.Count = count
				exchange.Recipient.Items = append(exchange.Recipient.Items, item)
				return item, exchange.Sender.Client, true
			} else {
				logger.Info.Println("Не найден предмет или неверное количество предметов")
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
				if exchange.ObjectId == objectid {
					return true
				}
			}
		}
		if exchanges.Recipient.ObjectId == client.GetObjectId() {
			for _, exchange := range exchanges.Recipient.Items {
				if exchange.ObjectId == objectid {
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

// UserClear Очистка информации трейде
func UserClear(client interfaces.CharacterI) bool {
	for index, exchange := range allTrade {
		if exchange.Sender.ObjectId == client.GetObjectId() || exchange.Recipient.ObjectId == client.GetObjectId() {
			allTrade = append(allTrade[:index], allTrade[index+1:]...)
			return true
		}
	}
	return false
}

//Информация о том какие предметы будем удалять из инвентарей, и какие предметы будем добавлять
type UpdateTradeData struct {
	Player     interfaces.CharacterI //Над каким персонажем производятся действия
	Item       models.MyItem         //Предмет
	Count      int64                 //Кол-во
	UpdateType int16
}

//TradeAddInventory Обмен предметами
func TradeAddInventory(clientI, player2I interfaces.CharacterI, exchange *Exchange) []UpdateTradeData {
	var UpdateInfo []UpdateTradeData

	client, ok := clientI.(*models.Character)
	if !ok {
		logger.Error.Panicln("TradeAddInventory clientI not character")
	}
	player2, ok := player2I.(*models.Character)
	if !ok {
		logger.Error.Panicln("TradeAddInventory clientI not character")
	}
	for _, itm := range exchange.Sender.Items {
		if exchange.Sender.ObjectId == client.GetObjectId() {
			UpdateInfo = removeAndAdd(client, player2, itm, itm.Count)
		} else {
			UpdateInfo = removeAndAdd(player2, client, itm, itm.Count)
		}
	}

	for _, itm := range exchange.Recipient.Items {
		if exchange.Sender.ObjectId == client.ObjectId {
			UpdateInfo = removeAndAdd(player2, client, itm, itm.Count)
		} else {
			UpdateInfo = removeAndAdd(client, player2, itm, itm.Count)
		}
	}
	return UpdateInfo
}

//Удаление и добавление в массив
func removeAndAdd(client, player2 *models.Character, itm *models.MyItem, count int64) []UpdateTradeData {
	var UpdateInfo []UpdateTradeData

	item, count, updtype, ok := models.RemoveItem(client, itm, itm.Count)

	UpdateInfo = append(UpdateInfo, UpdateTradeData{
		Player:     client,
		Item:       item,
		Count:      count,
		UpdateType: updtype,
	})

	item, count, updtype, ok = models.AddInventoryItem(player2, *itm, itm.Count)
	if !ok {
		logger.Info.Println("НЕ ОК")
	}
	UpdateInfo = append(UpdateInfo, UpdateTradeData{
		Player:     player2,
		Item:       item,
		Count:      count,
		UpdateType: updtype,
	})
	return UpdateInfo
}
