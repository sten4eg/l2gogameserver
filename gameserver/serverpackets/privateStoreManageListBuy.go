package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func PrivateStoreManageListBuy(character interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()
	itemList := character.GetInventory().GetUniqueItems(character, false, true, true)
	buyList := character.GetBuyList().GetItems()

	buffer.WriteSingleByte(0xbd)

	buffer.WriteD(character.GetObjectId())
	buffer.WriteQ(character.GetInventory().GetAdenaCount())

	buffer.WriteD(int32(len(itemList)))
	for _, item := range itemList {
		item.WriteItem(buffer)
		buffer.WriteQ(int64(item.GetDefaultPrice() * 2))
	}

	buffer.WriteD(int32(len(buyList)))
	for _, item := range buyList {
		item.WriteItem(buffer)
		buffer.WriteQ(item.GetPrice())
		buffer.WriteQ(int64(item.GetDefaultPrice() * 2))
		buffer.WriteQ(item.GetCount())
	}

	return buffer
}
