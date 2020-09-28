package clientpackets

import (
	"l2gogameserver/gameserver/models"
)

type Location struct {
	TargetX int32
	TargetY int32
	TargetZ int32
	OriginX int32
	OriginY int32
	OriginZ int32
}

func NewMoveBackwardToLocation(client *models.Client, data []byte) *Location {

	var location Location

	location.TargetX = client.Reader.RreadInt32()
	location.TargetY = client.Reader.RreadInt32()
	location.TargetZ = client.Reader.RreadInt32()
	location.OriginX = client.Reader.RreadInt32()
	location.OriginY = client.Reader.RreadInt32()
	location.OriginZ = client.Reader.RreadInt32()

	//location.TargetX = packet.ReadInt32()
	//location.TargetY = packet.ReadInt32()
	//location.TargetZ = packet.ReadInt32()
	//location.OriginX = packet.ReadInt32()
	//location.OriginY = packet.ReadInt32()
	//location.OriginZ = packet.ReadInt32()

	return &location
}
