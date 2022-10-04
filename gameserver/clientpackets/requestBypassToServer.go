package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/packets"
	"strconv"
	"strings"
)

func BypassToServer(data []byte, client interfaces.ReciverAndSender) {
	var packet = packets.NewReader(data)
	command := packet.ReadString()

	if strings.HasPrefix(command, "admin_create_item") {
		addAdminItem(command, client)
	}
}

func addAdminItem(command string, client interfaces.ReciverAndSender) {
	// 0 - префикс(комманда) 1- Id предмета 2 - количество
	s := strings.Split(command, " ")
	if len(s) != 3 {
		return
	}
	itemId, err := strconv.Atoi(s[1])
	if err != nil {
		return
	}
	count, err := strconv.Atoi(s[2])
	if err != nil {
		return
	}
	item, ok := items.GetItemInfo(itemId)
	if ok {
		client.GetCurrentChar().GetInventory().AddItem2(int32(itemId), count, item.IsStackable())
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
