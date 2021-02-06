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
		itm = unEquipedAndRecord(ii, myItems, client, conn)
	} else {
		itm = equipItemAndRecord(ii, myItems, client, conn)
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

func unEquipedAndRecord(item items.Item, myItems []items.Item, client *models.Client, conn *pgx.Conn) items.Item {
	var i items.Item
	switch item.Bodypart {
	case 128: // rHand
		i = setPaperdollItem(items.PAPERDOLL_RHAND, nil, myItems, client, conn)
	}

	return i
}

func equipItemAndRecord(item items.Item, myItems []items.Item, client *models.Client, conn *pgx.Conn) items.Item {
	var i items.Item
	switch item.Bodypart {
	case items.SlotRHand: // rHand
		_ = setPaperdollItem(items.PAPERDOLL_RHAND, nil, myItems, client, conn)
		i = setPaperdollItem(items.PAPERDOLL_RHAND, &item, myItems, client, conn)
	}

	return i
}

func setPaperdollItem(slot uint8, item *items.Item, myItems []items.Item, client *models.Client, conn *pgx.Conn) items.Item {

	var old items.Item
	var k int
	for i, v := range myItems {
		if v.LocData == int32(slot) { // todo if locdata or slot == 0
			k = i
			old = v
		}
	}

	if old.Id != 0 {
		old.Loc = "INVENTORY"
		old.LocData = items.GetEmptySlot(client.CurrentChar.CharId, conn)
		items.SaveSlotInDB(conn, old)
		myItems[k] = old
		return old
	} else {
		if item == nil {
			return items.Item{}
		}
		item.LocData = int32(slot)
		item.Loc = "PAPERDOLL"
	}

	return *item
}
