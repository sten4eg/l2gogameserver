package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewValidationPosition(data []byte, client *models.Character) {
	var packet = packets.NewReader(data)
	var l models.Coordinates

	l.X = packet.ReadInt32()
	l.Y = packet.ReadInt32()
	l.Z = packet.ReadInt32()
	_ = packet.ReadInt32() // heading
	_ = packet.ReadInt32() //data

	clientX, clientY, clientZ := client.GetXYZ()
	dx := l.X - clientX
	dy := l.Y - clientY
	_ = l.Z - clientZ
	diffSq := (dx * dx) + (dy * dy)

	if diffSq < 360000 {
		client.SetXYZ(clientX, clientY, l.Z)
	} else {
		client.SetXYZ(l.X, l.Y, l.Z)
	}
}
