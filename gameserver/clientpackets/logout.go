package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func Logout(client interfaces.ReciverAndSender, data []byte) {
	pkg := serverpackets.LogoutToClient(data, client)
	client.EncryptAndSend(pkg)
}
