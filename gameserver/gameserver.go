package gameserver

import (
	"fmt"
	"github.com/jackc/pgx"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"log"
	"net"
	"os"
	"runtime/pprof"
	"runtime/trace"
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
			go g.handleClientPackets(client)
		}
	}
}
func kickClient() {
	f, err := os.Create("f.pprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	//runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	ff, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer ff.Close()

	err = trace.Start(ff)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
}

func (g *GameServer) handleClientPackets(client *models.Client) {
	defer kickClient()

	for {
		opcode, data, err := client.Receive()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Closing the connection...")
			break
		}
		log.Println("income ", opcode)
		switch opcode {
		case 14:
			_ = clientpackets.NewprotocolVersion(data)
			serverpackets.NewKeyPacket(client)
			err := client.SimpleSend(client.Buffer.Bytes(), false)
			if err != nil {
				log.Println(err)
			}
			log.Println("Send NewKeyPacket")

		case 00:
			fmt.Println("A game server sent a request to register")
		case 43:
			login := clientpackets.NewAuthLogin(data)
			client.Account.Login = login
			client.Account = serverpackets.NewCharSelectionInfo(g.database, client)
			client.Account.Login = login
			err := client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
			log.Println("Send NewCharSelectionInfo")
		case 19:
			serverpackets.NewCharacterSuccess(client)
			err := client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
			log.Println("Send NewCharacterSuccess")
		case 12:
			reason := clientpackets.NewCharacterCreate(data, g.database, client.Account.Login)
			if reason != clientpackets.ReasonOk {
				serverpackets.NewCharCreateFail(client, reason)
				err := client.SimpleSend(client.Buffer.Bytes(), true)
				if err != nil {
					log.Println(err)
				}
			} else {
				serverpackets.NewCharCreateOk(client)
				err = client.SimpleSend(client.Buffer.Bytes(), true)
				if err != nil {
					log.Println(err)
				}
				log.Println("send NewCharCreateOk")
			}
		case 18:
			client.Account.CharSlot = clientpackets.NewCharSelected(data)
			pkg := serverpackets.NewSSQInfo()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			log.Println("sendSSQ")

			_ = serverpackets.NewCharSelected(client.Account.Char[client.Account.CharSlot], client) // return charId
			client.CC = client.Account.Char[client.Account.CharSlot]

			rg := models.GetRegion(client.CC.Coordinates.X, client.CC.Coordinates.Y)
			rg.AddVisibleObject(client.CC)
			client.CC.CurrentRegion = rg
			g.addOnlineChar(client.CC)
			err = client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
			log.Println("Send CharSelected")
		case 208:
			if len(data) >= 2 {
				switch data[0] {
				case 1:
					serverpackets.NewExSendManorList(client)
					err := client.SimpleSend(client.Buffer.Bytes(), true)
					if err != nil {
						log.Println(err)
					}
					log.Println("Send ExSendManorList")
				case 54:
					client.Account = serverpackets.NewCharSelectionInfo(g.database, client)
					err := client.SimpleSend(client.Buffer.Bytes(), true)
					if err != nil {
						log.Println(err)
					}
					log.Println("Send NewCharSelectionInfo")
				}

			}

		case 193:
			serverpackets.NewObservationReturn(client.CC, client)
			err := client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
		case 108:
			serverpackets.NewShowMiniMap(client)
			err := client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
		case 17:
			pkg := serverpackets.NewUserInfo(client.CC)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			pkg = serverpackets.NewExBrExtraUserInfo()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			pkg = serverpackets.NewSendMacroList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewItemList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewExQuestItemList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewGameGuardQuery()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewExGetBookMarkInfoPacket()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewShortCutInit()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewExBasicActionList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewSkillList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewHennaInfo()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewQuestList()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			pkg = serverpackets.NewStaticObject()
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			log.Println("Send NewUserInfo")
		case 166:
			pkg := serverpackets.NewSkillCoolTime()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		case 15:
			location := clientpackets.NewMoveBackwardToLocation(data)
			pkg := serverpackets.NewMoveToLocation(location, client)
			var info PacketByte
			info.SetB(pkg)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			Broad(g, client.CC, info)

			log.Println("Send NewMoveToLocation")
		case 73:
			//	say := clientpackets.NewSay(data)
			var info PacketByte
			//info.b = serverpackets.NewCreatureSay(say, client.CC)
			//err := client.Send(info.GetB(), true)
			//if err != nil {
			//	log.Println(err)
			//}
			info.b = serverpackets.NewCharInfo(client.CC)
			Broad(g, client.CC, info)
		case 89:
			clientpackets.NewValidationPosition(data, client.CC)
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}
	}
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
			if p.CC.CharId == w && p.CC.CharId != my.CharId {
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
		if p.CC.CharId == my.CharId {
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
