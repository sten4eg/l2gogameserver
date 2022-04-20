package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestItemList(client interfaces.ReciverAndSender, data []byte) {
	pkg := serverpackets.ItemList(client)
	client.EncryptAndSend(pkg)
}
