package models

import (
	"bytes"
	"errors"
	"l2gogameserver/gameserver/crypt"
	"l2gogameserver/packets"
	"log"
	"net"
)

type Client struct {
	Socket          net.Conn
	ScrambleModulus []byte
	Buffer          *packets.Buffer
	NeedCrypt       bool
	OutKey          []int32
	InKey           []int32
	CurrentChar     *Character
	Account         *Account
	ReadBuffer      bytes.Buffer
}

func NewClient() *Client {
	buff := new(packets.Buffer)
	return &Client{
		Buffer:    buff,
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
		CurrentChar: GetNewCharacterModel(),
	}
}

func (c *Client) Send(data []byte, need bool) error {
	if need {
		data = crypt.Encrypt(data, c.OutKey)
	}
	// Calculate the packet length
	length := int16(len(data) + 2)
	// Put everything together
	buffer := packets.NewBuffer()
	buffer.WriteH(length)
	_, err := buffer.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Socket.Write(buffer.Bytes())
	if err != nil {
		return errors.New("The packet couldn't be sent.")
	}
	return nil
}

func (c *Client) SimpleSend(data []byte, needCrypt bool) {
	if needCrypt {
		data = crypt.SimpleEncrypt(data, c.OutKey)
	}

	length := int16(len(data))
	data[0], data[1] = uint8(length&0xff), uint8(length>>8)

	_, err := c.Socket.Write(data)
	c.Buffer.Reset()
	if err != nil {
		log.Println("ERROR!!!")
	}
}

func (c *Client) Receive() (opcode byte, data []byte, e error) {
	// Read the first two bytes to define the packet size
	header := make([]byte, 2)

	n, err := c.Socket.Read(header)
	//fmt.Println(n)
	if n != 2 || err != nil {
		return 0x00, nil, errors.New("12An error occured while reading the packet header.")
	}

	// Calculate the packet size
	dataSize := (int(header[0]) | int(header[1])<<8) - 2 //hack bits

	// Allocate the appropriate size for our data
	data = make([]byte, dataSize)

	// Read the encrypted part of the packet
	n, err = c.Socket.Read(data)
	if n != dataSize || err != nil {
		return 0x00, nil, errors.New("An error occured while reading the packet data.")
	}

	// Print the raw packet
	//fmt.Printf("header packet : %X\n  Raw: %X\n", header, data)
	data = crypt.Decrypt(data, &c.NeedCrypt, c.InKey)
	// Extract the op code
	opcode = data[0]
	data = data[1:]
	e = nil
	return
}
