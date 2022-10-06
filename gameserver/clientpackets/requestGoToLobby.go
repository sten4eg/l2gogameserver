package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestGoToLobby(client interfaces.ReciverAndSender, data []byte) {
	pkg := serverpackets.CharSelectionInfo(client)
	client.SendBuf(pkg)
}
