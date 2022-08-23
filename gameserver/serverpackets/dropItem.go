package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

/**
Данная функция подлежит рефакторингу, проверки и доработки безопастности
- После структурирования инвентаря, нужно будет доделать удаление предмета из инвентаря
*/
//Выбрасывание предмета из инвентаря
func DropItem(clientI interfaces.ReciverAndSender, objectId int32, count int64, x, y, z int32) []byte {
	if count == 0 {
		return []byte{}
	}
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return []byte{}
	}

	itemId := 0             // ID предмета
	var inCount int64 = 0   //Кол-во
	var remainder int64 = 0 //Кол-во которое осталось после выброса у персонажа
	var isStackable int32   //0 стыкуется, 1 не стыкуется
	for _, e := range client.CurrentChar.Inventory.Items {
		if e.ObjectId == objectId {
			itemId = e.Id
			inCount = e.Count
			remainder = e.Count - count
			isStackable = int32(e.ConsumeType)
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
		buffer.WriteD(client.CurrentChar.ObjectId)
		buffer.WriteD(objectId)
		buffer.WriteD(int32(itemId))
		buffer.WriteD(x)
		buffer.WriteD(y)
		buffer.WriteD(z)
		buffer.WriteD(isStackable)
		buffer.WriteQ(count)
		buffer.WriteD(0x01)

		return buffer.Bytes()
	}

	return []byte{}
}
