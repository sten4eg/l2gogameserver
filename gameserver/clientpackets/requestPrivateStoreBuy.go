package clientpackets

import (
	"fmt"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestPrivateStoreBuy(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	storePlayerId := reader.ReadInt32()
	count := reader.ReadInt32()

	if count <= 0 || count > 250 || len(data)-int(count)*BatchLength > BatchLength { //TODO заменить на вызовы методов
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

		items[i] = models.NewItemRequest(objectId, cnt, price)
	}

	activeChar := client.GetCurrentChar()
	if activeChar == nil {
		return
	}

	if items == nil {
		pkg := serverpackets.ActionFailed(client)
		client.EncryptAndSend(pkg)
		return
	}

	//TODO floodProtectors

	storeCharacter := getCharacter(activeChar, storePlayerId)
	if storeCharacter == nil {
		return
	}

	//TODO if (player.isCursedWeaponEquipped())

	//TODO if (!player.isInsideRadius(storePlayer, INTERACTION_DISTANCE, true, false))

	//TODO if ((player.getInstanceId() != storePlayer.getInstanceId()) && (player.getInstanceId() != -1))

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

	result := storeList.PrivateStoreBuy(activeChar, items)
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
		broadcast.BroadcastUserInfo(client)
	}

}

func getCharacter(character interfaces.CharacterI, objectId int32) interfaces.CharacterI {
	for _, region := range character.GetCurrentRegion().GetNeighbors() {
		char, ok := region.GetChar(objectId)
		if ok {
			return char
		}
	}
	return nil
}
