package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
)

func RequestPrivateStoreQuitBuy(client interfaces.ReciverAndSender) {
	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	character.SetPrivateStoreType(privateStoreType.NONE)
	if character.IsSittings() {
		ChangeWaitType(client)
	}
	broadcast.BroadcastUserInfo(client)
}
