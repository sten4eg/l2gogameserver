package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
)

//Удаление предмета
func DestroyItem(data []byte, clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.Client)
	if !ok {
		return []byte{}
	}
	var packet = packets.NewReader(data)

	objectId := packet.ReadInt32()
	count := packet.ReadInt32()

	if count == 0 {
		log.Println("Нельзя удалить ноль предметов")
		return []byte{}
	}

	item := client.CurrentChar.ExistItemInInventory(objectId)
	if item == nil {
		log.Println("Не найден предмет")
		return []byte{}
	}

	//Удаляем из инвентаря предмет
	models.RemoveItemCharacter(client.CurrentChar, item, int64(count))
	log.Println("Предмет был удален!")

	buff := packets.Get()
	defer packets.Put(buff)
	pkg := serverpackets.InventoryUpdate(client, item, models.UpdateTypeRemove)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buff.Bytes()
}
