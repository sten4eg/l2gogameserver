package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ItemList(clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.Client)
	if !ok {
		return []byte{}
	}

	buffer := packets.Get()

	myItems := client.CurrentChar.Inventory.Items

	buffer.WriteSingleByte(0x11)
	buffer.WriteH(0)                   // 1 - открывает окно инвентаря
	buffer.WriteH(int16(len(myItems))) // количество всех предметов в инвентаре и на персонаже

	for i := range myItems {
		e := &myItems[i]
		buffer.WriteD(e.ObjId)              // уникальный object_id из бд
		buffer.WriteD(int32(e.Id))          // id предмета в клиенте
		buffer.WriteD(e.LocData)            // номер ячейки в инвентаре или на персонаже, где находится предмет
		buffer.WriteQ(e.Count)              // количество предмета
		buffer.WriteH(int16(e.ItemType))    // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		buffer.WriteH(0)                    // Постоянно 0
		buffer.WriteH(e.IsEquipped())       //  0 - в инвентаре, 1 - одето на персонаже
		buffer.WriteD(int32(e.SlotBitType)) // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		buffer.WriteH(int16(e.Enchant))     // Enchant level (pet level shown in control item)

		buffer.WriteH(0)             // Pet name exists or not shown in control item
		buffer.WriteD(0)             // getAugmentationBonus
		buffer.WriteD(int32(e.Mana)) // Mana(оставшееся время) для Шадоу оружия
		buffer.WriteD(0)             // time   TemporalLifeTime

		buffer.WriteH(int16(e.GetAttackElement())) // Тип Аттрибута атаки
		buffer.WriteH(int16(e.AttackAttributeVal)) // Значение Аттрибута атаки

		// Аттрибут в броне
		for _, a := range e.AttributeDefend {
			buffer.WriteH(a)
		}

		////// АУГМЕНТАЦИЯ
		if e.ObjId == 16 {
			buffer.WriteH(30)
			buffer.WriteH(10)
			buffer.WriteH(20)
		} else {
			buffer.WriteH(0)
			buffer.WriteH(0)
			buffer.WriteH(0)
		}

		/////////////////////

	}

	buffer.WriteH(0) //writeInventoryBlock

	defer packets.Put(buffer)
	return buffer.Bytes()
}
