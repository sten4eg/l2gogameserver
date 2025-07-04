package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestPrivateStoreQuitSell(client interfaces.ReciverAndSender) {
	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	character.SetPrivateStoreType(privateStoreType.NONE)
	if client.GetCurrentChar().IsSittings() {
		ChangeWaitType(client)
	}
	pkg := serverpackets.UserInfo(character)
	broadcast.BroadCastPkgToAroundPlayer(client, pkg)
}
