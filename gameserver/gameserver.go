package gameserver

import (
	"fmt"
	"github.com/jackc/pgx"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"log"
	"net"
	"time"
)

type GameServer struct {
	clientsListener  net.Listener
	clients          []*models.Client
	Socket           net.Conn
	database         *pgx.Conn
	OnlineCharacters *models.OnlineCharacters
}

func New() *GameServer {
	return &GameServer{}
}
func (g *GameServer) Init() {

	var err error
	globalConfig := config.Read()
	dbConfig := pgx.ConnConfig{
		Host:              globalConfig.LoginServer.Database.Host,
		Port:              globalConfig.LoginServer.Database.Port,
		Database:          globalConfig.LoginServer.Database.Name,
		User:              globalConfig.LoginServer.Database.User,
		Password:          globalConfig.LoginServer.Database.Password,
		TLSConfig:         nil,
		FallbackTLSConfig: nil,
	}

	g.database, err = pgx.Connect(dbConfig)
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Successful database connection")
	}
	g.clientsListener, err = net.Listen("tcp", ":7777")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Login server is listening on port 7777")
	}
	var onlineChars models.OnlineCharacters
	x := make(map[int32]*models.Character)
	onlineChars.Char = x
	g.OnlineCharacters = &onlineChars
}

func (g *GameServer) Start() {
	go g.Tick()
	defer g.clientsListener.Close()
	for {
		var err error
		client := models.NewClient()
		client.Socket, err = g.clientsListener.Accept()
		g.clients = append(g.clients, client)
		if err != nil {
			fmt.Println("Couldn't accept the incoming connection.")
			continue
		} else {
			go g.handler(client)
		}
	}
}
func kickClient(client *models.Client) {
	err := client.Socket.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Socket Close For: ", client.CurrentChar.CharName)
}

func Broad(my *models.Client, pkg models.PacketByte) {
	reg := models.GetRegion(my.CurrentChar.Coordinates.X, my.CurrentChar.Coordinates.Y)
	var charIds []int32
	for _, iii := range reg.Sur {
		iii.CharsInRegion.Range(func(key, value interface{}) bool {
			val := value.(*models.Character)
			if val.CharId == my.CurrentChar.CharId {
				return true
			}
			val.Conn.Send(pkg.GetB(), true)
			charIds = append(charIds, val.CharId)
			return true
		})
	}

	//if len(charIds) == 1 { //todo я всегда буду в этом регионе поэтому 1
	//	return
	//}

	//for _, p := range g.clients { // 3_000_000
	//	for _, w := range charIds { // 300
	//		if p.CurrentChar.CharId == w && p.CurrentChar.CharId != my.CharId {
	//			p.Send(pkg.GetB(), true)
	//		}
	//	}

	//}
}
func (g *GameServer) addOnlineChar(character *models.Character) {
	g.OnlineCharacters.Mu.Lock()
	g.OnlineCharacters.Char[character.CharId] = character
	g.OnlineCharacters.Mu.Unlock()
}
func (g *GameServer) Tick() {

	for {
		for _, v := range g.clients {
			if v.CurrentChar.Coordinates == nil {
				continue
			}
			x, y, _ := v.CurrentChar.GetXYZ()
			reg := models.GetRegion(x, y)
			if reg != v.CurrentChar.CurrentRegion && v.CurrentChar.CurrentRegion != nil {
				v.CurrentChar.CurrentRegion.CharsInRegion.Delete(v.CurrentChar.CharId)
				reg.CharsInRegion.Store(v.CurrentChar.CharId, v.CurrentChar)
				v.CurrentChar.CurrentRegion = reg

				var info models.PacketByte
				info.B = serverpackets.NewCharInfo(v.CurrentChar)
				Broad(v, info)
				BroadCastToMe(g, v.CurrentChar)
				log.Println(v.CurrentChar.CharId, " change Region ")

			}
		}
		time.Sleep(1 * time.Second)
	}
}

func BroadCastToMe(g *GameServer, my *models.Character) {
	reg := models.GetRegion(my.Coordinates.X, my.Coordinates.Y)
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

	var me *models.Client
	for _, p := range g.clients {
		if p.CurrentChar.CharId == my.CharId {
			me = p
			break
		}
	}

	for _, v := range charIds {
		var info models.PacketByte
		info.B = serverpackets.NewCharInfo(g.OnlineCharacters.Char[v])
		me.Send(info.GetB(), true)
	}

}
