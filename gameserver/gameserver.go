package gameserver

import (
	"fmt"
	"github.com/jackc/pgx"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"log"
	"net"
	"sync"
	"time"
)

type GameServer struct {
	clientsListener  net.Listener
	clients          []*models.Client
	Socket           net.Conn
	database         *pgx.Conn
	onlineCharacters *OnlineCharacters
}
type OnlineCharacters struct {
	char map[int32]*models.Character
	mu   sync.Mutex
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

		log.Fatal("Failed to connect to database: ", err.Error())
	} else {
		fmt.Println("Successful database connection")
	}
	g.clientsListener, err = net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal("Failed to connect to port 7777:", err.Error())
	} else {
		fmt.Println("Login server is listening on port 7777")
	}
	var onlineChars OnlineCharacters
	x := make(map[int32]*models.Character)
	onlineChars.char = x
	g.onlineCharacters = &onlineChars
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

type PacketByte struct {
	b []byte
}

func (i *PacketByte) GetB() []byte {
	cl := make([]byte, len(i.b))
	_ = copy(cl, i.b)
	return cl
}
func (i *PacketByte) SetB(v []byte) {
	cl := make([]byte, len(v))
	i.b = cl
	copy(i.b, v)
}

func Broad(g *GameServer, my *models.Character, pkg PacketByte) {

	reg := models.GetRegion(my.Coordinates.X, my.Coordinates.Y)
	var charIds []int32
	for _, iii := range reg.Sur {
		iii.CharsInRegion.Range(func(key, value interface{}) bool {
			val := value.(*models.Character)
			charIds = append(charIds, val.CharId)
			return true
		})
	}

	if len(charIds) == 1 { //todo я всегда буду в этом регионе поэтому 1
		return
	}

	for _, p := range g.clients {
		for _, w := range charIds {
			if p.CurrentChar.CharId == w && p.CurrentChar.CharId != my.CharId {
				p.Send(pkg.GetB(), true)
			}
		}

	}
}
func (g *GameServer) addOnlineChar(character *models.Character) {
	g.onlineCharacters.mu.Lock()
	g.onlineCharacters.char[character.CharId] = character
	g.onlineCharacters.mu.Unlock()
}
func (g *GameServer) Tick() {

	for {
		for _, v := range g.onlineCharacters.char {
			x, y, _ := v.GetXYZ()
			reg := models.GetRegion(x, y)
			if reg != v.CurrentRegion && v.CurrentRegion != nil {
				v.CurrentRegion.CharsInRegion.Delete(v.CharId)
				reg.CharsInRegion.Store(v.CharId, v)
				v.CurrentRegion = reg

				var info PacketByte
				info.b = serverpackets.NewCharInfo(v)
				Broad(g, v, info)
				BroadCastToMe(g, v)
				log.Println(v.CharId, " change Region ")

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
		var info PacketByte
		info.b = serverpackets.NewCharInfo(g.onlineCharacters.char[v])
		me.Send(info.GetB(), true)
	}

}
