package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func PrivateStoreListBuy(character, storeCharacter interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()
	storeCharacter.GetSellList().UpdateItems()
	items := storeCharacter.GetBuyList().GetAvailableItems(character.GetInventory())

	buffer.WriteSingleByte(0xbe)
	buffer.WriteD(storeCharacter.GetObjectId())
	buffer.WriteQ(character.GetInventory().GetAdenaCount())

	buffer.WriteD(int32(len(items)))

	for _, item := range items {
		item.WriteItem(buffer)
		buffer.WriteD(item.GetObjectId())
		buffer.WriteQ(item.GetPrice())
		buffer.WriteQ(int64(item.GetDefaultPrice() * 2))
		buffer.WriteQ(item.GetStoreCount())
	}

	return buffer
}
