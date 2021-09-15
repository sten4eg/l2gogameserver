package models

import (
	"errors"
	"l2gogameserver/gameserver/crypt"
	"l2gogameserver/packets"
	"log"
	"net"
	"sync"
)

type Client struct {
	m               sync.RWMutex
	Socket          net.Conn
	ScrambleModulus []byte
	Buffer          *packets.Buffer
	// NeedCrypt - флаг, при создании клиента false
	// указывает первый пакет пришедший от клиента не нужно расшифровывать
	// костыль
	NeedCrypt   bool
	OutKey      []int32
	InKey       []int32
	CurrentChar *Character
	Account     *Account
	// ToSendBuffer буффер полностью готовых к отправке пакета/пакетов
	ToSendBuffer *packets.Buffer
}

func NewClient() *Client {
	buff := new(packets.Buffer)
	toS := new(packets.Buffer)
	c := &Client{
		Buffer:       buff,
		ToSendBuffer: toS,
		NeedCrypt:    false,
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
		CurrentChar: &Character{},
	}

	return c
}

// Send отправка массив data персонажу
// need - флаг, указывает надо ли шифровать данные
func (c *Client) Send(data []byte, need bool) error {
	if need {
		data = crypt.Encrypt(data, c.OutKey)
	}
	// вычисление длинны пакета, 2 первых байта - размер пакета
	length := int16(len(data) + 2)

	s, f := byte(length>>8), byte(length&0xff)

	data = append([]byte{f, s}, data...)

	err := c.sendDataToSocket(data)
	if err != nil {
		return errors.New("Пакет не отправлен, ошибка: " + err.Error())
	}

	return nil
}

// SaveAndCryptDataInBufferToSend подготавливает данные из
// c.Buffer ---> c.ToSendBuffer
func (c *Client) SaveAndCryptDataInBufferToSend(needCrypt bool) {
	c.Buffer.Mu.Lock()
	data := c.Buffer.Bytes()
	c.Buffer.Reset()
	c.Buffer.Mu.Unlock()

	if len(data) == 0 {
		return
	}
	log.Println("Пакет с опкодом : ", data[0], " подготовлен к отправке")
	// добавление первых двух байт для длинны пакета
	data = append([]byte{0, 0}, data...)

	if needCrypt {
		data = crypt.SimpleEncrypt(data, c.OutKey)
	}

	length := int16(len(data))
	data[0], data[1] = uint8(length&0xff), uint8(length>>8)

	c.ToSendBuffer.Mu.Lock()
	c.ToSendBuffer.WriteSlice(data)
	c.ToSendBuffer.Mu.Unlock()
}

// SentToSend отправляет пользователю данные из c.ToSendBuffer
func (c *Client) SentToSend() {
	c.ToSendBuffer.Mu.Lock()
	data := c.ToSendBuffer.Bytes()
	c.ToSendBuffer.Reset()
	c.ToSendBuffer.Mu.Unlock()

	if len(data) == 0 {
		return
	}
	err := c.sendDataToSocket(data)

	if err != nil {
		panic(err)
	}
}

func (c *Client) Receive() (opcode byte, data []byte, e error) {
	// чтение первых 2 байта для определения размера всего пакета
	header := make([]byte, 2)

	n, err := c.Socket.Read(header)

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

	n, err = c.Socket.Read(data)
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

func (c *Client) sendDataToSocket(data []byte) error {
	c.m.Lock()
	_, err := c.Socket.Write(data)
	c.m.Unlock()
	return err
}
