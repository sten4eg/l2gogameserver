package ls2gs

import (
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func KickPlayer(data []byte, ls loginServerInterface) {
	reader := packets.NewReader(data)

	account := reader.ReadString()
	client := ls.GetClientFromGS(account)

	if client != nil {
		client.EncryptAndSend(serverpackets.ServerClose())
		client.EncryptAndSend(sysmsg.SystemMessage(sysmsg.AnotherLoginWithAccount))

		ls.SendLogoutFromGS(account)
		ls.RemoveClientFromGS(account)
		client.CloseConnection()
	}
}