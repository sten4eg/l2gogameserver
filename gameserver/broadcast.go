package gameserver

import (
	"fmt"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/chat"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/utils"
)

// BroadCastToAroundPlayersInRadius отправляет всем персонажам в радиусе radius
// информацию из пакета pkg
func (g *GameServer) BroadCastToAroundPlayersInRadius(my *models.Client, pkg utils.PacketByte, radius int32) {

	charsIds := models.GetAroundPlayersInRadius(my.CurrentChar, radius)
	for _, v := range charsIds {
		g.OnlineCharacters.Char[v.ObjectId].Conn.Send(pkg.GetB(), true)
	}
}

func (g *GameServer) BroadCastToAroundPlayers(my *models.Client, pkg utils.PacketByte) {
	charsIds := models.GetAroundPlayer(my.CurrentChar)
	for _, v := range charsIds {
		v.Conn.Send(pkg.GetB(), true)
	}
}

// BroadCastUserInfoInRadius отправляет всем персонажам в радиусе radius
// информацию о персонаже, Самому персонажу отправляет полный UserInfo
func (g *GameServer) BroadCastUserInfoInRadius(me *models.Client, radius int32) {
	ui := serverpackets.UserInfo(me)
	me.Send(ui, true)

	charsIds := models.GetAroundPlayersInRadius(me.CurrentChar, radius)
	if len(charsIds) == 0 {
		return
	}

	var ci utils.PacketByte
	ci.SetB(serverpackets.CharInfo(me.CurrentChar))

	var exUi utils.PacketByte
	exUi.SetB(serverpackets.ExBrExtraUserInfo(me.CurrentChar))

	g.OnlineCharacters.Mu.Lock()
	for _, v := range charsIds {
		g.OnlineCharacters.Char[v.ObjectId].Conn.Send(ci.GetB(), true)
		g.OnlineCharacters.Char[v.ObjectId].Conn.Send(exUi.GetB(), true)
	}
	g.OnlineCharacters.Mu.Unlock()
}

func (g *GameServer) BroadCastChat(me *models.Client, say models.Say) {
	switch say.Type {
	case chat.All:
		cs := serverpackets.CreatureSay(&say, me.CurrentChar)
		var pb utils.PacketByte
		pb.SetB(cs)
		me.SSend(me.CryptAndReturnPackageReadyToShip(pb.GetB()))
		g.BroadCastToAroundPlayersInRadius(me, pb, chat.AllChatRange)
	case chat.Tell:
		cs := serverpackets.CreatureSay(&say, me.CurrentChar)
		var pb utils.PacketByte
		pb.SetB(cs)
		ok := g.BroadCastToCharacterByName(pb, say.To)
		if ok {
			me.SSend(me.CryptAndReturnPackageReadyToShip(pb.GetB()))
		} else {
			// systemMSG что не найден перс
		}
	case chat.Shout:
		cs := serverpackets.CreatureSay(&say, me.CurrentChar)
		var pb utils.PacketByte
		pb.SetB(cs)
		me.SSend(me.CryptAndReturnPackageReadyToShip(pb.GetB()))
		g.BroadCastToAroundPlayersInRadius(me, pb, chat.ShoutChatRange)
	case chat.SpecialCommand:
		if me.CurrentChar.Target == 0 {
			return
		}
		qwe := g.OnlineCharacters.Char[me.CurrentChar.Target]
		q := models.CalculateDistance(qwe.Coordinates.X, qwe.Coordinates.Y, qwe.Coordinates.Z, me.CurrentChar.Coordinates.X, me.CurrentChar.Coordinates.Y, me.CurrentChar.Coordinates.Z, false, false)
		say.Text = fmt.Sprintf("%f", q)
		say.Type = chat.All

		cs := serverpackets.CreatureSay(&say, me.CurrentChar)
		var pb utils.PacketByte
		pb.SetB(cs)
		me.SSend(me.CryptAndReturnPackageReadyToShip(pb.GetB()))
		g.BroadCastToAroundPlayersInRadius(me, pb, chat.AllChatRange)
	}
}

// BroadCastToCharacterByName отправляет pkg персонажу с ником to
// true если отправлен, false если персонаж не найден
func (g *GameServer) BroadCastToCharacterByName(pkg utils.PacketByte, to string) bool {
	g.OnlineCharacters.Mu.Lock()
	defer g.OnlineCharacters.Mu.Unlock()
	for _, v := range g.OnlineCharacters.Char {
		if v.CharName == to {
			v.Conn.Send(pkg.GetB(), true)
			return true
		}
	}
	return false
}

// GetCharInfoAboutCharactersInRadius отправляет me CharInfo персонажей
// в радиусе radius
func (g *GameServer) GetCharInfoAboutCharactersInRadius(me *models.Client, radius int32) {
	charsIds := models.GetAroundPlayersInRadius(me.CurrentChar, radius)
	for _, v := range charsIds {
		me.SSend(me.CryptAndReturnPackageReadyToShip(serverpackets.CharInfo(v)))
	}
}

// GetCharInfoAboutCharacters отправляет me CharInfo персонажей
func (g *GameServer) GetCharInfoAboutCharacters(me *models.Client) {
	charsIds := models.GetAroundPlayer(me.CurrentChar)
	for _, v := range charsIds {
		me.SSend(me.CryptAndReturnPackageReadyToShip(serverpackets.CharInfo(v)))
	}
}

func (g *GameServer) Checkaem(client *models.Client, l models.BackwardToLocation) {
	var ut utils.PacketByte
	ut.SetB(serverpackets.MoveToLocation(&l, client))
	client.SSend(client.CryptAndReturnPackageReadyToShip(ut.GetB()))
	g.BroadCastToAroundPlayers(client, ut)
}

//func (g *GameServer) Tick() {
//	for {
//		g.clients.Range(func(k, v interface{}) bool {
//			client := v.(*models.Client)
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
//				var info utils.PacketByte
//				info.B = serverpackets.CharInfo(client.CurrentChar)
//				g.BroadToAroundPlayers(client, info)
//				BroadCastToMe(g, client.CurrentChar)
//				log.Println(client.CurrentChar.ObjectId, " change Region ")
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
//	var me *models.Client
//
//	g.clients.Range(func(k, v interface{}) bool {
//		client := v.(*models.Client)
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
//		var info utils.PacketByte
//		info.B = serverpackets.CharInfo(g.OnlineCharacters.Char[v])
//		me.Send(info.GetB(), true)
//	}
//}
