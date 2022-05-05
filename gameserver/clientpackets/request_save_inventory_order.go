package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

const limit = 125 // client limit
func RequestSaveInventoryOrder(clientI interfaces.ReciverAndSender, data []byte) {
	client, ok := clientI.(*models.Client)
	if !ok {
		return
	}

	var reader = packets.NewReader(data[2:])
	size := reader.ReadInt32()
	if size > limit {
		size = limit
	}

	type InventoryOrder struct {
		ObjId int32
		Order int32
	}
	newOrder := make([]InventoryOrder, 0, size)

	for i := int32(0); i < size; i++ {
		var io InventoryOrder
		io.ObjId = reader.ReadInt32()
		io.Order = reader.ReadInt32()
		newOrder = append(newOrder, io)
	}

	items := client.CurrentChar.Inventory.Items

	//todo переделать без n^2
	for _, io := range newOrder {
		for i := range items {
			item := &items[i]
			if io.ObjId == item.ObjectId && item.Location == models.InventoryLoc {
				items[i].LocData = io.Order
			}
		}
	}
}
