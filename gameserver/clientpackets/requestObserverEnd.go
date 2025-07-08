package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestObserverEnd(client interfaces.NewClientCtxInterface, data []byte) {
	pkg := serverpackets.ObservationReturn(client.GetAccount().GetCurrentChar())
	client.EncryptAndSend(pkg)
}
