package loginserver

import (
	"fmt"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/crypt"
	"l2gogameserver/gameserver/crypt/blowfish"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

type LoginServer struct {
	conn       *net.TCPConn
	blowfish   *blowfish.Cipher
	gameServer gameServerInterface
	sync.Mutex
}

var loginServerInstance *LoginServer

type gameServerInterface interface {
	ExistsWaitClient(login string) bool
	GetClient(login string) interfaces.ClientInterface
	RemoveWaitingClient(login string)
	RemoveClient(login string)
}

func (ls *LoginServer) AttachGs(gs gameServerInterface) {
	ls.gameServer = gs
}

func GetLoginServerInstance() *LoginServer {
	return loginServerInstance
}

var initBlowfishKey = []byte{95, 59, 118, 46, 93, 48, 53, 45, 51, 49, 33, 124, 43, 45, 37, 120, 84, 33, 94, 91, 36, 0}

func HandlerInit() error {
	loginServerInstance = new(LoginServer)

	loginServerInstance.Lock()
	defer loginServerInstance.Unlock()

	loginServerInstance.SetConn()

	go loginServerInstance.Run()

	return nil
}

func (ls *LoginServer) SetConn() {
	port := config.GetLoginServerPort()
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return //err>
	}
	addr := new(net.TCPAddr)
	addr.Port = intPort
	addr.IP = net.IP{127, 0, 0, 1} //TODO IP тоже брать из конфига
	var conn *net.TCPConn
	for conn == nil {
		conn, err = net.DialTCP("tcp4", nil, addr)
		if err != nil {
			log.Printf("не удалось подключиться к логин-серверу : %v\n", err.Error())
		}
		time.Sleep(howLongNeedSleep())
	}

	ls.conn = conn
}

func (ls *LoginServer) tryReconnectToLS() {
	log.Println("попытка реконнекта к логин серверу")
	ls.Lock()
	defer ls.Unlock()
	if ls.conn == nil {
		log.Println("реконнект к логин серверу")
		ls.SetConn()

		go ls.Run()
	}
	return
}

func (ls *LoginServer) CloseConnAndReconnectLS() {
	ls.conn = nil
	ls.tryReconnectToLS()
}

func (ls *LoginServer) Run() {
	defer ls.CloseConnAndReconnectLS()

	ls.setBlowFish(initBlowfishKey)

	for {
		header := make([]byte, 2)

		n, err := ls.conn.Read(header)
		if err != nil {
			log.Println(err)
			return
		}
		if n != 2 {
			log.Println("Должно быть 2 байта размера")
			return
		}

		dataSize := (int(header[0]) | int(header[1])<<8) - 2

		data := make([]byte, dataSize)
		n, err = ls.conn.Read(data)
		if err != nil {
			log.Println(err)
			return
		}
		if n != dataSize {
			log.Println("Прочитанно байт меньше чем объявлено в размере пакета")
			return
		}

		for i := 0; i < dataSize; i += 8 {
			ls.blowfish.Decrypt(data, data, i, i)
		}

		ok := crypt.VerifyCheckSum(data, dataSize)
		if !ok {
			fmt.Println("Неверная контрольная сумма пакета, закрытие соединения.")
			return
		}

		ls.HandlePacket(data)
	}
}

func generateNewBlowFish() []byte {
	bfk := make([]byte, 40)
	_, _ = rand.Read(bfk)

	// главное чтобы не 0
	if bfk[0] == 0 {
		const blowFishFirstByte byte = 113
		bfk[0] = blowFishFirstByte
	}
	return bfk
}

func (ls *LoginServer) setBlowFish(blowfishKey []byte) {
	c := make([]byte, len(blowfishKey))
	copy(c, blowfishKey)
	cipher, err := blowfish.NewCipher(c)
	if err != nil {
		panic(err)
	}
	ls.blowfish = cipher
}

func (ls *LoginServer) Send(buf *packets.Buffer) {
	size := buf.Len() + 4
	size = (size + 8) - (size % 8) // padding

	data := make([]byte, size)
	copy(data, buf.Bytes())

	defer packets.Put(buf)

	rs := crypt.AppendCheckSum(data, size)

	for i := 0; i < size; i += 8 {
		ls.blowfish.Encrypt(rs, rs, i, i)
	}

	rs = rs[:size]
	leng := len(rs) + 2

	s, f := byte(leng>>8), byte(leng&0xff)
	res := append([]byte{f, s}, rs...)

	_, err := ls.conn.Write(res)

	if err != nil {
		log.Println(err)
	}

}

var attempt int

func howLongNeedSleep() time.Duration {
	attempt++
	if attempt < 5 {
		return time.Duration(attempt) * time.Second
	}
	return time.Second * 5
}

func (ls *LoginServer) ExistsWaitClientOnGameServer(login string) bool {
	return ls.gameServer.ExistsWaitClient(login)
}
func (ls *LoginServer) GetClientFromGS(login string) interfaces.ClientInterface {
	return ls.gameServer.GetClient(login)
}
func (ls *LoginServer) RemoveWaitingClientFromGS(login string) {
	ls.gameServer.RemoveWaitingClient(login)
}
func (ls *LoginServer) RemoveClientFromGS(login string) {
	ls.gameServer.RemoveClient(login)
}
