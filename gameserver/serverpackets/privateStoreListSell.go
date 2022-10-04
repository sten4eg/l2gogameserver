package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func PrivateStoreListSell(character, storeCharacter interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xA1)
	buffer.WriteD(storeCharacter.GetObjectId())
	buffer.WriteD(utils.BoolToInt32(storeCharacter.GetSellList().IsPackaged()))
	buffer.WriteQ(character.GetInventory().GetAdenaCount())
	buffer.WriteD(int32(len(storeCharacter.GetSellList().GetItems())))

	for _, item := range storeCharacter.GetSellList().GetItems() {
		item.WriteItem(buffer)
		buffer.WriteQ(item.GetPrice())
		buffer.WriteQ(int64(item.GetDefaultPrice() * 2))
	}
	return buffer
}
