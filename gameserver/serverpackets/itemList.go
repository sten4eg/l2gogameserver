package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ItemList(character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	myItems := character.GetInventory().GetItems()

	buffer.WriteSingleByte(0x11)
	buffer.WriteH(0)                   // 1 - открывает окно инвентаря
	buffer.WriteH(int16(len(myItems))) // количество всех предметов в инвентаре и на персонаже

	for _, item := range myItems {
		buffer.WriteD(item.GetObjectId())  // уникальный object_id из бд
		buffer.WriteD(item.GetId())        // id предмета в клиенте
		buffer.WriteD(item.GetLocData())   // номер ячейки в инвентаре или на персонаже, где находится предмет
		buffer.WriteQ(item.GetCount())     // количество предмета
		buffer.WriteH(item.GetItemType2()) // Item tType 2 : 00-weapon, 01-shield/armor, 02-ring/earring/necklace, 03-questitem, 04-adena, 05-item
		buffer.WriteH(0)                   // Постоянно 0
		buffer.WriteH(item.IsEquipped())   //  0 - в инвентаре, 1 - одето на персонаже
		buffer.WriteD(item.GetBodyPart())  // Slot : 0006-lr.ear, 0008-neck, 0030-lr.finger, 0040-head, 0100-l.hand, 0200-gloves, 0400-chest, 0800-pants, 1000-feet, 4000-r.hand, 8000-r.hand
		buffer.WriteH(item.GetEnchant())   // Enchant level (pet level shown in control item)

		buffer.WriteH(0)              // Pet name exists or not shown in control item
		buffer.WriteD(0)              // getAugmentationBonus
		buffer.WriteD(item.GetMana()) // Mana(оставшееся время) для Шадоу оружия
		buffer.WriteD(0)              // time   TemporalLifeTime

		buffer.WriteH(int16(item.GetAttackElementType())) // Тип Аттрибута атаки
		buffer.WriteH(item.GetAttackElementPower())       // Значение Аттрибута атаки

		// Аттрибут в броне
		for _, a := range item.GetElementDefAttr() {
			buffer.WriteH(a)
		}

		////// АУГМЕНТАЦИЯ
		if item.GetObjectId() == 16 {
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
