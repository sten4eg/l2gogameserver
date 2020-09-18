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
	for {
		opcode, data, err := client.Receive()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Closing the connection...")
			break
		}

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

		default:
			fmt.Println("Can't recognize the packet sent by the gameserver")
		}

	}
}
