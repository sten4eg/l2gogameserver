package clientpackets

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/gameserver/world"
	"l2gogameserver/packets"
	"strconv"
)

func Action(data []byte, client interfaces.NewClientCtxInterface,
	f func(client interfaces.NewClientCtxInterface,
		l world.BackwardToLocation),
	db *sql.DB,
) *world.BackwardToLocation {

	reAppeal := false // повторное обращение к объекту
	var packet = packets.NewReader(data)
	objectId := packet.ReadInt32() //Target
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	actionId := packet.ReadSingleByte() // Action identifier : 0-Simple click, 1-Shift click

	if objectId == client.GetAccount().GetCurrentChar().GetTarget() {
		reAppeal = true
	}
	//else {
	//	client.CurrentChar.Target = objectId
	//}

	object := getTargetByObjectId(objectId, client.GetAccount().GetCurrentChar().GetCurrentRegion())

	switch target := object.(type) {
	case interfaces.MyItemInterface:
		itemAction(client, target, actionId, db)
	case interfaces.CharacterI:
		characterAction(client, target, actionId)
	case interfaces.Npcer:
		npcAction(client, target, actionId)
	default:
		logger.Info.Println("Wrong object type")
	}

	_, _ = actionId, reAppeal

	f(client, *MoveToLocation(client, originX, originY, originZ))
	return MoveToLocation(client, originX, originY, originZ)

	/*
		npc, npcx, npcy, npcz, err := models.GetNpcObject(objectId)
		if err != nil {
			logger.Info.Println(err)
		}

		//Прост тест вызова HTML при клике
		if actionId == 1 {
			NpcHtmlMessage := NpcHtmlMessage(client, npc.NpcId)
			client.Send(NpcHtmlMessage)
		}
		//Если повторный клик по нпц
		if reAppeal {
			//npc, npcx, npcy, npcz, err := models.GetNpcObject(objectId)
			//if err != nil {
			//	logger.Info.Println(err)
			//}
			x, y, z := client.CurrentChar.GetXYZ()
			distance := models.CalculateDistance(npcx, npcy, npcz, x, y, z, false, false)
			_, _ = distance, npc

			//подбегаем
			if distance <= 60 {
				logger.Info.Println("Расстояние до NPC подходящее")
				if models.GetDialogNPC(npc.tType) == 0 {
					//НПЦ для разговора, открываем диалог
					//Пускай макс. дистанция разговора будет 60 поинтов
					//Пока откроем ID нпц
					NpcHtmlMessage := NpcHtmlMessage(client, npc.NpcId)
					client.Send(NpcHtmlMessage)
				} else {
					//бьем нпц
					client.Send(Attack(data, client))
				}
			} else {
				logger.Info.Println("Расстояние до NPC слишком больше, необходимо подбежать")
				return MoveToLocation(client, npcx, npcy, npcz)

			}

		}

	*/
	return nil
}

// TODO мб переместить в медоты региона и чекать соседние регионы
func getTargetByObjectId(objId int32, region interfaces.WorldRegioner) any {
	var target any
	var ok bool

	for _, v := range region.GetNeighbors() {
		target, ok = v.GetItem(objId)
		if ok {
			return target
		}
		target, ok = v.GetChar(objId)
		if ok {
			return target
		}
		target, ok = v.GetNpc(objId)
		if ok {
			return target
		}
	}

	return nil
}

func itemAction(client interfaces.NewClientCtxInterface, item interfaces.MyItemInterface, actionId byte, db *sql.DB) {
	switch actionId {
	case 0:
		doActionOnItem(client, item, db)
	case 1:
		doActionShiftOnItem(client, item)
	default:
		logger.Info.Panicln("Wrong actionId")
	}
}

func doActionOnItem(client interfaces.NewClientCtxInterface, item interfaces.MyItemInterface, db *sql.DB) {
	buffer := serverpackets.GetItem(item, client.GetAccount().GetCurrentChar().GetObjectId())
	broadcast.BroadCastBufferToAroundPlayers(client, buffer)

	buffer2 := serverpackets.DeleteObject(item.GetObjectId())
	broadcast.BroadCastBufferToAroundPlayers(client, buffer2)

	updateItem := client.GetAccount().GetCurrentChar().GetInventory().AddItem2(item.GetId(), int(item.GetCount()), item.IsStackable(), db)
	client.GetAccount().GetCurrentChar().GetCurrentRegion().DeleteVisibleItem(item)

	items := []interfaces.MyItemInterface{updateItem}
	msg := serverpackets.InventoryUpdate(items)
	client.EncryptAndSend(msg)
}

func doActionShiftOnItem(client interfaces.NewClientCtxInterface, item interfaces.MyItemInterface) {
	//TODO Генерировать html нормальный образом
	//html := "<html><body><center><font color=\"LEVEL\">Item Info</font></center><br></body></html>"

	itemObjId := strconv.FormatInt(int64(item.GetObjectId()), 10)
	itemId := strconv.FormatInt(int64(item.GetId()), 10)
	x, y, z := item.GetCoordinate()
	itemLocation := strconv.FormatInt(int64(x), 10) + " " + strconv.FormatInt(int64(y), 10) + " " + strconv.FormatInt(int64(z), 10)
	itemClass := "*models.MyItem"

	html := "<html><body><center><font color=\"LEVEL\">Item Info</font></center><br><table border=0>" + "<tr><td>Object ID: </td><td>" + itemObjId + "</td></tr><tr><td>Item ID: </td><td>" + itemId + "</td></tr><tr><td>Owner ID: </td><td>" + "0" + "</td></tr><tr><td>Location: </td><td>" + itemLocation + "</td></tr><tr><td><br></td></tr><tr><td>Class: </td><td>" + itemClass + "</td></tr></table></body></html>"

	pkg := serverpackets.NpcHtmlMessage2(0, html, 0)
	client.SendBuf(pkg)

}

func characterAction(client interfaces.NewClientCtxInterface, char interfaces.CharacterI, actionId byte) {
	switch actionId {
	case 0:
		doActionOnCharacter(client, char)
	case 1:
		//TODO Доделать окно с информации о персонаже
		break
	default:
		logger.Info.Panicln("Wrong actionId")
	}
}

func doActionOnCharacter(client interfaces.NewClientCtxInterface, targetChar interfaces.CharacterI) {
	if client.GetAccount().GetCurrentChar().GetTarget() != targetChar.GetObjectId() {
		client.GetAccount().GetCurrentChar().SetTarget(targetChar.GetObjectId())
		x, y, z := targetChar.GetXYZ()
		pkg := serverpackets.TargetSelected(client.GetAccount().GetCurrentChar().GetObjectId(), targetChar.GetObjectId(), x, y, z)
		client.SendBuf(pkg)
	} else {
		if targetChar.GetPrivateStoreType() == privateStoreType.SELL || targetChar.GetPrivateStoreType() == privateStoreType.PACKAGE_SELL {
			pkg := serverpackets.PrivateStoreListSell(client.GetAccount().GetCurrentChar(), targetChar)
			client.SendBuf(pkg)
		} else if targetChar.GetPrivateStoreType() == privateStoreType.BUY {
			pkg := serverpackets.PrivateStoreListBuy(client.GetAccount().GetCurrentChar(), targetChar)
			client.SendBuf(pkg)
		}
	}
}

func npcAction(client interfaces.NewClientCtxInterface, npc interfaces.Npcer, actionId byte) {
	switch actionId {
	case 0:
		doActionOnNpc(client, npc, true)
	case 1:
		break
	default:
		logger.Info.Panicln("Wrong actionId")
	}
}

func doActionOnNpc(client interfaces.NewClientCtxInterface, npc interfaces.Npcer, interact bool) {
	if !npc.IsTargetable() {
		return
	}

	//client.CurrentChar.SetLastFolkNPC(npc)

	if npc.GetObjectId() != client.GetAccount().GetCurrentChar().GetTarget() {
		maxHp := npc.GetMaxHp()
		attributes := []serverpackets.Attributes{{Id: serverpackets.MaxHp, Value: maxHp}, {Id: serverpackets.CurHp, Value: maxHp}} // TODO поменять на текущее хп
		client.GetAccount().GetCurrentChar().SetTarget(npc.GetObjectId())
		//TODO проверка isAutoAttackable

		client.SendBuf(serverpackets.MyTargetSelected(npc.GetObjectId()))
		client.SendBuf(serverpackets.StatusUpdate(npc.GetObjectId(), attributes))

	} else if interact {
		//TODO взаимодействие с нпц
	}
}
