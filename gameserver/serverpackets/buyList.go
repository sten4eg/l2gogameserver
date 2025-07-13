package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trader"
	"l2gogameserver/packets"
)

func BuyList(trader trader.GmShopData, listId int, player interfaces.CharacterI) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xB7)
	buffer.WriteD(0x00)

	buffer.WriteQ(player.GetInventory().GetAdenaCount()) // money
	buffer.WriteD(int32(listId))                         // List ID
	buffer.WriteH(int16(len(trader.Items)))              // Size

	for _, item := range trader.Items {
		buffer.WriteD(item.Item.GetId())
		buffer.WriteD(item.Item.GetId())
		buffer.WriteD(0)
		buffer.WriteQ(-1)
		buffer.WriteH(item.Item.GetItemType2())
		buffer.WriteH(int16(item.Item.GetItemType1()))
		buffer.WriteH(0x00)
		buffer.WriteD(item.Item.GetBodyPart())
		buffer.WriteH(0x00)  // Enchant
		buffer.WriteH(0x00)  // Custom Type
		buffer.WriteD(0x00)  // Augment
		buffer.WriteD(-1)    // Mana
		buffer.WriteD(-9999) // Time
		buffer.WriteH(0x00)  // Element Type
		buffer.WriteH(0x00)  // Element Power
		for i := 0; i < 6; i++ {
			buffer.WriteH(0x00)
		}
		buffer.WriteH(0x00)
		buffer.WriteH(0x00)
		buffer.WriteH(0x00)

		buffer.WriteQ(item.Item.GetPrice())
	}

	return buffer
}
