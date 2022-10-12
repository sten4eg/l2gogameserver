package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestNewCharacter(client interfaces.ReciverAndSender) {
	client.SendBuf(serverpackets.CharacterSuccess())
}
