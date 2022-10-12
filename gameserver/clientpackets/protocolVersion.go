package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func ProtocolVersion(clientI interfaces.ClientInterface, data []byte) {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return
	}

	var packet = packets.NewReader(data)
	protocolVersion := packet.ReadUInt16()
	//273 - последний протокол для HF
	if protocolVersion != 273 && protocolVersion != 268 {
		logger.Info.Println(client.GetRemoteAddr(), " хотел подключиться с версией протококла:", protocolVersion)
		return
	}

	clientI.AddLengthAndSand(serverpackets.KeyPacket(client))
}
