package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestTargetCancel(data []byte, client interfaces.ReciverAndSender) {

	var packet = packets.NewReader(data)
	unselect := packet.ReadUInt16()
	_ = unselect

	pkg := serverpackets.TargetUnselected(client.GetCurrentChar())
	client.EncryptAndSend(pkg)
}
