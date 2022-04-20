package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
)

func ProtocolVersion(clientI interfaces.ReciverAndSender, data []byte) {
	client, ok := clientI.(*models.Client)
	if !ok {
		return
	}

	var packet = packets.NewReader(data)
	protocolVersion := packet.ReadUInt16()
	if protocolVersion != 273 && protocolVersion != 268 {
		log.Println(client.Socket.RemoteAddr(), " хотел подключиться с версией протококла:", protocolVersion)
		return
	}

	client.AddLengthAndSand(serverpackets.KeyPacket(client))
}
