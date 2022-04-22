package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func TradeSendRequest(target interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x70)
	buffer.WriteD(target.GetObjectId())

	return buffer.Bytes()
}

func TradeStart(player *models.Character) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x14)
	buffer.WriteD(player.ObjectId)
	buffer.WriteH(int16(len(player.Inventory.Items)))
	for _, item := range player.Inventory.Items {
		buffer.WriteD(item.ObjId)           // ObjectId
		buffer.WriteD(int32(item.Id))       // ItemId
		buffer.WriteD(item.LocData)         // T1
		buffer.WriteQ(item.Count)           // Quantity
		buffer.WriteH(int16(item.ItemType)) // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		buffer.WriteH(0)                    // Filler (always 0)
		//00 шмот который надет, будет виден в торге
		buffer.WriteH(item.IsEquipped())       // Equipped : 00-No, 01-yes
		buffer.WriteD(int32(item.SlotBitType)) // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		buffer.WriteH(int16(item.Enchant))     // Enchant level (pet level shown in control item)
		buffer.WriteH(0)                       // Pet name exists or not shown in control item
		buffer.WriteD(0)                       //getAugmentationBonus
		buffer.WriteD(int32(item.Mana))
		buffer.WriteD(int32(item.Time))

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
