package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewSay(data []byte) *models.Say {
	var packet = packets.NewReader(data)
	var say models.Say
	text := packet.ReadString()

	say.Text = text
	say.Type = packet.ReadInt32()
	return &say
}
