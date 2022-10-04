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

const BuyItemBatchLength = 40

func SetPrivateStoreListBuy(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	count := reader.ReadInt32()
	//TODO 3 проверка - проверить что в массиве не осталось предметов (возможно не правильно сделано, проверить).
	if count < 1 || count > 250 || len(data)-int(count)*BuyItemBatchLength > BuyItemBatchLength {
		return
	}

	items := make([]Item, count)

	for i := int32(0); i < count; i++ {
		itemId := reader.ReadInt32()
		_ = reader.ReadInt32()
		cnt := reader.ReadInt64()
		price := reader.ReadInt64()

		if itemId < 1 || cnt < 1 || price < 0 {
			items = nil
			return
		}
		_ = reader.ReadInt32()
		_ = reader.ReadInt32()
		_ = reader.ReadInt32()
		_ = reader.ReadInt32()

		items[i] = Item{itemId: itemId, count: cnt, price: price}
	}

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	if items == nil {
		character.SetPrivateStoreType(privateStoreType.NONE)
		broadcast.BroadcastUserInfo(client)
		character.SendSysMsg(sysmsg.YouHaveExceededQuantityThatCanBeInputted)
		return
	}

	//TODO if (!player.getAccessLevel().allowTransaction())

	//TODO if (AttackStanceTaskManager.getInstance().hasAttackStanceTask(player) || player.isInDuel())

	//TODO if (player.isInsideZone(ZoneId.NO_STORE))

	tradeList := character.GetBuyList()
	tradeList.Clear()

	if len(items) > 3 { //TODO заменить на player.getPrivateBuyStoreLimit()
		pkg := serverpackets.PrivateStoreManageListBuy(character)
		client.SendBuf(pkg)

	}

	var totalCost int64
	for _, item := range items {
		if !item.addToTradeListByItemId(tradeList, character) {
			//Читер
			return
		}

		totalCost += item.getCost()
		if totalCost > config.MaxAdena {
			//Читер
			return
		}
	}

	if totalCost > character.GetInventory().GetAdenaCount() {
		pkg := serverpackets.PrivateStoreManageListBuy(character)
		client.SendBuf(pkg)
		character.SendSysMsg(sysmsg.ThePurchasePriceIsHigherThanMoney)
		return
	}

	if !character.IsSittings() {
		ChangeWaitType(client)
	}
	character.SetPrivateStoreType(privateStoreType.BUY)
	broadcast.BroadcastUserInfo(client)
	pkg := serverpackets.PrivateStoreMsgBuy(character)
	client.SendBuf(pkg)
}
