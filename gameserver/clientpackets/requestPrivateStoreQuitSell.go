package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestPrivateStoreQuitSell(client interfaces.ReciverAndSender) {
	activeChar := client.GetCurrentChar()
	if activeChar == nil {
		return
	}

	activeChar.SetPrivateStoreType(privateStoreType.NONE)
	if client.GetCurrentChar().IsSittings() {
		ChangeWaitType(client)
	}
	pkg := serverpackets.UserInfo(activeChar)
	broadcast.BroadCastPkgToAroundPlayer(client, pkg)
}
