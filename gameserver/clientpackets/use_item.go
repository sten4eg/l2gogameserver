package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewUseItem(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	ctrlPressed := packet.ReadInt32() != 0
	_ = ctrlPressed

	var selectedItem items.MyItem

	for _, v := range client.CurrentChar.Inventory {
		if v.ObjId == objId {
			selectedItem = v
			break
		}
	}

	if selectedItem.IsEquipped() == 1 {
		unEquipAndRecord(selectedItem, client.CurrentChar.Inventory)
	} else {
		equipItemAndRecord(selectedItem, client.CurrentChar.Inventory)
	}

	items.SaveInventoryInDB(client.CurrentChar.Inventory)

	serverpackets.NewInventoryUpdate(client, items.UpdateTypeModify)

	client.CurrentChar.Paperdoll = items.RestoreVisibleInventory(client.CurrentChar.CharId)

	serverpackets.NewUserInfo(client.CurrentChar, client)

	client.SentToSend()
}

func unEquipAndRecord(item items.MyItem, myItems []items.MyItem) {
	switch item.SlotBitType {
	case items.SlotRHand:
		setPaperdollItem(items.PAPERDOLL_RHAND, nil, myItems)
	case items.SlotLegs:
		setPaperdollItem(items.PAPERDOLL_LEGS, nil, myItems)
	}
}

// equipItemAndRecord
func equipItemAndRecord(item items.MyItem, myItems []items.MyItem) {
	switch item.SlotBitType {
	case items.SlotRHand:
		setPaperdollItem(items.PAPERDOLL_RHAND, &item, myItems)
	case items.SlotLegs:
		setPaperdollItem(items.PAPERDOLL_LEGS, &item, myItems)
	}
}

func setPaperdollItem(slot uint8, item *items.MyItem, myItems []items.MyItem) {

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

	var old items.MyItem
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

func getFirstEmptySlot(myItems []items.MyItem) int32 {
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
