package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

const limit = 125 // client limit
func RequestSaveInventoryOrder(clientI interfaces.NewClientCtxInterface, data []byte) {

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

	newItems := clientI.GetAccount().GetCurrentChar().GetInventory().GetItems()
	//todo переделать без n^2
	for _, io := range newOrder {
		for _, i := range newItems {
			if io.ObjId == i.GetObjectId() && i.GetLocation() == models.InventoryLoc {
				i.SetLocData(io.Order)
			}
		}
	}
}
