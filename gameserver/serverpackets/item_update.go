package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func ItemUpdate(client *models.Client, updateType int16, objId int32) {
	client.Buffer.Mu.Lock()
	items := client.CurrentChar.Inventory
	var item models.MyItem

	for _, v := range items {
		if v.ObjId == objId {
			item = v
		}
	}
	if item.ObjId == 0 {
		return
	}

	client.Buffer.WriteSingleByte(0x21)
	client.Buffer.WriteH(1)

	client.Buffer.WriteH(updateType)              // Update type : 01-add, 02-modify, 03-remove
	client.Buffer.WriteD(item.ObjId)              //idItemInDB
	client.Buffer.WriteD(int32(item.Id))          //getDisplayId idItemsInLineage
	client.Buffer.WriteD(item.LocData)            //Location
	client.Buffer.WriteQ(item.Count)              //Count
	client.Buffer.WriteH(int16(item.ItemType))    // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
	client.Buffer.WriteH(0)                       // Filler (always 0)
	client.Buffer.WriteH(item.IsEquipped())       //  Equipped : 00-No, 01-yes
	client.Buffer.WriteD(int32(item.SlotBitType)) // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
	client.Buffer.WriteH(int16(item.Enchant))     // Enchant level (pet level shown in control item)

	client.Buffer.WriteH(0)                // Pet name exists or not shown in control item
	client.Buffer.WriteD(0)                // getAugmentationBonus
	client.Buffer.WriteD(int32(item.Mana)) // mana
	client.Buffer.WriteD(-9999)            // time

	client.Buffer.WriteH(int16(item.GetAttackElement())) //getAttackElementType
	client.Buffer.WriteH(int16(item.AttackAttributeVal)) //getAttackElementPower

	// Аттрибут в броне
	for _, a := range item.AttributeDefend {
		client.Buffer.WriteH(a)
	}

	////// АУГМЕНТАЦИЯ
	client.Buffer.WriteH(0)
	client.Buffer.WriteH(0)
	client.Buffer.WriteH(0)
	/////////////////////
	client.Buffer.Mu.Unlock()

	client.SaveAndCryptDataInBufferToSend(true)

}
