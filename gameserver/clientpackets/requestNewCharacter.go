package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestNewCharacter(client interfaces.NewClientCtxInterface) {
	client.SendBuf(serverpackets.CharacterSuccess())
}
