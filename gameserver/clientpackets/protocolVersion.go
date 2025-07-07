package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"net"
)

type protocolVersionInterface interface {
	AddLengthAndSand([]byte)
	GetRemoteAddr() net.Addr
}

func ProtocolVersion(client protocolVersionInterface, data []byte) {

	var packet = packets.NewReader(data)
	protocolVersion := packet.ReadUInt16()
	//273 - последний протокол для HF
	if protocolVersion != 273 && protocolVersion != 268 {
		logger.Info.Println(client.GetRemoteAddr(), " хотел подключиться с версией протококла:", protocolVersion)
		return
	}

	client.AddLengthAndSand(serverpackets.KeyPacket())
}
