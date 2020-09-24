package gameserver

import (
	"fmt"
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"log"
	"net"
)

type GameServer struct {
	clientsListener net.Listener
	clients         []*models.Client
	Socket          net.Conn
}

func New() *GameServer {
	return &GameServer{}
}

func (g *GameServer) Init() {
	var err error
	g.clientsListener, err = net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal("Failed to connect to port 7777:", err.Error())
	} else {
		fmt.Println("Login server is listening on port 7777")
	}
}

func (g *GameServer) Start() {
	defer g.clientsListener.Close()

	done := make(chan bool)

	go func() {
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
		done <- true
	}()
	for i := 0; i < 1; i++ {
		<-done
	}
}

func (g *GameServer) handleClientPackets(client *models.Client) {
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
			pkg := serverpackets.NewKeyPacket()
			err := client.Send(pkg, false)
			if err != nil {
				log.Println(err)
			}
		case 00:
			fmt.Println("A game server sent a request to register")
		case 43:
			clientpackets.NewAuthLogin(data)
			pkg := serverpackets.NewCharSelectionInfo()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		case 19:
			pkg := serverpackets.NewCharacterSuccess()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		case 18:
			clientpackets.NewCharSelected(data)
			pkg := serverpackets.NewSSQInfo()
			//		d := client.Ssend(pkg)
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			log.Println("sendSSQ")
			pkg = serverpackets.NewCharSelected()
			//			dd := client.Ssend(pkg)
			//			q := append(d,dd...)
			//			client.SSS(q)
			err = client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
			log.Println("CharSelected")
		case 208:
			if i == 0 {
				pkg := serverpackets.NewExSendManorList()
				err := client.Send(pkg, true)
				if err != nil {
					log.Println(err)
				}
				log.Println("ExSendManorList")
			}
			i++
		case 193:
			pkg := serverpackets.NewObservationReturn()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		case 17:
			pkg := serverpackets.NewUserInfo()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}

			//	p := serverpackets.NewItemList()
			//	client.Send(p, true)
			//	p = serverpackets.NewExGetBookMarkInfoPacket()
			//	client.Send(p, true)
			//	p = serverpackets.NewShortCutInit()
			//	client.Send(p, true)
			//	p = serverpackets.NewExBasicActionList()
			//	client.Send(p, true)
			//	p = serverpackets.NewSkillList()
			//	client.Send(p, true)
			////	p = serverpackets.NewHennaInfo()
			////	client.Send(p, true)
			//	p = serverpackets.NewQuestList()
			//	client.Send(p, true)
			//	p = serverpackets.NewEtcStatusUpdate()
			//	client.Send(p, true)
			//	p = serverpackets.NewExStorageMaxCount()
			//	client.Send(p, true)
			//	p = serverpackets.NewFriendList()
			//	client.Send(p, true)
			//	//WELCOM_TO_LINEAGE
			//
			//	p = serverpackets.NewSkillCoolTime()
			//	client.Send(p, true)
			//	p = serverpackets.NewExVoteSystemInfo()
			//	client.Send(p, true)
			//	p = serverpackets.NewExNevitAdventPointInfoPacket()
			//	client.Send(p, true)
			//	p = serverpackets.NewExNevitAdventTimeChange()
			//	client.Send(p, true)
			//	p = serverpackets.NewExShowContactList()
			//	client.Send(p, true)
			//
			//	p = serverpackets.NewActionFailed()
			//	client.Send(p, true)
			log.Println("NewUserInfo")
		case 166:
			pkg := serverpackets.NewSkillCoolTime()
			err := client.Send(pkg, true)
			if err != nil {
				log.Println(err)
			}
		default:
			log.Println("Not Found case with opcode: ", opcode)
		}

	}
}
