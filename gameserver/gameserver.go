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
)

type GameServer struct {
	clientsListener net.Listener
	clients         []*models.Client
	Socket          net.Conn
	database        *pgx.Conn
	account         *serverpackets.Account
	mp              map[int32]serverpackets.Character
}

func New() *GameServer {
	return &GameServer{}
}
func (g *GameServer) Init() {
	gm := make(map[int32]serverpackets.Character)
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
	g.mp = gm
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

}

func (g *GameServer) Start() {
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
}

func (g *GameServer) handleClientPackets(client *models.Client) {
	defer kickClient()
	var i = 0

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
			//var pkg []byte
			login := clientpackets.NewAuthLogin(data)
			g.account = serverpackets.NewCharSelectionInfo(g.database, client, login) //TODO пересмотреть
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
			reason, err := clientpackets.NewCharacterCreate(data, g.database, client.CurrentChar.Login)
			if err != nil {
				serverpackets.NewCharCreateFail(client, reason)
			} else {
				log.Println("sozdal")
			}
		case 18:
			g.account.CharSlot = clientpackets.NewCharSelected(data)
			pkg := serverpackets.NewSSQInfo()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			log.Println("sendSSQ")

			client.CurrentChar.CharId = serverpackets.NewCharSelected(g.account.Char[g.account.CharSlot], client)
			g.mp[g.account.CharSlot] = *g.account.Char[g.account.CharSlot]
			err = client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
			log.Println("Send CharSelected")
		case 208:
			if i == 0 {
				serverpackets.NewExSendManorList(client)
				err := client.SimpleSend(client.Buffer.Bytes(), true)
				if err != nil {
					log.Println(err)
				}
				log.Println("Send ExSendManorList")
			}
			i++
		case 193:
			serverpackets.NewObservationReturn(g.account.Char[g.account.CharSlot], client)
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
			pkg := serverpackets.NewUserInfo(g.account.Char[g.account.CharSlot])
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
			serverpackets.NewMoveToLocation(location, client, client.CurrentChar.CharId)
			err := client.SimpleSend(client.Buffer.Bytes(), true)
			if err != nil {
				log.Println(err)
			}
			client.CurrentChar.Spawn.Z = location.TargetZ
			client.CurrentChar.Spawn.X = location.TargetX
			client.CurrentChar.Spawn.Y = location.TargetY
			client.Buffer.Reset()
			Broad(g)
			log.Println("Send NewMoveToLocation")
		case 73:
			say := clientpackets.NewSay(data)
			pkg := serverpackets.NewCreatureSay(say)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}
	}
}

func Broad(g *GameServer) {

	for _, p := range g.clients {
		CI := serverpackets.NewCharInfo(p.CurrentChar)
		err := p.Send(CI, true)
		if err != nil {
			log.Fatal(err)
		}
	}

}
