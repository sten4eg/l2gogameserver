package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func ExStorageMaxCount(client *models.Client) {

	client.Buffer.WriteSingleByte(0xFE)
	client.Buffer.WriteH(0x2F)

	client.Buffer.WriteD(int32(client.CurrentChar.GetInventoryLimit())) // _inventory Limit
	client.Buffer.WriteD(0)                                             // _warehouse Limit
	client.Buffer.WriteD(0)                                             // _clan Limit
	client.Buffer.WriteD(0)                                             // _privateSell
	client.Buffer.WriteD(0)                                             // _privateBuy
	client.Buffer.WriteD(0)                                             // _recipeD (dworf)
	client.Buffer.WriteD(0)                                             //_recipe
	client.Buffer.WriteD(0)                                             // _inventoryExtraSlots
	client.Buffer.WriteD(0)                                             // _inventoryQuestItems

	client.SaveAndCryptDataInBufferToSend(true)

}
