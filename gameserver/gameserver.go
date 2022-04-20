package gameserver

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"log"
)

var OnlineCharacters *models.OnlineCharacters

func GetNetConnByCharacterObjectId(objId int32) interfaces.ReciverAndSender {
	return OnlineCharacters.Char[objId].Conn
}
func GetNetConnByCharacterName(name string) interfaces.ReciverAndSender {
	for i, v := range OnlineCharacters.Char {
		if v.GetName() == name {
			return OnlineCharacters.Char[i].Conn
		}
	}
	return nil
}

func AddOnlineChar(character interfaces.CharacterI) {
	ch, ok := character.(*models.Character)
	if !ok {
		panic("addOnlineCharPanic")
	}
	OnlineCharacters.Char[character.GetObjectId()] = ch
}

func CharOffline(client interfaces.ReciverAndSender) {
	OnlineCharacters.Mu.Lock()
	delete(OnlineCharacters.Char, client.GetCurrentChar().GetObjectId())
	OnlineCharacters.Mu.Unlock()
	client.GetCurrentChar().GetCurrentRegion().DeleteVisibleChar(client.GetCurrentChar())

	client.GetCurrentChar().CloseChannels()

	//todo close all character goroutine, save info in DB
	log.Println("Socket Close For: ", client.GetCurrentChar().GetName())
}

//
//import (
//	"fmt"
//	"l2gogameserver/gameserver/broadcast"
//	"l2gogameserver/gameserver/interfaces"
//	"l2gogameserver/gameserver/models"
//	"l2gogameserver/gameserver/serverpackets"
//	"l2gogameserver/packets"
//	"l2gogameserver/utils"
//	"log"
//	"net"
//)
//
//type GameServer struct {
//	clientsListener  net.Listener
//	OnlineCharacters *models.OnlineCharacters
//
//	//clients          sync.Map
//}
//
//var GS *GameServer
//
////func (g *GameServer) AddClient(c *models.Client) {
////	g.clients.Store(c.Account.Login, c)
////}
//func New() *GameServer {
//	return &GameServer{}
//}
//
//func (g *GameServer) Start() {
//	var err error
//	GS = g
//	/* #nosec */
//	g.clientsListener, err = net.Listen("tcp4", ":7777")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	var onlineChars broadcast.OnlineCharacters
//	onlineChars.Char = make(map[int32]interfaces.ReciverAndSender)
//	//g.OnlineCharacters = &onlineChars
//
//	//go g.Tick()
//	defer g.clientsListener.Close()
//	for {
//		client := models.NewClient()
//		client.Socket, err = g.clientsListener.Accept()
//
//		if err != nil {
//			fmt.Println("Couldn't accept the incoming connection.", err)
//			continue
//		}
//
//		//g.AddClient(client)
//		go g.Handler(client)
//
//	}
//}
//
//func (g *GameServer) ChannelListener(client interfaces.ReciverAndSender) {
//	ch, ok := client.(*models.Client)
//	if !ok {
//		panic("ChannelListenerPanic")
//	}
//
//	for q := range ch.CurrentChar.ChannelUpdateShadowItem {
//		pkg := serverpackets.ItemUpdate(client, q.UpdateType, q.ObjId)
//		i := client.CryptAndReturnPackageReadyToShip(pkg)
//		client.Send(i)
//		if q.UpdateType == models.UpdateTypeRemove {
//			g.BroadCastUserInfoInRadius(client, 2000)
//		}
//	}
//}
//
//func (g *GameServer) NpcListener(client interfaces.ReciverAndSender) {
//	ch, ok := client.(*models.Client)
//	if !ok {
//		panic("NpcListenerPanic")
//	}
//	for q := range ch.CurrentChar.NpcInfo {
//		buff := packets.Get()
//		for i := range q {
//			pkg := serverpackets.NpcInfo(q[i])
//			buff.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
//		}
//		client.Send(buff.Bytes())
//		packets.Put(buff)
//	}
//}
//
//func (g *GameServer) MoveListener(client interfaces.ReciverAndSender) {
//	ch, ok := client.(*models.Client)
//	if !ok {
//		panic("NpcListenerPanic")
//	}
//
//	pkg := utils.GetPacketByte()
//	defer pkg.Release()
//
//	for q := range ch.CurrentChar.CharInfoTo {
//		pkg.SetData(serverpackets.CharInfo(ch.CurrentChar))
//		for _, v := range q {
//			g.OnlineCharacters.Mu.Lock()
//			g.OnlineCharacters.Char[v].Conn.EncryptAndSend(pkg.GetData())
//			g.OnlineCharacters.Mu.Unlock()
//		}
//	}
//
//	pkg.Free()
//
//	for q := range ch.CurrentChar.DeleteObjectTo {
//		pkg.SetData(serverpackets.DeleteObject(ch.CurrentChar))
//		for _, v := range q {
//			g.OnlineCharacters.Mu.Lock()
//			g.OnlineCharacters.Char[v].Conn.EncryptAndSend(pkg.GetData())
//			g.OnlineCharacters.Mu.Unlock()
//		}
//	}
//
//}
//
//func (g *GameServer) charOffline(client interfaces.ReciverAndSender) {
//	g.OnlineCharacters.Mu.Lock()
//	delete(g.OnlineCharacters.Char, client.GetCurrentChar().GetObjectId())
//	g.OnlineCharacters.Mu.Unlock()
//	client.GetCurrentChar().GetCurrentRegion().DeleteVisibleChar(client.GetCurrentChar())
//
//	client.GetCurrentChar().CloseChannels()
//
//	//todo close all character goroutine, save info in DB
//	log.Println("Socket Close For: ", client.GetCurrentChar().GetName())
//}
//
//func (g *GameServer) addOnlineChar(character interfaces.CharacterI) {
//	ch, ok := character.(*models.Character)
//	if !ok {
//		panic("addOnlineCharPanic")
//	}
//	g.OnlineCharacters.Mu.Lock()
//	g.OnlineCharacters.Char[character.GetObjectId()] = ch
//	g.OnlineCharacters.Mu.Unlock()
//}
//
//func (g *GameServer) brdsct(me *models.Client, pkg utils.PacketByte) {
//	charsIds := models.GetAroundPlayer(me.CurrentChar)
//	for _, v := range charsIds {
//		v.EncryptAndSend(pkg.GetData())
//		//me.Send(me.CryptAndReturnPackageReadyToShip(serverpackets.CharInfo(v)))
//	}
//}
