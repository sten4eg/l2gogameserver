package clientpackets

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/trader"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
	"os"
	"strconv"
	"strings"
)

func BypassToServer(data []byte, client interfaces.NewClientCtxInterface, db *sql.DB) {
	var packet = packets.NewReader(data)
	command := packet.ReadString()
	logger.LogInfo("BypassToServer: %s", command)
	if strings.HasPrefix(command, "admin_create_item") {
		addAdminItem(command, client, db)
	} else if strings.HasPrefix(command, "admin_gmshop") {
		openGmShop(client)
	} else if strings.HasPrefix(command, "admin_move_to") {
		moveAdminTo(command, client, db)
	} else if strings.HasPrefix(command, "admin_help") || strings.HasPrefix(command, "admin_html") || strings.HasPrefix(command, "admin_show_html") {
		openHelp(command, client)
	} else if strings.HasPrefix(command, "admin_buy") {
		buyAdminItem(command, client, db)
	}
}

func buyAdminItem(command string, client interfaces.NewClientCtxInterface, db *sql.DB) {
	args := strings.Fields(command)
	logger.LogInfo("buyAdminItem: args=%v", args)

	shopID, err := strconv.Atoi(args[1])
	if err != nil {
		logger.LogError("Invalid item ID: %v", err)
		return
	}

	shopItemList, ok := trader.GetShopByID(shopID)
	if !ok {
		logger.LogError("Invalid shop ID: %v", shopID)
		return
	}

	pkg := serverpackets.BuyList(*shopItemList, shopID, client.GetAccount().GetCurrentChar())
	client.SendBuf(pkg)

	//Не работает, что-то там неверно, вероятно всё неверно с закрытием окна
	//pkg = serverpackets.CloseBuyList(client.GetAccount().GetCurrentChar().GetInventory().GetAdenaCount(), int32(shopID))
	//client.SendBuf(pkg)
}

func openHelp(command string, client interfaces.NewClientCtxInterface) {
	args := strings.Fields(command) // split by any whitespace
	logger.LogInfo("openHelp: args=%v", args)

	page := args[1]

	data, err := os.ReadFile("./datapack/html/admin/" + page)
	if err != nil {
		logger.LogError("Error reading file: %v", err)
		return
	}
	htmlContent := utils.B2s(data)
	client.SendBuf(serverpackets.NpcHtmlMessage2(0, htmlContent, 0))
}

// Мб потом нужно сделать сохранение в БД координат игрока или удалить db ибо пока не знаю как сохрениня состояний перса реализовано
func moveAdminTo(command string, client interfaces.NewClientCtxInterface, db *sql.DB) {
	args := strings.Fields(command) // split by any whitespace
	logger.LogInfo("moveAdminTo: args=%v", args)

	if len(args) == 4 {
		locX, err := strconv.Atoi(args[1])
		if err != nil {
			logger.LogError("Invalid X coordinate: %v", err)
			return
		}
		locY, err := strconv.Atoi(args[2])
		if err != nil {
			logger.LogError("Invalid Y coordinate: %v", err)
			return
		}
		locZ, err := strconv.Atoi(args[3])
		if err != nil {
			logger.LogError("Invalid Z coordinate: %v", err)
			return
		}

		pkg := serverpackets.TeleportToLocation(client.GetAccount().GetCurrentChar(), locX, locY, locZ, 0)

		if err := client.EncryptAndSend(pkg); err != nil {
			logger.Info.Printf("Ошибка отправки пакета телепортации: %v", err)
		}
		return
	}

	data, err := os.ReadFile("./datapack/html/admin/teleports/WorldAreas.htm")
	if err != nil {
		logger.LogError("Error reading file: %v", err)
		return
	}
	htmlContent := utils.B2s(data)
	client.SendBuf(serverpackets.NpcHtmlMessage2(0, htmlContent, 0))

}

func openGmShop(client interfaces.NewClientCtxInterface) {
	data, err := os.ReadFile("./datapack/html/admin/gmshop.htm")
	if err != nil {
		logger.LogError("Error reading file: %v", err)
		return
	}
	htmlContent := utils.B2s(data)
	client.SendBuf(serverpackets.NpcHtmlMessage2(0, htmlContent, 0))
}

func addAdminItem(command string, client interfaces.NewClientCtxInterface, db *sql.DB) {
	args := strings.Fields(command) // split by any whitespace
	logger.LogInfo("addAdminItem: args=%v", args)

	switch len(args) {
	case 1:
		// Только команда без параметров — показываем HTML
		htmlPath := "./datapack/html/admin/create_item.htm"
		data, err := os.ReadFile(htmlPath)
		if err != nil {
			logger.LogError("Failed to read HTML file '%s': %v", htmlPath, err)
			return
		}
		client.SendBuf(serverpackets.NpcHtmlMessage2(0, utils.B2s(data), 0))
		return

	case 2, 3:
		// Есть ID предмета и, возможно, количество
		itemID, err := strconv.Atoi(args[1])
		if err != nil {
			logger.LogError("Invalid item ID: %v", err)
			return
		}

		count := 1
		if len(args) == 3 {
			count, err = strconv.Atoi(args[2])
			if err != nil {
				logger.LogError("Invalid count value: %v", err)
				return
			}
			if count <= 0 {
				logger.LogError("Count must be positive: %d", count)
				return
			}
		}

		item, found := items.GetItemInfo(itemID)
		if !found {
			logger.LogError("Item not found: ID=%d", itemID)
			return
		}

		char := client.GetAccount().GetCurrentChar()
		if char == nil {
			logger.LogError("No current character")
			return
		}

		char.GetInventory().AddItem2(int32(itemID), count, item.IsStackable(), db)

		pkg := serverpackets.ItemList(client.GetAccount().GetCurrentChar())
		client.EncryptAndSend(pkg)

		logger.LogInfo("Item added: ID=%d, count=%d", itemID, count)
		return

	default:
		logger.LogError("Too many parameters: got %d, expected 1–3", len(args))
		return
	}
}

//TODO всё что ниже написал logan22, может быть что то и понадобится
/*
	Пока заметка, направление как делать.
	Разберем на будущее парсинг bypass
	Все запросы на открытие страницы будут начинатся с _bbspage
	следующие параметры разделены двоеточием:
	[вызов страницы]:[команда]:[информация]:[информация]:[информация]...
	_bbspage:open:/page/index.htm (аналог _bbspage:open:page) - открыть файл
	_bbspage:buffer:combo:3 - наложение комбо баффа с ID 3
	_bbspage:buffer:save - сохранить бафф персонажа
	_bbspage:buffer:get:3 - бафф персонажа (ранее сохраненным баффом) с ID 3
	// Другие аналогия
	_bbspage:gmshop:multisell:1531 - открыть мультиселл 1531
	_bbspage:teleport:id:152 - Телепорт по координатам с ID 152
	_bbspage:teleport:save	- сохранение позиции (xyz) персонажа
	_bbspage:teleport:to:5 - телепорт ранее сохраненную позицию с ID 5
	_bbspage:teleport:remove:5 - удаление сохраненной точки с ID 5
	...
*/
//func BypassToServer(data []byte, client interfaces.ReciverAndSender) {
//	var bypassRequest = packets.NewReader(data).ReadString()
//	bypassInfo := strings.Split(bypassRequest, ":")
//	for i, s := range bypassInfo {
//		logger.Info.Println("#", i, "->", s)
//	}
//	logger.Info.Println(bypassInfo)
//	if bypassInfo[0] == "_bbshome" && bypassRequest == "_bbshome" {
//		//Открытие диалога по умолчанию
//		SendOpenDialogBBS(client, "./datapack/html/community/index.htm")
//	} else if bypassInfo[0] == "_bbspage" {
//		commandname := bypassInfo[1]
//		switch commandname {
//		//Запрос открытия диалога
//		case "open":
//			SendOpenDialogBBS(client, "./datapack/html/community/"+bypassInfo[2])
//
//		//Функции телепортации
//		case "teleport":
//			switch bypassInfo[2] {
//			case "id":
//				teleportID, err := strconv.Atoi(bypassInfo[3])
//				if err != nil {
//					logger.Info.Println(err)
//					return
//				}
//				pkg := community.UserTeleport(client, teleportID)
//				client.EncryptAndSend(pkg)
//			case "save":
//				logger.Info.Println("Сохранение позиции игрока")
//			case "to":
//				logger.Info.Println("Телепорт по сохраненной позиции игрока #", bypassInfo[3])
//			case "remove":
//				logger.Info.Println("Удаление по сохраненной позиции игрока #", bypassInfo[3])
//			}
//
//		case "gmshop":
//			switch bypassInfo[2] {
//			case "multisell": //Open multisell
//				id, err := strconv.Atoi(bypassInfo[3])
//				if err != nil {
//					logger.Info.Println(err)
//					return
//				}
//				logger.Info.Println("Открыть мультиселл с ID", id)
//				multisellList, ok := multisell.Get(client, id)
//				if !ok {
//					logger.Info.Println("Не найден запрашиваемый мультисел#")
//				}
//				pkg := serverpackets.MultiSell(multisellList)
//				client.EncryptAndSend(pkg)
//			}
//
//		}
//
//	}
//}
//
////SendOpenDialogBBS Открытие диалога и отправка клиенту диалога
//func SendOpenDialogBBS(client interfaces.ReciverAndSender, filename string) {
//	logger.Info.Println(filename)
//	htmlDialog, err := htm.Open(filename)
//	if err != nil {
//		logger.Info.Println(err)
//		return
//	}
//	htmlDialog = parseVariableBoard(client, htmlDialog)
//	bufferDialog := packets.Get()
//	defer packets.Put(bufferDialog)
//	bufferDialog1 := packets.Get()
//	defer packets.Put(bufferDialog1)
//	bufferDialog2 := packets.Get()
//	defer packets.Put(bufferDialog2)
//
//	if len(*htmlDialog) < 8180 {
//		bufferDialog.WriteSlice(models.ShowBoard(*htmlDialog, "101"))
//		bufferDialog1.WriteSlice(models.ShowBoard("", "102"))
//		bufferDialog2.WriteSlice(models.ShowBoard("", "103"))
//	} else if len(*htmlDialog) < 8180*2 {
//		bufferDialog.WriteSlice(models.ShowBoard((*htmlDialog)[:8180], "101"))
//		bufferDialog1.WriteSlice(models.ShowBoard((*htmlDialog)[8180:], "102"))
//		bufferDialog2.WriteSlice(models.ShowBoard("", "103"))
//	} else if len(*htmlDialog) < 8180*3 {
//		bufferDialog.WriteSlice(models.ShowBoard((*htmlDialog)[:8180], "101"))
//		bufferDialog1.WriteSlice(models.ShowBoard((*htmlDialog)[8180:8180*2], "102"))
//		bufferDialog2.WriteSlice(models.ShowBoard((*htmlDialog)[8180*2:], "103"))
//	}
//	buffer := packets.Get()
//	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(bufferDialog.Bytes()))
//	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(bufferDialog1.Bytes()))
//	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(bufferDialog2.Bytes()))
//	client.Send(buffer.Bytes())
//
//	packets.Put(buffer)
//}
//
////parseVariableBoard Псевдопеременные из html комьюнити заменяем реальными
//func parseVariableBoard(client interfaces.ReciverAndSender, html *string) *string {
//	r := strings.NewReplacer(
//		"<?player_name?>", client.GetCurrentChar().GetName(),
//		"<?player_class?>", strconv.Itoa(int(client.GetCurrentChar().GetClassId())),
//		"<?cb_time?>", time.Now().Format(time.RFC850),
//	)
//	result := r.Replace(*html)
//	return &result
//}
