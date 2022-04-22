package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func DestroyItem(data []byte, clientI interfaces.ReciverAndSender) {
	client, ok := clientI.(*models.Client)
	if !ok {
		return
	}
	var packet = packets.NewReader(data)

	objectId := packet.ReadInt32()
	count := packet.ReadInt32()

	if count == 0 {
		logger.Info.Println("Нельзя удалить ноль предметов")
		return
	}

	item := client.CurrentChar.ExistItemInInventory(objectId)
	if item == nil {
		logger.Info.Println("Не найден предмет")
		return
	}

	//Удаляем из инвентаря предмет
	models.RemoveItemCharacter(client.CurrentChar, item, int64(count))
	logger.Info.Println("Предмет был удален!")

	pkg := serverpackets.InventoryUpdate(*item, models.UpdateTypeRemove)
	client.EncryptAndSend(pkg)
}
