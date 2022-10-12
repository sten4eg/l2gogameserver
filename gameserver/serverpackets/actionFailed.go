package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
)

func ActionFailed(client interfaces.ReciverAndSender) []byte {
	return []byte{0x1f}
}
