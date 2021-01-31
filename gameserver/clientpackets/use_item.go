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
		itm = unEquipedAndRecord(ii, myItems)
	} else {
		itm = equipItemAndRecord(ii, myItems)
	}
	myItems[q] = itm

	items.SaveInventoryInDB(conn, myItems)

	dataq := serverpackets.NewInventoryUpdate(myItems)
	client.SimpleSend(dataq, true)

	client.CurrentChar.Paperdoll = items.RestoreVisibleInventory(client.CurrentChar.CharId, conn)
	pkg := serverpackets.NewUserInfo(client.CurrentChar)
	err := client.Send(pkg, true)
	if err != nil {
		log.Println(err)
	}
	//items.UseEquippableItem(client.CurrentChar.Inventory)
}

func unEquipedAndRecord(item items.Item, myItems []items.Item) items.Item {
	var i items.Item
	switch item.Bodypart {
	case 128: // rHand
		i = setPaperdollItem(items.PAPERDOLL_RHAND, nil, myItems)
	}

	return i
}

func equipItemAndRecord(item items.Item, myItems []items.Item) items.Item {
	var i items.Item
	switch item.Bodypart {
	case 128: // rHand
		i = setPaperdollItem(items.PAPERDOLL_RHAND, &item, myItems)
	}

	return i
}

func setPaperdollItem(slot uint8, item *items.Item, myItems []items.Item) items.Item {

	var old items.Item
	for _, v := range myItems {
		if v.LocData == int32(slot) { // todo if locdata or slot == 0
			old = v
		}
	}

	if old.Id != 0 {
		old.Loc = "INVENTORY"
		old.LocData = 23
		return old
	} else {
		item.LocData = int32(slot)
		item.Loc = "PAPERDOLL"
	}

	return *item
}
