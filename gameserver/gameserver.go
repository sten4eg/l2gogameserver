package gameserver

import (
	"fmt"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"log"
	"net"
	"sync"
	"time"
)

type GameServer struct {
	clientsListener  net.Listener
	clients          sync.Map
	Socket           net.Conn
	OnlineCharacters *models.OnlineCharacters
}

func (g *GameServer) AddClient(c *models.Client) {
	g.clients.Store(c.Account.Login, c)
}
func New() *GameServer {
	return &GameServer{}
}

func (g *GameServer) Start() {
	var err error

	/* #nosec */
	g.clientsListener, err = net.Listen("tcp", ":7777")
	if err != nil {
		panic(err.Error())
	}

	var onlineChars models.OnlineCharacters

	onlineChars.Char = make(map[int32]*models.Character)
	g.OnlineCharacters = &onlineChars

	go g.Tick()
	defer g.clientsListener.Close()
	for {
		client := models.NewClient()
		client.Socket, err = g.clientsListener.Accept()

		if err != nil {
			fmt.Println("Couldn't accept the incoming connection.", err)
			continue
		} else {
			g.AddClient(client)
			go g.handler(client)
		}
	}
}
func kickClient(client *models.Client) {
	err := client.Socket.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Socket Close For: ", client.CurrentChar.CharName)
}

func (g *GameServer) BroadToAroundPlayers(my *models.Client, pkg models.PacketByte) {

	charsIds := models.GetAroundPlayers(my.CurrentChar)
	for _, v := range charsIds {
		_ = g.OnlineCharacters.Char[v].Conn.Send(pkg.GetB(), true)
	}

}
func (g *GameServer) addOnlineChar(character *models.Character) {
	g.OnlineCharacters.Mu.Lock()
	g.OnlineCharacters.Char[character.CharId] = character
	g.OnlineCharacters.Mu.Unlock()
}

func (g *GameServer) Tick() {
	for {
		g.clients.Range(func(k, v interface{}) bool {
			client := v.(*models.Client)
			if client.CurrentChar.Coordinates == nil {
				return true
			}

			x, y, _ := client.CurrentChar.GetXYZ()
			reg := models.GetRegion(x, y)
			if reg != client.CurrentChar.CurrentRegion && client.CurrentChar.CurrentRegion != nil {
				client.CurrentChar.CurrentRegion.CharsInRegion.Delete(client.CurrentChar.CharId)
				reg.CharsInRegion.Store(client.CurrentChar.CharId, client.CurrentChar)
				client.CurrentChar.CurrentRegion = reg

				var info models.PacketByte
				info.B = serverpackets.CharInfo(client.CurrentChar)
				g.BroadToAroundPlayers(client, info)
				BroadCastToMe(g, client.CurrentChar)
				log.Println(client.CurrentChar.CharId, " change Region ")
			}

			return true // if false, Range stops
		})

		time.Sleep(1 * time.Second)
	}
}

func BroadCastToMe(g *GameServer, my *models.Character) {
	x, y, _ := my.GetXYZ()
	reg := models.GetRegion(x, y)
	var charIds []int32

	for _, iii := range reg.Sur {
		iii.CharsInRegion.Range(func(key, value interface{}) bool {
			val := value.(*models.Character)
			if val.CharId != my.CharId {
				charIds = append(charIds, val.CharId)
			}
			return true
		})
	}

	if charIds == nil {
		return
	}

	var me *models.Client

	g.clients.Range(func(k, v interface{}) bool {
		client := v.(*models.Client)
		if client.CurrentChar.CharId == my.CharId {
			me = client
			return false
		}
		return true
	})

	if me == nil {
		return // todo need log
	}
	for _, v := range charIds {
		var info models.PacketByte
		info.B = serverpackets.CharInfo(g.OnlineCharacters.Char[v])
		_ = me.Send(info.GetB(), true)
	}
}
