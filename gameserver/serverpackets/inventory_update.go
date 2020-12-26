package serverpackets

import (
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/packets"
)

func NewInventoryUpdate(items []items.Item) []byte {
	buffer := new(packets.Buffer)

	buffer.WriteH(0)
	buffer.WriteSingleByte(0x21)

	buffer.WriteH(int16(len(items)))
	for _, e := range items {
		buffer.WriteH(2)
		buffer.WriteD(e.ObjId)
		buffer.WriteD(e.Id)
		buffer.WriteD(e.LocData)
		buffer.WriteQ(e.Count)
		buffer.WriteH(e.ItemType)
		buffer.WriteH(0) // Filler (always 0)
		buffer.WriteH(e.IsEquipped())
		buffer.WriteD(e.Bodypart)
		buffer.WriteH(e.Enchant)
		//buffer.WriteD(21)    //idItemInDB
		//buffer.WriteD(1147)  //getDisplayId idItemsInLineage
		//buffer.WriteD(0)     //Location
		//buffer.WriteQ(1)     //Count
		//buffer.WriteH(00)    // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		//buffer.WriteH(0)     // Filler (always 0)
		//buffer.WriteH(0)     //  Equipped : 00-No, 01-yes
		//buffer.WriteD(0)     // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		//buffer.WriteH(0)     // Enchant level (pet level shown in control item)
		buffer.WriteH(0)     // Pet name exists or not shown in control item
		buffer.WriteD(0)     // getAugmentationBonus
		buffer.WriteD(-1)    // mana
		buffer.WriteD(-9999) // time

		buffer.WriteH(-2) //getAttackElementType
		buffer.WriteH(0)  //getAttackElementPower

		////// ELEM DEF /////
		buffer.WriteH(0)
		buffer.WriteH(0)
		buffer.WriteH(0)
		buffer.WriteH(0)
		buffer.WriteH(0)
		buffer.WriteH(0)
		////////////////////

		////// ENCHANT OPTION
		buffer.WriteH(0)
		buffer.WriteH(0)
		buffer.WriteH(0)
		/////////////////////

	}

	return buffer.Bytes()
}
