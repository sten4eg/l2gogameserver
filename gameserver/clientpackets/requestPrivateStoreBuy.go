package clientpackets

import (
	"database/sql"
	"fmt"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestPrivateStoreBuy(client interfaces.ReciverAndSender, data []byte, db *sql.DB) {
	reader := packets.NewReader(data)

	storeCharacterId := reader.ReadInt32()
	count := reader.ReadInt32()
	//TODO 3 проверка - проверить что в массиве не осталось предметов (возможно не правильно сделано, проверить).
	if count <= 0 || count > 250 || len(data)-int(count)*SellItemBatchLength > SellItemBatchLength { //TODO заменить на вызовы методов
		return
	}

	items := make([]interfaces.ItemRequestInterface, count)

	for i := int32(0); i < count; i++ {
		objectId := reader.ReadInt32()
		cnt := reader.ReadInt64()
		price := reader.ReadInt64()

		if objectId < 1 || cnt < 1 || price < 0 {
			items = nil
			return
		}

		items[i] = models.NewItemRequestWithoutItemId(objectId, cnt, price)
	}

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	if items == nil {
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)
		return
	}

	//TODO floodProtectors

	storeCharacter := character.GetCurrentRegion().GetCharacterInRegions(storeCharacterId)
	if storeCharacter == nil {
		return
	}

	//TODO if (player.isCursedWeaponEquipped())

	ox, oy, oz := character.GetXYZ()
	mx, my, mz := storeCharacter.GetXYZ()
	if models.CalculateDistance(ox, oy, oz, mx, my, mz, true, false) > 150.0 {
		return
	}

	//TODO if ((player.getInstanceId() != storeCharacter.getInstanceId()) && (player.getInstanceId() != -1))

	if !(storeCharacter.GetPrivateStoreType() == privateStoreType.SELL || storeCharacter.GetPrivateStoreType() == privateStoreType.PACKAGE_SELL) {
		return
	}

	storeList := storeCharacter.GetSellList()
	if storeList == nil {
		return
	}

	//TODO if (!player.getAccessLevel().allowTransaction())

	if storeCharacter.GetPrivateStoreType() == privateStoreType.PACKAGE_SELL && len(storeList.GetItems()) > len(items) {
		//TODO Читер забанить
		return
	}

	result := storeList.PrivateStoreBuy(character, items, db)
	if result > 0 {
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)
		if result > 1 {
			fmt.Println("failed")
		}
		return
	}

	if len(storeList.GetItems()) == 0 {
		storeCharacter.SetPrivateStoreType(privateStoreType.NONE)
		broadcast.BroadcastUserInfo(storeCharacter)
	}

}
