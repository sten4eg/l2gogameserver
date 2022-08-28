package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

type gameServerInterface interface {
	AddWaitClient(string, interfaces.ClientInterface)
	AddClient(string, interfaces.ClientInterface) bool
}

func AuthLogin(data []byte, client interfaces.ClientInterface, gs gameServerInterface) {
	var packet = packets.NewReader(data)

	login := packet.ReadString()
	client.SetLogin(login)
	playKey1 := packet.ReadUInt32()
	playKey2 := packet.ReadUInt32()
	loginKey1 := packet.ReadUInt32()
	loginKey2 := packet.ReadUInt32()
	//TODO проверить что они приходят в правильном порядке

	_, _, _, _ = playKey1, playKey2, loginKey1, loginKey2

	if client.GetCurrentChar() == nil {
		if gs.AddClient(login, client) {
			client.SetSessionKey(playKey1, playKey2, loginKey1, loginKey2)
			gs.AddWaitClient(login, client)
		} else {
			//TODO client.CLOSE()
		}

		pkg := serverpackets.CharSelectionInfo(client)
		client.EncryptAndSend(pkg)
	}

}
