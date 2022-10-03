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

const BatchLength = 20

type Item struct {
	itemId int32
	count  int64
	price  int64
}

func (i *Item) addToTradeList(list interfaces.TradeListInterface, char interfaces.CharacterI) bool {
	if (config.MaxAdena / i.count) < i.price {
		return false
	}

	list.AddItem(i.itemId, i.count, char, i.price)
	return true
}

func (i *Item) getPrice() int64 {
	return i.count * i.price
}

func SetPrivateStoreListSell(client interfaces.ReciverAndSender, data []byte) {
	packet := packets.NewReader(data)

	packageSale := packet.ReadInt32() == 1
	count := packet.ReadInt32()

	if count < 1 || count > 250 || len(data)-int(count)*BatchLength > BatchLength { //TODO заменить на вызовы методов
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

	activeChar := client.GetCurrentChar()
	if activeChar == nil {
		return
	}

	if items == nil {
		activeChar.EncryptAndSend(sysmsg.SystemMessage(sysmsg.IncorrectItemCount))
		activeChar.SetPrivateStoreType(privateStoreType.NONE)
		broadcast.BroadcastUserInfo(client)
	}

	//TODO проверка на AccessLevel

	if len(items) > 3 { //TODO > getPrivateSellStoreLimit()
		pkg := serverpackets.PrivateStoreManageListSell(activeChar, packageSale)
		activeChar.SendBuf(pkg)
		activeChar.EncryptAndSend(sysmsg.SystemMessage(sysmsg.YouHaveExceededQuantityThatCanBeInputted))
		return
	}

	tradeList := activeChar.GetSellList()
	tradeList.Clear()
	tradeList.SetPackaged(packageSale)

	totalCost := activeChar.GetInventory().GetAdenaCount()
	for _, item := range items {
		if !item.addToTradeList(tradeList, activeChar) {
			return
		}

		totalCost += item.getPrice()

		if totalCost > config.MaxAdena {
			return
		}

	}

	if !activeChar.IsSittings() {
		ChangeWaitType(client)
	}
	if packageSale {
		activeChar.SetPrivateStoreType(privateStoreType.PACKAGE_SELL)
	} else {
		activeChar.SetPrivateStoreType(privateStoreType.SELL)
	}

	broadcast.BroadcastUserInfo(client)

	if packageSale {
		pkg := serverpackets.ExPrivateStoreSetWholeMsg(activeChar, activeChar.GetSellList().GetTitle())
		broadcast.BroadCastBufferToAroundPlayers(client, pkg)
	} else {
		pkg := serverpackets.PrivateStoreMsgSell(activeChar)
		broadcast.BroadCastBufferToAroundPlayers(client, pkg)
	}
}
