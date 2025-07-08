package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestPrivateStoreQuitSell(client interfaces.NewClientCtxInterface) {
	character := client.GetAccount().GetCurrentChar()
	if character == nil {
		return
	}

	character.SetPrivateStoreType(privateStoreType.NONE)
	if character.IsSittings() {
		ChangeWaitType(client)
	}
	pkg := serverpackets.UserInfo(character)
	broadcast.BroadCastPkgToAroundPlayer(client, pkg)
}
