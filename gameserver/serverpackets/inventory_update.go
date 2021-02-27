package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
)

func NewInventoryUpdate(client *models.Client, items []items.Item) {

	client.Buffer.WriteSingleByte(0x21)
	client.Buffer.WriteH(int16(len(items)))
	for _, e := range items {
		client.Buffer.WriteH(2)
		client.Buffer.WriteD(e.ObjId)
		client.Buffer.WriteD(e.Id)
		client.Buffer.WriteD(e.LocData)
		client.Buffer.WriteQ(e.Count)
		client.Buffer.WriteH(e.ItemType)
		client.Buffer.WriteH(0) // Filler (always 0)
		client.Buffer.WriteH(e.IsEquipped())
		client.Buffer.WriteD(e.Bodypart)
		client.Buffer.WriteH(e.Enchant)
		//client.Buffer.WriteD(21)    //idItemInDB
		//client.Buffer.WriteD(1147)  //getDisplayId idItemsInLineage
		//client.Buffer.WriteD(0)     //Location
		//client.Buffer.WriteQ(1)     //Count
		//client.Buffer.WriteH(00)    // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		//client.Buffer.WriteH(0)     // Filler (always 0)
		//client.Buffer.WriteH(0)     //  Equipped : 00-No, 01-yes
		//client.Buffer.WriteD(0)     // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		//client.Buffer.WriteH(0)     // Enchant level (pet level shown in control item)
		client.Buffer.WriteH(0)     // Pet name exists or not shown in control item
		client.Buffer.WriteD(0)     // getAugmentationBonus
		client.Buffer.WriteD(-1)    // mana
		client.Buffer.WriteD(-9999) // time

		client.Buffer.WriteH(-2) //getAttackElementType
		client.Buffer.WriteH(0)  //getAttackElementPower

		////// ELEM DEF /////
		client.Buffer.WriteH(0)
		client.Buffer.WriteH(0)
		client.Buffer.WriteH(0)
		client.Buffer.WriteH(0)
		client.Buffer.WriteH(0)
		client.Buffer.WriteH(0)
		////////////////////

		////// ENCHANT OPTION
		client.Buffer.WriteH(0)
		client.Buffer.WriteH(0)
		client.Buffer.WriteH(0)
		/////////////////////

	}
	client.SaveAndCryptDataInBufferToSend(true)

}
