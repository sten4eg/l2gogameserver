package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewprotocolVersion(data []byte, client *models.Client) {

	var packet = packets.NewReader(data)
	protocolVersion := packet.ReadUInt16() //todo check !=273
	_ = protocolVersion
	serverpackets.NewKeyPacket(client)
}
