package gameserver

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type GameServer struct {
	clientsListener     net.Listener
	clients             []*Client
	Socket net.Conn
}
type Client struct {
	Socket          net.Conn
	ScrambleModulus []byte
}
func NewClient() *Client {

	return &Client{}
}
func New() *GameServer {
	return &GameServer{}
}
func (g *GameServer) Init()  {
	var err error
	g.clientsListener, err = net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal("Failed to connect to port 7777:", err.Error())
	} else {
		fmt.Println("Login server is listening on port 7777")
	}
}

func (g *GameServer) Start()  {
	defer g.clientsListener.Close()

	done := make(chan bool)

	go func() {
		for {
			var err error
			client := NewClient()
			g.Socket, err = g.clientsListener.Accept()
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

func (g *GameServer) handleClientPackets(client *Client)  {
	for {
		opcode, _, err := g.Receive()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Closing the connection...")
			break
		}

		switch opcode {
		case 00:
			fmt.Println("A game server sent a request to register")
		default:
			fmt.Println("Can't recognize the packet sent by the gameserver")
		}
	}
}

func (g *GameServer) Receive() (opcode byte, data []byte, e error) {
	// Read the first two bytes to define the packet size
	header := make([]byte, 2)
	n, err := g.Socket.Read(header)

	if n != 2 || err != nil {
		return 0x00, nil, errors.New("12An error occured while reading the packet header.")
	}

	// Calculate the packet size
	size := 0
	size = size + int(header[0])
	size = size + int(header[1])*256

	// Allocate the appropriate size for our data (size - 2 bytes used for the length
	data = make([]byte, size-2)

	// Read the encrypted part of the packet
	n, err = g.Socket.Read(data)

	if n != size-2 || err != nil {
		return 0x00, nil, errors.New("An error occured while reading the packet data.")
	}

	// Print the raw packet
	fmt.Printf("header packet : %X\n  Raw: %X\n", header, data)

	// Extract the op code
	opcode = data[0]
	data = data[1:]
	e = nil
	return
}