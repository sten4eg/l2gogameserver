package clientpackets

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestBuyItem(client interfaces.NewClientCtxInterface, data []byte, db *sql.DB) {
	var packet = packets.NewReader(data)

	listID := packet.ReadInt32()
	_ = listID
	size := packet.ReadInt32()
	// TODO: тут не помещает защита от большого size пакета потому что могут флудить :D

	for i := 0; i < int(size); i++ {
		itemID := packet.ReadInt32()
		count := packet.ReadInt64()
		item, ok := items.GetItemInfo(int(itemID))
		if !ok {
			logger.LogError("[BuyItem] Item Not Found ID:", itemID)
			return
		}
		client.GetAccount().GetCurrentChar().GetInventory().AddItem2(itemID, int(count), item.IsStackable(), db)
	}

	pkg := serverpackets.ItemList(client.GetAccount().GetCurrentChar())
	client.EncryptAndSend(pkg)

}
