package ls2gs

import (
	"l2gogameserver/config"
	"l2gogameserver/loginserver/network/gs2ls"
	"l2gogameserver/packets"
	"log"
)

type loginServerInterfaceAuthResponse interface {
	Send(buffer *packets.Buffer)
}

func AuthResponse(data []byte, ls loginServerInterfaceAuthResponse) {
	reader := packets.NewReader(data)
	serverId := reader.ReadSingleByte()
	serverName := reader.ReadString()

	_ = serverName //todo вообще наверное не нужен
	sId := config.GetServerId()

	if byte(sId) != serverId {
		log.Println("serverId в конфиге и то что прилал логинсервер не совпали")
	}

	buf := gs2ls.ServerStatus()
	ls.Send(buf)

}
