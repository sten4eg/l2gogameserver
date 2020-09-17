package clientpackets

import (
	"l2gogameserver/packets"
)

func NewprotocolVersion(data []byte) bool {

	var packet = packets.NewReader(data)
	protocolVersion := packet.ReadUInt16()

	_ = protocolVersion
	return true // todo
}
