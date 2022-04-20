package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestObserverEnd(client interfaces.ReciverAndSender, data []byte) {
	pkg := serverpackets.ObservationReturn(client)
	client.EncryptAndSend(pkg)
}
