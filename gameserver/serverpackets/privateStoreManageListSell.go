package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func PrivateStoreManageListSell(character interfaces.CharacterI, isPackageSale bool) *packets.Buffer {
	buffer := packets.Get()
	character.GetSellList().UpdateItems()
	itemList := character.GetInventory().GetAvailableItems(character.GetSellList(), character)
	sellList := character.GetSellList().GetItems()

	buffer.WriteSingleByte(0xA0)

	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(utils.BoolToInt32(isPackageSale))
	buffer.WriteQ(character.GetInventory().GetAdenaCount())

	buffer.WriteD(int32(len(itemList)))
	for _, item := range itemList {
		item.WriteItem(buffer)
		buffer.WriteQ(int64(item.GetDefaultPrice()) * 2)
	}

	buffer.WriteD(int32(len(sellList)))
	for _, item := range sellList {
		item.WriteItem(buffer)
		buffer.WriteQ(item.GetPrice())
		buffer.WriteQ(int64(item.GetDefaultPrice()) * 2)
	}

	return buffer
}
