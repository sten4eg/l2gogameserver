package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestShowMiniMap(client interfaces.ReciverAndSender, data []byte) {
	pkg := serverpackets.ShowMiniMap(client)
	client.EncryptAndSend(pkg)
}
