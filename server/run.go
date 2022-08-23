package server

import (
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/handlers"
	"l2gogameserver/gameserver/models"
	"net"
)

type GameServer struct {
	clientsListener *net.TCPListener
	//OnlineCharacters *models.OnlineCharacters
	//clients          sync.Map
}

func New() *GameServer {
	return new(GameServer)
}

func (g *GameServer) Start() {
	var err error

	addr := new(net.TCPAddr)
	addr.Port = 7777
	addr.IP = net.IP{127, 0, 0, 1}

	/* #nosec */
	g.clientsListener, err = net.ListenTCP("tcp4", addr)
	if err != nil {
		logger.Error.Panicln(err.Error())
	}
	defer g.clientsListener.Close()

	var onlineChars models.OnlineCharacters
	onlineChars.Char = make(map[int32]*models.Character)
	gameserver.OnlineCharacters = &onlineChars

	//go g.Tick()

	for {
		client := models.NewClient()
		conn, err := g.clientsListener.AcceptTCP()
		if err != nil {
			fmt.Println("Couldn't accept the incoming connection.", err)
			continue
		}
		client.SetConn(conn)

		//g.AddClient(client) //todo надо ли добавлять клиентов в отдельную мапу или массив?
		go handlers.Handler(client)
	}
}
