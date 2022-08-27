package server

import (
	"fmt"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/handlers"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/loginserver"
	"log"
	"net"
	"sync"
)

type GameServer struct {
	clientsListener *net.TCPListener
	clients         sync.Map
}

func New() *GameServer {
	gs := &GameServer{}
	gs.clients.Store("q", "v")
	ls := loginserver.GetLoginServerInstance()
	ls.AttachGs(gs)
	return gs
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
		go handlers.Handler(client, g)
	}
}

func (g *GameServer) AddClient(login string, clientI interfaces.ClientInterface) {
	client, ok := clientI.(*models.ClientCtx)
	if !ok || client == nil {
		log.Println("нету клиента для добавления")
		return
	}

	//g.onlineCharacters[login] = client
}

func (g *GameServer) KickClientByLogin(login string) {
	//for i,v := range g.onlineCharacters
}
