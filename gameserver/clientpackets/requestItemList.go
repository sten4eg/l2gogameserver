package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestItemList(data []byte, client interfaces.NewClientCtxInterface) {
	pkg := serverpackets.ItemList(client.GetAccount().GetCurrentChar())
	client.EncryptAndSend(pkg)
}
