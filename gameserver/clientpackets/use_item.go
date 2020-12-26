package clientpackets

import (
	"github.com/jackc/pgx"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
)

func NewUseItem(data []byte, client *models.Client, conn *pgx.Conn) {
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	ctrlPressed := packet.ReadInt32() != 0
	log.Print(objId, ctrlPressed)

	myItems := items.GetMyItems(client.CurrentChar.CharId, conn)

	var ii items.Item
	var q int
	for i, v := range myItems {
		if v.Id == objId {
			ii = v
			q = i
			break
		}
	}

	var itm items.Item
	if ii.IsEquipped() == 1 {

	} else {
		itm = equipItemAndRecord(ii)
	}
	myItems[q] = itm

	items.SaveInventoryInDB(conn, myItems)

	dataq := serverpackets.NewInventoryUpdate(myItems)
	client.SimpleSend(dataq, true)

	//items.UseEquippableItem(client.CurrentChar.Inventory)
}

func equipItemAndRecord(item items.Item) items.Item {
	var i items.Item
	switch item.Bodypart {
	case 128: // rHand
		i = setPaperdollItem(items.PAPERDOLL_RHAND, item)
	}

	return i
}

func setPaperdollItem(slot uint8, item items.Item) items.Item {
	item.LocData = 5
	item.Loc = "PAPERDOLL"
	return item
}
