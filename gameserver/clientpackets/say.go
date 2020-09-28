package clientpackets

import (
	"l2gogameserver/packets"
)

type Say struct {
	Text string
	Type int32
}

func NewSay(data []byte) *Say {
	var packet = packets.NewReader(data)
	var say Say
	text := packet.ReadString()

	say.Text = text
	say.Type = packet.ReadInt32()
	return &say
}
