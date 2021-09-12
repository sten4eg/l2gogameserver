package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func ItemList(client *models.Client) {

	myItems := client.CurrentChar.Inventory

	client.Buffer.Mu.Lock()
	client.Buffer.WriteSingleByte(0x11)
	client.Buffer.WriteH(0)                   // 1 - открывает окно инвентаря
	client.Buffer.WriteH(int16(len(myItems))) // количество всех предметов в инвентаре и на персонаже

	for _, e := range myItems {
		client.Buffer.WriteD(e.ObjId)              // уникальный object_id из бд
		client.Buffer.WriteD(int32(e.Id))          // id предмета в клиенте
		client.Buffer.WriteD(e.LocData)            // номер ячейки в инвентаре или на персонаже, где находится предмет
		client.Buffer.WriteQ(e.Count)              // количество предмета
		client.Buffer.WriteH(int16(e.ItemType))    // Item Type 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		client.Buffer.WriteH(0)                    // Постоянно 0
		client.Buffer.WriteH(e.IsEquipped())       //  0 - в инвентаре, 1 - одето на персонаже
		client.Buffer.WriteD(int32(e.SlotBitType)) // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		client.Buffer.WriteH(int16(e.Enchant))     // Enchant level (pet level shown in control item)

		client.Buffer.WriteH(0)             // Pet name exists or not shown in control item
		client.Buffer.WriteD(0)             // getAugmentationBonus
		client.Buffer.WriteD(int32(e.Mana)) // Mana(оставшееся время) для Шадоу оружия
		client.Buffer.WriteD(0)             // time   TemporalLifeTime

		client.Buffer.WriteH(int16(e.GetAttackElement())) // Тип Аттрибута атаки
		client.Buffer.WriteH(int16(e.AttackAttributeVal)) // Значение Аттрибута атаки

		// Аттрибут в броне
		for _, a := range e.AttributeDefend {
			client.Buffer.WriteH(a)
		}

		////// АУГМЕНТАЦИЯ
		if e.ObjId == 16 {
			client.Buffer.WriteH(30)
			client.Buffer.WriteH(10)
			client.Buffer.WriteH(20)
		} else {
			client.Buffer.WriteH(0)
			client.Buffer.WriteH(0)
			client.Buffer.WriteH(0)
		}

		/////////////////////

	}

	client.Buffer.WriteH(0) //writeInventoryBlock
	client.Buffer.Mu.Unlock()
	client.SaveAndCryptDataInBufferToSend(true)
}
