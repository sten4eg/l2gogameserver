package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func ChangeWaitType(client interfaces.ReciverAndSender) {
	pkg := serverpackets.ChangeWaitType(client.GetCurrentChar())
	broadcast.BroadCastPkgToAroundPlayer(client, pkg)
}
