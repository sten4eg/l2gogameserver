package models

import (
	"errors"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/crypt"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/packets"
	"net"
	"sync"
)

type ClientCtx struct {
	m               sync.RWMutex
	conn            *net.TCPConn
	ScrambleModulus []byte
	// NeedCrypt - флаг, при создании клиента false
	// первый пакет пришедший от клиента не нужно расшифровывать, после 1 пакета NeedCrypt = true
	// костыль
	NeedCrypt   bool
	OutKey      []int32
	InKey       []int32
	CurrentChar *Character
	Account     *Account
	state       clientStates.State
	sessionKey  SessionKey
}
type SessionKey struct {
	PlayOk1  uint32
	PlayOk2  uint32
	LoginOk1 uint32
	LoginOk2 uint32
}

func NewClient() *ClientCtx {
	c := &ClientCtx{
		NeedCrypt: false,
		OutKey: []int32{
			0x6b,
			0x60,
			0xcb,
			0x5b,
			0x82,
			0xce,
			0x90,
			0xb1,
			0xc8,
			0x27,
			0x93,
			0x01,
			0xa1,
			0x6c,
			0x31,
			0x97,
		},
		InKey: []int32{
			0x6b,
			0x60,
			0xcb,
			0x5b,
			0x82,
			0xce,
			0x90,
			0xb1,
			0xc8,
			0x27,
			0x93,
			0x01,
			0xa1,
			0x6c,
			0x31,
			0x97,
		},
		Account:     new(Account),
		CurrentChar: nil,
		state:       clientStates.Connected,
	}

	return c
}

// AddLengthAndSand добавляет 2 байта длинны и отправляет клиенту
func (c *ClientCtx) AddLengthAndSand(data []byte) {
	// вычисление длинны пакета, 2 первых байта - размер пакета
	length := int16(len(data) + 2)

	s, f := byte(length>>8), byte(length&0xff)

	data = append([]byte{f, s}, data...)
	c.Send(data)
}

func (c *ClientCtx) EncryptAndSend(data []byte) error {
	data = crypt.Encrypt(data, c.OutKey)
	// вычисление длинны пакета, 2 первых байта - размер пакета
	length := int16(len(data) + 2)

	s, f := byte(length>>8), byte(length&0xff)

	data = append([]byte{f, s}, data...)

	err := c.sendDataToSocket(data)
	if err != nil {
		//logger.Error.Panicln("Пакет не отправлен, ошибка: " + err.Error())
		return err
	}
	return nil
}
func (c *ClientCtx) SendBuf(buffer *packets.Buffer) error {
	if buffer == nil {
		return nil //todo мб ошибку кинуть?
	}

	data := buffer.Bytes()
	defer packets.Put(buffer)

	data = crypt.Encrypt(data, c.OutKey)
	// Вычисление длинны пакета
	length := uint16(len(data) + 2)

	toSend := packets.Get()
	toSend.WriteHU(length)
	toSend.WriteSlice(data) //TODO очень много выделяет
	defer packets.Put(toSend)

	err := c.sendDataToSocket(toSend.Bytes())
	if err != nil {
		logger.Error.Panicln("Пакет не отправлен, ошибка: " + err.Error())
	}

	return nil
}
func (c *ClientCtx) Send(d []byte) {
	if len(d) == 0 {
		logger.Info.Println("Пакет пуст")
		return
	}
	err := c.sendDataToSocket(d)
	if err != nil {
		logger.Error.Panicln("Пакет не отправлен, ошибка: " + err.Error())
	}
}

func (c *ClientCtx) SendSysMsg(num interface{}, options ...string) error {
	smsg := num.(sysmsg.SysMsg)

	return c.EncryptAndSend(sysmsg.SystemMessage(smsg))
}

func (c *ClientCtx) CryptAndReturnPackageReadyToShip(data []byte) []byte {
	data = crypt.Encrypt(data, c.OutKey)
	// вычисление длинны пакета, 2 первых байта - размер пакета
	length := int16(len(data) + 2)

	s, f := byte(length>>8), byte(length&0xff)

	data = append([]byte{f, s}, data...)

	return data
}

func (c *ClientCtx) Receive() (opcode byte, data []byte, e error) {
	// чтение первых 2 байта для определения размера всего пакета
	header := make([]byte, 2)

	n, err := c.conn.Read(header)

	if err != nil {
		return 0, nil, err
	}

	if n != 2 {
		return 0, nil, errors.New("байтов длинны пакета должно быть 2")
	}

	// длинна пакета
	dataSize := (int(header[0]) | int(header[1])<<8) - 2

	// аллокация требуемого массива байт для входящего пакета
	data = make([]byte, dataSize)

	n, err = c.conn.Read(data)
	if err != nil {
		return 0, nil, err
	}

	if n != dataSize {
		return 0, nil, errors.New("длинна прочитанного пакета не соответствует требуемому размеру")
	}

	// Если это первый пакет от юзера то его не расшифровываем
	// todo можно ли приудмать что нибудь лучше?
	if c.NeedCrypt {
		data = crypt.Decrypt(data, c.InKey)
	} else {
		c.NeedCrypt = true
	}

	// первый байт opcode, остальные полезная нагрузка
	opcode = data[0]
	data = data[1:]
	e = nil
	return
}

func (c *ClientCtx) sendDataToSocket(data []byte) error {
	c.m.Lock()
	_, err := c.conn.Write(data)
	c.m.Unlock()
	return err
}

func (c *ClientCtx) GetCurrentChar() interfaces.CharacterI {
	if c.CurrentChar == nil {
		return nil
	}
	return c.CurrentChar
}

func (c *ClientCtx) SetConn(conn *net.TCPConn) {
	c.conn = conn
}

func (c *ClientCtx) GetConn() *net.TCPConn {
	return c.conn
}

func (c *ClientCtx) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *ClientCtx) SetLogin(login string) {
	c.Account.Login = login
}

func (c *ClientCtx) RemoveCurrentChar() {
	c.CurrentChar = nil
}

func (c *ClientCtx) SetState(state clientStates.State) {
	c.state = state
}

func (c *ClientCtx) GetState() clientStates.State {
	return c.state
}

func (c *ClientCtx) SetSessionKey(playOk1, playOk2, loginOk1, loginOk2 uint32) {
	c.sessionKey.PlayOk1 = playOk1
	c.sessionKey.PlayOk2 = playOk2
	c.sessionKey.LoginOk1 = loginOk1
	c.sessionKey.LoginOk2 = loginOk2
}

func (c *ClientCtx) GetSessionKey() (playOk1, playOk2, loginOk1, loginOk2 uint32) {
	return c.sessionKey.PlayOk1, c.sessionKey.PlayOk2, c.sessionKey.LoginOk1, c.sessionKey.LoginOk2
}

func (c *ClientCtx) GetAccountLogin() string {
	return c.Account.Login
}

func (c *ClientCtx) CloseConnection() {
	c.conn.Close()
}
