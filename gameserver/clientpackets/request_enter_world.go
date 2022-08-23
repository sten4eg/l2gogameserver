package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestEnterWorld(clientI interfaces.ReciverAndSender, data []byte) {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return
	}
	buff := packets.Get()

	pkg := serverpackets.UserInfo(client.GetCurrentChar())
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	pkg2 := serverpackets.ExBrExtraUserInfo(client.CurrentChar)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	//Если персонажа никогда не заходил в игру, выдадим ему какие-то стартовые предметы
	if client.CurrentChar.FirstEnterGame {
		client.CurrentChar.SaveFirstInGamePlayer()
		logger.Info.Println("Выдача предметов новому персонажу: ", client.CurrentChar.CharName)

		client.CurrentChar.Inventory = models.AddItem(models.MyItem{
			ObjectId: idfactory.GetNext(),
			Item: items.Item{
				Id: 6379,
			},
			Count: 1,
		}, client.CurrentChar)

		client.CurrentChar.Inventory = models.AddItem(models.MyItem{
			ObjectId: idfactory.GetNext(),
			Item: items.Item{
				Id: 6382,
			},
			Count: 1,
		}, client.CurrentChar)

		client.CurrentChar.Inventory = models.AddItem(models.MyItem{
			ObjectId: idfactory.GetNext(),
			Item: items.Item{
				Id: 6381,
			},
			Count: 1,
		}, client.CurrentChar)

		client.CurrentChar.Inventory = models.AddItem(models.MyItem{
			ObjectId: idfactory.GetNext(),
			Item: items.Item{
				Id: 6380,
			},
			Count: 1,
		}, client.CurrentChar)

	}

	count := uint8(len(client.CurrentChar.Macros))
	for index, macro := range client.CurrentChar.Macros {
		pkg3 := serverpackets.SendMacroList(client, macro, count, index)
		buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg3))
	}

	pkg4 := serverpackets.ItemList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg4))

	pkg5 := serverpackets.ExQuestItemList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg5))

	pkg6 := serverpackets.GameGuardQuery(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg6))

	pkg7 := serverpackets.ExGetBookMarkInfoPacket(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg7))

	pkg8 := serverpackets.ExStorageMaxCount(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg8))

	pkg9 := serverpackets.ShortCutInit(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg9))

	pkg10 := serverpackets.ExBasicActionList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg10))

	pkg11 := serverpackets.SkillList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg11))

	pkg12 := serverpackets.HennaInfo(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg12))

	pkg13 := serverpackets.QuestList(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg13))

	pkg14 := serverpackets.StaticObject(client)
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg14))

	pkg15 := serverpackets.ShortBuffStatusUpdate(client) //todo test
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg15))

	pkg16 := serverpackets.ActionList(client) //todo test
	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg16))

	client.Send(buff.Bytes())
	packets.Put(buff)
	//NPCdistance := client.CurrentChar.SpawnDistancePoint(5000)
	//logger.Info.Printf("Загружено возле игрока %d npc", len(NPCdistance))
	//for id, locdata := range NPCdistance {
	//	npc, err := models.GetNpcInfo(locdata.NpcId)
	//	if err != nil {
	//		//Вернется ошибка что NPC не найден
	//		//Крайне маловероятно что такое может случиться, но лучше подстаховаться.
	//		logger.Info.Println(err)
	//		continue
	//	}
	//	pkg17 := serverpackets.NpcInfo(client, id, npc, locdata)
	//	buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg17))
	//}

}
