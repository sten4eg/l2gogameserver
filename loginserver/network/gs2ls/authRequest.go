package gs2ls

import (
	"l2gogameserver/config"
	"l2gogameserver/packets"
)

const version byte = 14

func AuthRequest() *packets.Buffer {
	serverId := config.GetServerId()
	hexId := config.GetHexId()
	acceptAlternative := config.GetAcceptAlternateId()
	reserveHost := config.GetReserveHostOnLogin()
	port := config.GetPort()
	maxPlayer := config.GetMaxPlayer()
	subNet := config.GetSubNets()
	hosts := config.GetHosts() //TODO в будующем можно удалить , и удалить принятие этих данных на логинсервере, все эти хосты есть в subNet

	buf := packets.Get()
	buf.WriteSingleByte(0x01)
	buf.WriteSingleByte(version)
	buf.WriteSingleByte(byte(serverId))
	if acceptAlternative {
		buf.WriteSingleByte(0x01)
	} else {
		buf.WriteSingleByte(0x00)
	}

	if reserveHost {
		buf.WriteSingleByte(0x01)
	} else {
		buf.WriteSingleByte(0x00)
	}

	buf.WriteH(int16(port))
	buf.WriteD(int32(maxPlayer))
	buf.WriteD(int32(len(hexId)))
	buf.WriteSlice(hexId)
	buf.WriteD(int32(len(subNet)))

	for i := range subNet {
		buf.WriteS(subNet[i])
		buf.WriteS(hosts[i])
	}
	return buf
}
