package ls2gs

import (
	"database/sql"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/loginserver/network/gs2ls"
	"l2gogameserver/packets"
)

type loginServerInterface interface {
	ExistsWaitClientOnGameServer(login string) bool
	Send(buffer *packets.Buffer)
	GetClientFromGS(login string) interfaces.ClientInterface
	RemoveWaitingClientFromGS(login string)
	RemoveClientFromGS(login string)

	SendLogoutFromGS(login string)
}

func PlayerAuthResponse(data []byte, ls loginServerInterface, db *sql.DB) {
	reader := packets.NewReader(data)

	account := reader.ReadString()
	authed := reader.ReadSingleByte() != 0

	if ls.ExistsWaitClientOnGameServer(account) {
		client := ls.GetClientFromGS(account)
		if authed {
			ls.Send(gs2ls.PlayerInGame([]string{account}))
			client.SetState(clientStates.Authed)
			//todo чекнуть sessionKey
			pkg := serverpackets.CharSelectionInfo(client, db)
			client.SendBuf(pkg)
		} else {
			//log
			client.EncryptAndSend(serverpackets.LoginFail(serverpackets.SystemErrorLoginLater))
			ls.RemoveClientFromGS(account)
		}

		ls.RemoveWaitingClientFromGS(account)
	}

}
