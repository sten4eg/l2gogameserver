package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

/**
Данная функция подлежит рефакторингу, проверки и доработки безопастности
- После структурирования инвентаря, нужно будет доделать удаление предмета из инвентаря
*/
//Выбрасывание предмета из инвентаря
func DropItem(client *models.Client, objectId int32, count int64, x, y, z int32) []byte {
	if count == 0 {
		return []byte{}
	}

	itemId := 0                 // ID предмета
	var inCount int64 = 0       //Кол-во
	var remainder int64 = 0     //Кол-во которое осталось после выброса у персонажа
	var isStackable byte = 0x01 //0 стыкуется, 1 не стыкуется
	for _, e := range client.CurrentChar.Inventory {
		if e.ObjId == objectId {
			itemId = e.Id
			inCount = e.Count
			remainder = e.Count - count
			if e.ConsumeType == "consume_type_stackable" {
				isStackable = 0x00
			}
			break
		}
	}
	_ = remainder //Нужна проверка на остаточное кол-во предметов

	//Если id 0 тогда предмет не был найден в инвентаре, вероятно жулик!!!
	if itemId == 0 {
		return []byte{}
	}

	if inCount >= count {
		buffer := packets.Get()
		defer packets.Put(buffer)

		buffer.WriteSingleByte(0x16)
		buffer.WriteD(client.CurrentChar.CharId)
		buffer.WriteD(objectId)
		buffer.WriteD(int32(itemId))
		buffer.WriteD(x)
		buffer.WriteD(y)
		buffer.WriteD(z)
		buffer.WriteD(int32(isStackable))
		buffer.WriteQ(count)
		buffer.WriteD(0x01)

		return buffer.Bytes()
	}

	return []byte{}
}
