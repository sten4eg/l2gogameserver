package clientpackets

import (
	"l2gogameserver/packets"
	"log"
)

func NewAuthLogin(data []byte) {

	var packet = packets.NewReader(data)

	login := packet.ReadString()
	log.Println(login)
	_ = login

}
