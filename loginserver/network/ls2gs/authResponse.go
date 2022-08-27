package ls2gs

import (
	"l2gogameserver/config"
	"l2gogameserver/packets"
	"log"
)

func AuthResponse(data []byte) {
	reader := packets.NewReader(data)
	serverId := reader.ReadSingleByte()
	serverName := reader.ReadString()

	_ = serverName //todo вообще наверное не нужен
	sId := config.GetServerId()

	if byte(sId) != serverId {
		log.Println("serverId в конфиге и то что прилал логинсервер не совпали")
	}

}
