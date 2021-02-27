package serverpackets

import (
	"github.com/jackc/pgx"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
)

func NewItemList(client *models.Client, conn *pgx.Conn) {

	client.Buffer.WriteSingleByte(0x11)
	client.Buffer.WriteH(0)

	myItems := items.GetMyItems(client.CurrentChar.CharId, conn)
	client.Buffer.WriteH(int16(len(myItems)))

	for _, e := range myItems {
		client.Buffer.WriteD(e.ObjId)        //idItemInDB
		client.Buffer.WriteD(e.Id)           //getDisplayId idItemsInLineage
		client.Buffer.WriteD(e.LocData)      //Location
		client.Buffer.WriteQ(e.Count)        //Count
		client.Buffer.WriteH(e.ItemType)     // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		client.Buffer.WriteH(0)              // Filler (always 0)
		client.Buffer.WriteH(e.IsEquipped()) //  Equipped : 00-No, 01-yes
		client.Buffer.WriteD(e.Bodypart)     // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		client.Buffer.WriteH(e.Enchant)      // Enchant level (pet level shown in control item)

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

	client.Buffer.WriteH(0) //writeInventoryBlock

	client.SaveAndCryptDataInBufferToSend(true)
}
