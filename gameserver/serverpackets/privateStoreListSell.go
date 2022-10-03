package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func PrivateStoreListSell(player, storePlayer interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xA1)
	buffer.WriteD(storePlayer.GetObjectId())
	buffer.WriteD(utils.BoolToInt32(storePlayer.GetSellList().IsPackaged()))
	buffer.WriteQ(player.GetInventory().GetAdenaCount())
	buffer.WriteD(int32(len(storePlayer.GetSellList().GetItems())))

	for _, item := range storePlayer.GetSellList().GetItems() {
		item.WriteItem(buffer)
		buffer.WriteQ(item.GetPrice())
		buffer.WriteQ(int64(item.GetDefaultPrice() * 2))
	}
	return buffer
}
