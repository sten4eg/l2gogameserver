package clientpackets

import (
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/serverpackets"
)

type logoutInterface interface {
	EncryptAndSend(data []byte)
}

func Logout(data []byte, client logoutInterface, state clientStates.State) {
	var pkg []byte
	if state == clientStates.InGame {
		pkg = serverpackets.LogoutWithInGameState()
	} else {
		pkg = serverpackets.LogoutWithAuthedState()
	}
	client.EncryptAndSend(pkg)
}
