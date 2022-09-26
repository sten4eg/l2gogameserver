package server

import (
	"fmt"
	"github.com/puzpuzpuz/xsync"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/handlers"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/loginserver"
	"l2gogameserver/loginserver/network/gs2ls"
	"net"
)

type GameServer struct {
	clientsListener *net.TCPListener
	clients         *xsync.MapOf[interfaces.ClientInterface]
	waitingClients  *xsync.MapOf[interfaces.ClientInterface]
	loginServer     *loginserver.LoginServer
}

func New() *GameServer {
	gs := &GameServer{}
	gs.clients = xsync.NewMapOf[interfaces.ClientInterface]()
	gs.waitingClients = xsync.NewMapOf[interfaces.ClientInterface]()
	ls := loginserver.GetLoginServerInstance()
	gs.loginServer = ls

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

// AddClient вернёт true если клинта небыло в мапе, false если клиент был обновлен в мапе
func (g *GameServer) AddClient(login string, clientI interfaces.ClientInterface) bool {
	_, loaded := g.clients.LoadOrStore(login, clientI)
	return !loaded
}

func (g *GameServer) DeleteClient(login string) {
	g.clients.Delete(login)
}

func (g *GameServer) AddWaitClient(login string, clientI interfaces.ClientInterface) {
	playOk1, playOk2, loginOk1, loginOk2 := clientI.GetSessionKey()
	g.loginServer.Send(gs2ls.PlayerAuthRequest(login, playOk1, playOk2, loginOk1, loginOk2))
	g.waitingClients.Store(login, clientI)
}

func (g *GameServer) ExistsWaitClient(login string) bool {
	_, exist := g.waitingClients.Load(login)
	return exist
}

func (g *GameServer) KickClientByLogin(login string) {
	//for i,v := range g.onlineCharacters
}

func (g *GameServer) GetClient(login string) interfaces.ClientInterface {
	v, ok := g.clients.Load(login)
	if !ok {
		return nil
	}
	return v
}

func (g *GameServer) RemoveWaitingClient(login string) {
	g.waitingClients.Delete(login)
}
func (g *GameServer) RemoveClient(login string) {
	g.clients.Delete(login)
}
func (g *GameServer) SendLogout(login string) {
	msg := gs2ls.PlayerLogout(login)
	g.loginServer.Send(msg)
}
