package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
	"log"
)

//Удаление предмета
func DestroyItem(data []byte, client *models.Client) []byte {
	var packet = packets.NewReader(data)

	objectId := packet.ReadInt32()
	count := packet.ReadInt32()

	if count == 0 {
		log.Println("Нельзя удалить ноль предметов")
		return []byte{}
	}

	item, ok := models.CheckIsItemCharacter(client.CurrentChar, objectId)
	if ok == false {
		log.Println("Не найден предмет")
		return []byte{}
	}
	if int32(item.Count) >= count && int32(item.Count) <= count {
		log.Println("Неверное количество предметов для удаления")
		return []byte{}
	}

	//Удаляем из инвентаря предмет
	if models.RemoveItemCharacter(client.CurrentChar, objectId, item.Count) == false {
		log.Println("Удаление не произошло, значит какая-то фигня")
		return []byte{}
	}

	pkgInventoryUpdate := InventoryUpdate(client, client.CurrentChar.ObjectId, models.UpdateTypeRemove)
	client.SSend(pkgInventoryUpdate)

	log.Println("Предмет был удален!")
	return []byte{}
}
