package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewRequestTargetCanceld(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)
	unselect := packet.ReadUInt16()
	_ = unselect
	serverpackets.NewTargetUnselected(client)
	client.SimpleSend(client.Buffer.Bytes(), true)
}
