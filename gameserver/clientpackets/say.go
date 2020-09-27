package clientpackets

import (
	"bytes"
	"l2gogameserver/packets"
)

type Say struct {
	Text string
	Type int32
}

func NewSay(data []byte) *Say {
	var packet = packets.NewReader(data)
	var say Say
	trimText := bytes.Trim([]byte(packet.ReadString()), string(rune(0)))
	say.Text = string(trimText)
	say.Type = packet.ReadInt32()
	return &say
}
