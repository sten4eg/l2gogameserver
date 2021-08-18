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

	var selectedItem items.Item

	for _, v := range myItems {
		if v.ObjId == objId {
			selectedItem = v
			break
		}
	}

	if selectedItem.IsEquipped() == 1 {
		unEquipAndRecord(selectedItem, myItems)
	} else {
		equipItemAndRecord(selectedItem, myItems)
	}

	items.SaveInventoryInDB(conn, myItems)

	serverpackets.NewInventoryUpdate(client, myItems)

	client.CurrentChar.Paperdoll = items.RestoreVisibleInventory(client.CurrentChar.CharId, conn)

	serverpackets.NewUserInfo(client.CurrentChar, client)

	client.SentToSend()
}

func unEquipAndRecord(item items.Item, myItems []items.Item) {
	switch item.Bodypart {
	case items.SlotRHand:
		setPaperdollItem(items.PAPERDOLL_RHAND, nil, myItems)
	case items.SlotLegs:
		setPaperdollItem(items.PAPERDOLL_LEGS, nil, myItems)
	}
}

// equipItemAndRecord
func equipItemAndRecord(item items.Item, myItems []items.Item) {
	switch item.Bodypart {
	case items.SlotRHand:
		setPaperdollItem(items.PAPERDOLL_RHAND, &item, myItems)
	case items.SlotLegs:
		setPaperdollItem(items.PAPERDOLL_LEGS, &item, myItems)
	}
}

func setPaperdollItem(slot uint8, item *items.Item, myItems []items.Item) {

	if item == nil {
		for i, v := range myItems {
			if v.LocData == int32(slot) {
				v.LocData = getFirstEmptySlot(myItems)
				v.Loc = items.Inventory
				myItems[i] = v
				break
			}
		}
		return
	}

	var old items.Item
	var k int
	var keyCurrentItem int
	for i, v := range myItems {
		if v.LocData == int32(slot) && v.Loc == items.Paperdoll { // todo if locdata or slot == 0
			k = i
			old = v
		}

		if v == *item {
			keyCurrentItem = i
		}

	}

	if old.Id != 0 {
		old.Loc = items.Inventory
		old.LocData = item.LocData
		myItems[k] = old
		item.LocData = int32(slot)
		item.Loc = items.Paperdoll
	} else {
		item.LocData = int32(slot)
		item.Loc = items.Paperdoll
	}

	myItems[keyCurrentItem] = *item
}

func getFirstEmptySlot(myItems []items.Item) int32 {
	var max int32
	for _, v := range myItems {
		if v.LocData > max {
			max = v.LocData
		}
	}

	var i int32
	for i = 0; i < max; i++ {
		flag := false
		for _, q := range myItems {
			if q.LocData == i && q.Loc != items.Paperdoll {
				flag = true
			}
		}

		if !flag {
			return i
		}
	}

	return max + 1
}
