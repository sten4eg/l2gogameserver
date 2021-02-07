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

	for _, v := range myItems {
		if v.Id == objId {
			ii = v

			break
		}
	}

	if ii.IsEquipped() == 1 {
		unEquipedAndRecord(ii, myItems)
	} else {
		equipItemAndRecord(ii, myItems)
	}

	items.SaveInventoryInDB(conn, myItems)

	dataq := serverpackets.NewInventoryUpdate(myItems)
	client.SimpleSend(dataq, true)

	client.CurrentChar.Paperdoll = items.RestoreVisibleInventory(client.CurrentChar.CharId, conn)
	pkg := serverpackets.NewUserInfo(client.CurrentChar)
	err := client.Send(pkg, true)
	if err != nil {
		log.Println(err)
	}
}

func unEquipedAndRecord(item items.Item, myItems []items.Item) {

	switch item.Bodypart {
	case 128: // rHand
		setPaperdollItem(items.PAPERDOLL_RHAND, nil, myItems)
	}
}

func equipItemAndRecord(item items.Item, myItems []items.Item) {
	switch item.Bodypart {
	case items.SlotRHand: // rHand
		setPaperdollItem(items.PAPERDOLL_RHAND, &item, myItems)
	}
}

func getEmtpySlot(items []items.Item) {
	//slot := 0
	//for _,v := range items {
	//	if v.Loc == "INVENTORY"
	//}
}
func setPaperdollItem(slot uint8, item *items.Item, myItems []items.Item) {

	if item == nil {
		for i, v := range myItems {
			if v.LocData == int32(slot) {
				v.LocData = 32
				v.Loc = "INVENTORY"
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
		if v.LocData == int32(slot) { // todo if locdata or slot == 0
			k = i
			old = v
		}

		if v == *item {
			keyCurrentItem = i
		}

	}

	if old.Id != 0 {
		old.Loc = "INVENTORY"
		old.LocData = item.LocData
		myItems[k] = old
		item.LocData = int32(slot)
		item.Loc = "PAPERDOLL"
	} else {
		item.LocData = int32(slot)
		item.Loc = "PAPERDOLL"
	}

	myItems[keyCurrentItem] = *item
}
