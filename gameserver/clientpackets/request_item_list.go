package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestItemList(data []byte, client interfaces.CharacterI) {
	pkg := serverpackets.ItemList(client)
	client.EncryptAndSend(pkg)
}
