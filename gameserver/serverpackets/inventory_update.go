package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func InventoryUpdate(item *models.MyItem, updateType int16) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x21)
	buffer.WriteH(1)

	buffer.WriteH(updateType)              // Update type : 01-add, 02-modify, 03-remove
	buffer.WriteD(item.ObjId)              //idItemInDB
	buffer.WriteD(int32(item.Id))          //getDisplayId idItemsInLineage
	buffer.WriteD(item.LocData)            //Location
	buffer.WriteQ(item.Count)              //Count
	buffer.WriteH(int16(item.ItemType))    // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
	buffer.WriteH(0)                       // Filler (always 0)
	buffer.WriteH(item.IsEquipped())       //  Equipped : 00-No, 01-yes
	buffer.WriteD(int32(item.SlotBitType)) // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
	buffer.WriteH(int16(item.Enchant))     // Enchant level (pet level shown in control item)

	buffer.WriteH(0)                // Pet name exists or not shown in control item
	buffer.WriteD(0)                // getAugmentationBonus
	buffer.WriteD(int32(item.Mana)) // mana
	buffer.WriteD(-9999)            // time

	buffer.WriteH(int16(item.GetAttackElement())) //getAttackElementType
	buffer.WriteH(int16(item.AttackAttributeVal)) //getAttackElementPower

	// Аттрибут в броне
	for _, a := range item.AttributeDefend {
		buffer.WriteH(a)
	}

	////// АУГМЕНТАЦИЯ
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	/////////////////////

	return buffer.Bytes()

}
