package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func ChangeWaitType(client interfaces.NewClientCtxInterface) {
	pkg := serverpackets.ChangeWaitType(client.GetAccount().GetCurrentChar())
	broadcast.BroadCastPkgToAroundPlayer(client, pkg)
}
