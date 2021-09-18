package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/chat"
	"l2gogameserver/packets"
	"strings"
)

func Say(client *models.Client, data []byte) models.Say {
	var packet = packets.NewReader(data)
	var say models.Say

	say.Text = packet.ReadString()
	say.Type = packet.ReadInt32()

	buffer := packets.Get()
	if strings.HasPrefix(say.Text, ".") {
		say.Type = chat.SpecialCommand
		say.Text = "tok"
	}
	defer packets.Put(buffer)

	switch say.Type {
	case chat.All:
		return say
	case chat.Tell:
		say.To = packet.ReadString()
		return say
	case chat.Shout:
		return say
	}

	return say
}
