package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func ChangeWaitType(client interfaces.ReciverAndSender) {
	pkg := serverpackets.ChangeWaitType(client)
	client.EncryptAndSend(pkg)
}
