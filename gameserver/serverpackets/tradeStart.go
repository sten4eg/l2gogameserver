package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func TradeStart(character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	items := character.GetInventory().GetItems()

	buffer.WriteSingleByte(0x14)
	buffer.WriteD(character.GetActiveTradeList().GetPartner().GetObjectId())
	buffer.WriteH(int16(len(items)))
	for _, item := range items {
		buffer.WriteD(item.GetObjectId())         // ObjectId
		buffer.WriteD(int32(item.GetId()))        // ItemId
		buffer.WriteD(item.GetLocData())          // T1
		buffer.WriteQ(item.GetCount())            // Quantity
		buffer.WriteH(int16(item.GetItemType2())) // Item tType 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		buffer.WriteH(0)                          // Filler (always 0)
		//00 шмот который надет, будет виден в торге
		buffer.WriteH(item.IsEquipped())         // Equipped : 00-No, 01-yes
		buffer.WriteD(int32(item.GetBodyPart())) // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		buffer.WriteH(int16(item.GetEnchant()))  // Enchant level (pet level shown in control item)
		buffer.WriteH(0)                         // Pet name exists or not shown in control item
		buffer.WriteD(0)                         //getAugmentationBonus
		buffer.WriteD(int32(item.GetMana()))
		buffer.WriteD(int32(item.GetTime()))

		buffer.WriteH(-2)
		buffer.WriteH(0)
		for i := 0; i < 6; i++ {
			buffer.WriteH(0)
		}
		buffer.WriteH(0)
		buffer.WriteH(0)
		buffer.WriteH(0)
	}
	return buffer.Bytes()
}
