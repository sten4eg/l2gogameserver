package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestManorList(client interfaces.NewClientCtxInterface, data []byte) {
	pkg := serverpackets.ExSendManorList(client)
	client.EncryptAndSend(pkg)
}
