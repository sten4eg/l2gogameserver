package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ExStorageMaxCount(clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return []byte{}
	}
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0x2F)

	buffer.WriteD(int32(client.CurrentChar.GetInventoryLimit())) // _inventory Limit
	buffer.WriteD(0)                                             // _warehouse Limit
	buffer.WriteD(0)                                             // _clan Limit
	buffer.WriteD(0)                                             // _privateSell
	buffer.WriteD(0)                                             // _privateBuy
	buffer.WriteD(0)                                             // _recipeD (dworf)
	buffer.WriteD(0)                                             //_recipe
	buffer.WriteD(0)                                             // _inventoryExtraSlots
	buffer.WriteD(0)                                             // _inventoryQuestItems

	defer packets.Put(buffer)
	return buffer.Bytes()

}
