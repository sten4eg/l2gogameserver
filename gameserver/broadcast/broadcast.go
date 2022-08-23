package broadcast

import (
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/chat"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/utils"
)

// ToAroundPlayerInRadius отправляет всем персонажам в радиусе radius
// информацию из пакета pkg
func ToAroundPlayerInRadius(my interfaces.ReciverAndSender, pkg *utils.PacketByte, radius int32) {
	charsIds := models.GetAroundPlayersInRadius(my.GetCurrentChar(), radius)
	for i := range charsIds {
		charsIds[i].EncryptAndSend(pkg.GetData())
	}
}

func BroadCastToAroundPlayers(my interfaces.ReciverAndSender, pkg *utils.PacketByte) {
	charsIds := models.GetAroundPlayer(my.GetCurrentChar())
	for i := range charsIds {
		charsIds[i].EncryptAndSend(pkg.GetData())
	}
}

// BroadCastUserInfoInRadius отправляет всем персонажам в радиусе radius
// информацию о персонаже, Самому персонажу отправляет полный UserInfo
func BroadCastUserInfoInRadius(me interfaces.ReciverAndSender, radius int32) {
	ui := serverpackets.UserInfo(me.GetCurrentChar())
	me.EncryptAndSend(ui)

	charsIds := models.GetAroundPlayersInRadius(me.GetCurrentChar(), radius)
	if len(charsIds) == 0 {
		return
	}

	ci := utils.GetPacketByte()
	defer ci.Release()

	ci.SetData(serverpackets.CharInfo(me.GetCurrentChar()))

	exUi := utils.GetPacketByte()
	defer exUi.Release()
	exUi.SetData(serverpackets.ExBrExtraUserInfo(me.GetCurrentChar()))

	//g.OnlineCharacters.Mu.Lock()
	for i := range charsIds {
		gameserver.OnlineCharacters.Char[charsIds[i].ObjectId].Conn.EncryptAndSend(ci.GetData())
		gameserver.OnlineCharacters.Char[charsIds[i].ObjectId].Conn.EncryptAndSend(exUi.GetData())
	}
	//g.OnlineCharacters.Mu.Unlock()
}

func BroadCastChat(me interfaces.ReciverAndSender, say models.Say) {
	pb := utils.GetPacketByte()
	defer pb.Release()

	switch say.Type {
	case chat.All:
		cs := serverpackets.CreatureSay(&say, me.GetCurrentChar())
		pb.SetData(cs)
		me.EncryptAndSend(pb.GetData())
		ToAroundPlayerInRadius(me, pb, chat.AllChatRange)
	case chat.Tell:
		cs := serverpackets.CreatureSay(&say, me.GetCurrentChar())
		pb.SetData(cs)
		ok := BroadCastToCharacterByName(pb, say.To)
		if ok {
			me.EncryptAndSend(pb.GetData())
		} else {
			// systemMSG что не найден перс
		}
	case chat.Shout:
		cs := serverpackets.CreatureSay(&say, me.GetCurrentChar())
		pb.SetData(cs)
		me.EncryptAndSend(pb.GetData())

		ToAroundPlayerInRadius(me, pb, chat.ShoutChatRange)
		//todo что за SpecialCommand ?
		//case chat.SpecialCommand:
		//	if me.CurrentChar.Target == 0 {
		//		return
		//	}
		//	qwe := g.OnlineCharacters.Char[me.CurrentChar.Target]
		//	q := models.CalculateDistance(qwe.Coordinates.X, qwe.Coordinates.Y, qwe.Coordinates.Z, me.CurrentChar.Coordinates.X, me.CurrentChar.Coordinates.Y, me.CurrentChar.Coordinates.Z, false, false)
		//	say.Text = fmt.Sprintf("%f", q)
		//	say.tType = chat.All
		//
		//	cs := serverpackets.CreatureSay(&say, me.CurrentChar)
		//	pb.SetData(cs)
		//	me.Send(me.CryptAndReturnPackageReadyToShip(pb.GetData()))
		//	ToAroundPlayerInRadius(me, pb, chat.AllChatRange)
	}
}

// BroadCastToCharacterByName отправляет pkg персонажу с ником to
// true если отправлен, false если персонаж не найден
func BroadCastToCharacterByName(pkg *utils.PacketByte, to string) bool {
	gameserver.OnlineCharacters.Mu.Lock()
	defer gameserver.OnlineCharacters.Mu.Unlock()

	conn := gameserver.GetNetConnByCharacterName(to)
	if conn != nil {
		conn.EncryptAndSend(pkg.GetData())
		return true
	}

	return false
}

// SendCharInfoAboutCharactersInRadius отправляет me CharInfo персонажей
// в радиусе radius
func SendCharInfoAboutCharactersInRadius(me interfaces.ReciverAndSender, radius int32) {
	charsIds := models.GetAroundPlayersInRadius(me.GetCurrentChar(), radius)
	for i := range charsIds {
		me.EncryptAndSend(serverpackets.CharInfo(charsIds[i]))
	}
}

// SendCharInfoAboutCharactersAround отправляет me CharInfo персонажей
func SendCharInfoAboutCharactersAround(me *models.ClientCtx) {
	charsIds := models.GetAroundPlayer(me.CurrentChar)
	for i := range charsIds {
		me.EncryptAndSend(serverpackets.CharInfo(charsIds[i]))
	}
}

func Checkaem(client interfaces.ReciverAndSender, l models.BackwardToLocation) {
	ut := utils.GetPacketByte()
	ut.SetData(serverpackets.MoveToLocation(&l, client))

	client.EncryptAndSend(ut.GetData())
	BroadCastToAroundPlayers(client, ut)
}

func GetCharacterByObjectId(id int32) interfaces.CharacterI {
	return gameserver.GetNetConnByCharObjectId(id)
}

//func (g *GameServer) Tick() {
//	for {
//		g.clients.Range(func(k, v interface{}) bool {
//			client := v.(*models.ClientCtx)
//			if client.CurrentChar.Coordinates == nil {
//				return true
//			}
//
//			x, y, _ := client.CurrentChar.GetXYZ()
//			reg := models.GetRegion(x, y)
//			if reg != client.CurrentChar.CurrentRegion && client.CurrentChar.CurrentRegion != nil {
//				client.CurrentChar.CurrentRegion.CharsInRegion.Delete(client.CurrentChar.ObjectId)
//				reg.CharsInRegion.Store(client.CurrentChar.ObjectId, client.CurrentChar)
//				client.CurrentChar.CurrentRegion = reg
//
//				var info utils.packetByte
//				info.B = serverpackets.CharInfo(client.CurrentChar)
//				g.BroadToAroundPlayers(client, info)
//				BroadCastToMe(g, client.CurrentChar)
//				logger.Info.Println(client.CurrentChar.ObjectId, " change Region ")
//			}
//
//			return true // if false, Range stops
//		})
//
//		time.Sleep(1 * time.Second)
//	}
//}

//func BroadCastToMe(g *GameServer, my *models.Character) {
//	x, y, z := my.GetXYZ()
//	reg := models.GetRegion(x, y,z)
//	var charIds []int32
//
//	for _, iii := range reg.Sur {
//		iii.CharsInRegion.Range(func(key, value interface{}) bool {
//			val := value.(*models.Character)
//			if val.ObjectId != my.ObjectId {
//				charIds = append(charIds, val.ObjectId)
//			}
//			return true
//		})
//	}
//
//	if charIds == nil {
//		return
//	}
//
//	var me *models.ClientCtx
//
//	g.clients.Range(func(k, v interface{}) bool {
//		client := v.(*models.ClientCtx)
//		if client.CurrentChar.ObjectId == my.ObjectId {
//			me = client
//			return false
//		}
//		return true
//	})
//
//	if me == nil {
//		return // todo need log
//	}
//	for _, v := range charIds {
//		var info utils.packetByte
//		info.B = serverpackets.CharInfo(g.OnlineCharacters.Char[v])
//		me.AddLenghtAndSand(info.GetData(), true)
//	}
//}
