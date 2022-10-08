package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ExStorageMaxCount(character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0x2F)

	buffer.WriteD(int32(character.GetInventoryLimit())) // _inventory Limit
	buffer.WriteD(0)                                    // _warehouse Limit
	buffer.WriteD(0)                                    // _clan Limit
	buffer.WriteD(3)                                    // _privateSell
	buffer.WriteD(3)                                    // _privateBuy
	buffer.WriteD(0)                                    // _recipeD (dworf)
	buffer.WriteD(0)                                    //_recipe
	buffer.WriteD(0)                                    // _inventoryExtraSlots
	buffer.WriteD(0)                                    // _inventoryQuestItems

	defer packets.Put(buffer)
	return buffer.Bytes()

}
