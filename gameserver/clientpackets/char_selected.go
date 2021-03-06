package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewCharSelected(data []byte, client *models.Client) {

	var buffer = packets.NewReader(data)
	charSlot := buffer.ReadInt32()
	_ = buffer.ReadUInt16() // unused, remove ?
	_ = buffer.ReadInt32()  // unused, remove ?
	_ = buffer.ReadInt32()  // unused, remove ?
	_ = buffer.ReadInt32()  // unused, remove ?

	client.Account.CharSlot = charSlot
	serverpackets.NewSSQInfo(client)
}
