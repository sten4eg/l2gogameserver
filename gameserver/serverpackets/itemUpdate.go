package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ItemUpdate(client interfaces.NewClientCtxInterface, updateType int16, objId int32) []byte {
	buffer := packets.Get()

	items := client.GetAccount().GetCurrentChar().GetInventory().GetItems()
	var item interfaces.MyItemInterface

	for _, v := range items {
		if v.GetObjectId() == objId {
			item = v
		}
	}
	if item.GetObjectId() == 0 {
		return []byte{}
	}

	buffer.WriteSingleByte(0x21)
	buffer.WriteH(1)

	buffer.WriteH(int16(updateType))            // Update type : 01-add, 02-modify, 03-remove
	buffer.WriteD(item.GetObjectId())           //idItemInDB
	buffer.WriteD(int32(item.GetId()))          //getDisplayId idItemsInLineage
	buffer.WriteD(item.GetLocData())            //Location
	buffer.WriteQ(item.GetCount())              //Count
	buffer.WriteH(int16(item.GetItemType2()))   // Item tType 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
	buffer.WriteH(0)                            // Filler (always 0)
	buffer.WriteH(item.IsEquipped())            //  Equipped : 00-No, 01-yes
	buffer.WriteD(int32(item.GetSlotBitType())) // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
	buffer.WriteH(int16(item.GetEnchant()))     // Enchant level (pet level shown in control item)

	buffer.WriteH(0)                     // Pet name exists or not shown in control item
	buffer.WriteD(0)                     // getAugmentationBonus
	buffer.WriteD(int32(item.GetMana())) // mana
	buffer.WriteD(-9999)                 // time

	buffer.WriteH(int16(item.GetAttackElementType()))  //getAttackElementType
	buffer.WriteH(int16(item.GetAttackElementPower())) //getAttackElementPower

	// Аттрибут в броне
	for _, a := range item.GetElementDefAttr() {
		buffer.WriteH(a)
	}

	////// АУГМЕНТАЦИЯ
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	/////////////////////

	return buffer.Bytes()

}
