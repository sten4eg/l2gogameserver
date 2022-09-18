package clientpackets

import (
	"l2gogameserver/gameserver/serverpackets"
)

type logoutInterface interface {
	EncryptAndSend(data []byte)
}

func Logout(data []byte, client logoutInterface) {
	pkg := serverpackets.LogoutToClient(data)
	client.EncryptAndSend(pkg)
}
