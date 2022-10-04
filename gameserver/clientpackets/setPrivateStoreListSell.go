package clientpackets

import (
	"l2gogameserver/config"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

const SellItemBatchLength = 20

type Item struct {
	itemId int32
	count  int64
	price  int64
}

func (i *Item) addToTradeListByObjectId(list interfaces.TradeListInterface, char interfaces.CharacterI) bool {
	if (config.MaxAdena / i.count) < i.price {
		return false
	}

	list.AddItem(i.itemId, i.count, char, i.price)
	return true
}

func (i *Item) addToTradeListByItemId(list interfaces.TradeListInterface, char interfaces.CharacterI) bool {
	if (config.MaxAdena / i.count) < i.price {
		return false
	}
	list.AddItemByItemId(i.itemId, i.count, i.price)
	return true
}

func (i *Item) getPrice() int64 {
	return i.count * i.price
}

func SetPrivateStoreListSell(client interfaces.ReciverAndSender, data []byte) {
	packet := packets.NewReader(data)

	packageSale := packet.ReadInt32() == 1
	count := packet.ReadInt32()
	//TODO 3 проверка - проверить что в массиве не осталось предметов (возможно не правильно сделано, проверить).
	if count < 1 || count > 250 || len(data)-int(count)*SellItemBatchLength > SellItemBatchLength { //TODO заменить на вызовы методов
		return
	}

	items := make([]Item, count)

	for i := int32(0); i < count; i++ {
		itemId := packet.ReadInt32()
		cnt := packet.ReadInt64()
		price := packet.ReadInt64()

		if itemId < 1 || cnt < 1 || price < 0 {
			items = nil
			return
		}

		items[i] = Item{itemId: itemId, count: cnt, price: price}
	}

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	if items == nil {
		character.EncryptAndSend(sysmsg.SystemMessage(sysmsg.IncorrectItemCount))
		character.SetPrivateStoreType(privateStoreType.NONE)
		broadcast.BroadcastUserInfo(client)
		return
	}

	//TODO проверка на AccessLevel

	if len(items) > 3 { //TODO > getPrivateSellStoreLimit()
		pkg := serverpackets.PrivateStoreManageListSell(character, packageSale)
		character.SendBuf(pkg)
		character.EncryptAndSend(sysmsg.SystemMessage(sysmsg.YouHaveExceededQuantityThatCanBeInputted))
		return
	}

	tradeList := character.GetSellList()
	tradeList.Clear()
	tradeList.SetPackaged(packageSale)

	totalCost := character.GetInventory().GetAdenaCount()
	for _, item := range items {
		if !item.addToTradeListByObjectId(tradeList, character) {
			return
		}

		totalCost += item.getPrice()

		if totalCost > config.MaxAdena {
			return
		}

	}

	if !character.IsSittings() {
		ChangeWaitType(client)
	}
	if packageSale {
		character.SetPrivateStoreType(privateStoreType.PACKAGE_SELL)
	} else {
		character.SetPrivateStoreType(privateStoreType.SELL)
	}

	broadcast.BroadcastUserInfo(client)

	if packageSale {
		pkg := serverpackets.ExPrivateStoreSetWholeMsg(character, character.GetSellList().GetTitle())
		broadcast.BroadCastBufferToAroundPlayers(client, pkg)
	} else {
		pkg := serverpackets.PrivateStoreMsgSell(character)
		broadcast.BroadCastBufferToAroundPlayers(client, pkg)
	}
}
