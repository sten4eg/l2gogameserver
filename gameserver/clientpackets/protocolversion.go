package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
)

func ProtocolVersion(data []byte, clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.Client)
	if !ok {
		return []byte{}
	}

	var packet = packets.NewReader(data)
	protocolVersion := packet.ReadUInt16()
	if protocolVersion != 273 && protocolVersion != 268 {
		log.Println(client.Socket.RemoteAddr(), " хотел подключиться с версией протококла:", protocolVersion)
		return []byte{}
	}

	buffer := packets.Get()

	pkg1 := serverpackets.KeyPacket(client)
	buffer.WriteSlice(client.ReturnPackageReadyToShip(pkg1))

	defer packets.Put(buffer)
	return buffer.Bytes()

}
