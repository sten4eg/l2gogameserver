package clientpackets

import (
	"l2gogameserver/packets"
)

type Location struct {
	TargetX int32
	TargetY int32
	TargetZ int32
	OriginX int32
	OriginY int32
	OriginZ int32
}

func NewMoveBackwardToLocation(data []byte) *Location {
	var packet = packets.NewReader(data)
	var location Location
	location.TargetX = packet.ReadInt32()
	location.TargetY = packet.ReadInt32()
	location.TargetZ = packet.ReadInt32()
	location.OriginX = packet.ReadInt32()
	location.OriginY = packet.ReadInt32()
	location.OriginZ = packet.ReadInt32()

	return &location
}
