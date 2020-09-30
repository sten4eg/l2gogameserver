package clientpackets

import (
	"l2gogameserver/packets"
)

func NewAuthLogin(data []byte) string {

	var packet = packets.NewReader(data)

	login := packet.ReadString()

	return login
}
