package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestGoToLobby(client interfaces.ReciverAndSender) {
	client.SendBuf(serverpackets.CharSelectionInfo(client))
}
