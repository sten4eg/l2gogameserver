package clientpackets

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func DestroyItem(data []byte, clientI interfaces.ReciverAndSender, db *sql.DB) {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return
	}
	char := client.GetCurrentChar()
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
	//models.RemoveItemCharacter(client.CurrentChar, item, int64(count))
	char.GetInventory().DestroyItem(item, int(count), db)
	logger.Info.Println("Предмет был удален!")

	//TODO сделать нормально удаление
	items := []interfaces.MyItemInterface{item}
	pkg := serverpackets.InventoryUpdate(items)
	client.EncryptAndSend(pkg)
}
