package serverpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func TradeStart(playerI interfaces.CharacterI) []byte {
	player, ok := playerI.(*models.Character)
	if !ok {
		logger.Error.Panicln("TradeStart playerI is not Character")
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x14)
	buffer.WriteD(player.ActiveTradeList.GetPartner().GetObjectId())
	buffer.WriteH(int16(len(player.Inventory.Items)))
	for _, item := range player.Inventory.Items {
		buffer.WriteD(item.ObjectId)         // ObjectId
		buffer.WriteD(int32(item.Id))        // ItemId
		buffer.WriteD(item.LocData)          // T1
		buffer.WriteQ(item.Count)            // Quantity
		buffer.WriteH(int16(item.ItemType2)) // Item tType 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		buffer.WriteH(0)                     // Filler (always 0)
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
