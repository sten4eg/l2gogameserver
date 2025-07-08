package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

type gameServerInterface interface {
	AddWaitClient(string, uint32, uint32, uint32, uint32)
	AddClient(string, interfaces.NewClientCtxInterface) bool
	RemoveClient(string)
	SendLogout(string)
}
type authLoginClientInterface interface {
	SetLogin(string)
	GetCurrentChar() interfaces.CharacterI
	SetSessionKey(uint32, uint32, uint32, uint32)
}

func AuthLogin(data []byte, client interfaces.NewClientCtxInterface, gs gameServerInterface) {
	var packet = packets.NewReader(data)

	login := packet.ReadString()
	client.GetAccount().SetLogin(login)
	playKey1 := packet.ReadUInt32()
	playKey2 := packet.ReadUInt32()
	loginKey1 := packet.ReadUInt32()
	loginKey2 := packet.ReadUInt32()
	//TODO проверить что они приходят в правильном порядке

	if client.GetAccount().GetCurrentChar() == nil {
		if gs.AddClient(login, client) {
			client.SetSessionKey(playKey1, playKey2, loginKey1, loginKey2)
			gs.AddWaitClient(login, playKey1, playKey2, loginKey1, loginKey2)
		} else {
			//TODO client.CLOSE()
		}

	}

}
