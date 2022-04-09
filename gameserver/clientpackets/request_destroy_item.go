package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
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

	buff := packets.Get()
	defer packets.Put(buff)
	pkg := serverpackets.InventoryUpdate(client, item, models.UpdateTypeRemove)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	//Удаляем из инвентаря предмет
	ok, _ = models.RemoveItemCharacter(client.CurrentChar, objectId, item.Count)
	if !ok {
		log.Println("Удаление не произошло, значит какая-то фигня")
		return []byte{}
	}
	log.Println("Предмет был удален!")

	return buff.Bytes()
}
