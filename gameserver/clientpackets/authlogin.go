package clientpackets

import (
	"l2gogameserver/packets"
)

func NewAuthLogin(data []byte) {

	var packet = packets.NewReader(data)

	login := packet.ReadString()

	_ = login

}
