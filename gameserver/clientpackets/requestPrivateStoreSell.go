package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestPrivateStoreSell(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	storeCharacterId := reader.ReadInt32()
	count := reader.ReadInt32()
	//TODO 3 проверка - проверить что в массиве не осталось предметов (возможно не правильно сделано, проверить).
	if count <= 0 || count > 250 || len(data)-int(count)*BuyItemBatchLength > BuyItemBatchLength {
		return
	}

	items := make([]interfaces.ItemRequestInterface, count)

	for i := int32(0); i < count; i++ {
		objectId := reader.ReadInt32()
		itemId := reader.ReadInt32()

		_ = reader.ReadInt16()
		_ = reader.ReadInt16()

		cnt := reader.ReadInt64()
		price := reader.ReadInt64()

		if objectId < 1 || itemId < 1 || cnt < 1 || price < 0 {
			items = nil
			return
		}
		items[i] = models.NewItemRequest(objectId, itemId, cnt, price)
	}

	player := client.GetCurrentChar()
	if player == nil {
		return
	}

	if items == nil {
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)
	}

	//TODO if (!getClient().getFloodProtectors().getTransaction().tryPerformAction("privatestoresell"))

	storeCharacter := player.GetCurrentRegion().GetCharacterInRegions(storeCharacterId)
	if storeCharacter == nil {
		return
	}

	ox, oy, oz := player.GetXYZ()
	mx, my, mz := storeCharacter.GetXYZ()
	if models.CalculateDistance(ox, oy, oz, mx, my, mz, true, false) > 150.0 {
		return
	}

	//TODO if ((player.getInstanceId() != storeCharacter.getInstanceId()) && (player.getInstanceId() != -1))

	if storeCharacter.GetPrivateStoreType() != privateStoreType.BUY {
		return
	}

	//TODO if (player.isCursedWeaponEquipped())

	storeList := storeCharacter.GetBuyList()
	if storeList == nil {
		return
	}

	//TODO if (!player.getAccessLevel().allowTransaction())

	if !storeList.PrivateStoreSell(player, items) {
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)
		return
	}

	if len(storeList.GetItems()) == 0 {
		storeCharacter.SetPrivateStoreType(privateStoreType.NONE)
		broadcast.BroadcastUserInfo(storeCharacter)
	}
}
