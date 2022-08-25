package loginserver

import (
	"fmt"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/crypt"
	"l2gogameserver/gameserver/crypt/blowfish"
	"l2gogameserver/loginserver/network/gs2ls"
	"l2gogameserver/loginserver/network/ls2gs"
	"l2gogameserver/packets"
	"log"
	"math/rand"
	"net"
	"strconv"
)

type LoginServer struct {
	conn     *net.TCPConn
	blowfish *blowfish.Cipher
}

var loginServerInstance *LoginServer

var initBlowfishKey = []byte{95, 59, 118, 46, 93, 48, 53, 45, 51, 49, 33, 124, 43, 45, 37, 120, 84, 33, 94, 91, 36, 0}

func HandlerInit() error {
	loginServerInstance = new(LoginServer)

	port := config.GetLoginServerPort()
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	addr := new(net.TCPAddr)
	addr.Port = intPort
	addr.IP = net.IP{127, 0, 0, 1}

	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		return err
	}
	loginServerInstance.conn = conn

	go loginServerInstance.Run()

	return nil
}

func (ls *LoginServer) Run() {
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

func (ls *LoginServer) HandlePacket(data []byte) {
	opCode := data[0]
	data = data[1:]
	fmt.Println(opCode)

	switch opCode {
	default:
		fmt.Printf("неопознаный опкод от логинсервера: %v\n", opCode)
	case 0x00:
		pubKey := ls2gs.InitLs(data)
		bfk := generateNewBlowFish()
		buf := gs2ls.BlowFishKey(pubKey, bfk)

		ls.Send(buf)
		ls.setBlowFish(bfk)
		buf = gs2ls.AuthRequest()
		ls.Send(buf)
	case 0x02:
		ls2gs.AuthResponse(data)
		buf := gs2ls.ServerStatus()
		ls.Send(buf)
	}
}
func generateNewBlowFish() []byte {
	bfk := make([]byte, 40)
	_, _ = rand.Read(bfk)
	bfk[0] = 119 // главное чтобы не 0
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
