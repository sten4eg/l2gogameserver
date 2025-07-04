package clientpackets

import (
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/serverpackets"
)

type logoutInterface interface {
	EncryptAndSend(data []byte) error
	GetAccountLogin() string
}

func Logout(client logoutInterface, state clientStates.State, gs gameServerInterface) {
	var pkg []byte
	if state == clientStates.InGame {
		pkg = serverpackets.LogoutWithInGameState()
	} else {
		pkg = serverpackets.LogoutWithAuthedState()
	}

	client.EncryptAndSend(pkg)
	login := client.GetAccountLogin()

	gs.SendLogout(login)
	gs.RemoveClient(login)

}
