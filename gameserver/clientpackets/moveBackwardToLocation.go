package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/world"
	"l2gogameserver/packets"
)

func MoveBackwardToLocation(client interfaces.NewClientCtxInterface, data []byte) world.BackwardToLocation {

	var location world.BackwardToLocation
	var packet = packets.NewReader(data)

	location.TargetX = packet.ReadInt32()
	location.TargetY = packet.ReadInt32()
	location.TargetZ = packet.ReadInt32()
	location.OriginX = packet.ReadInt32()
	location.OriginY = packet.ReadInt32()
	location.OriginZ = packet.ReadInt32()

	return location

}

func MoveToLocation(client interfaces.NewClientCtxInterface, targetX, targetY, targetZ int32) *world.BackwardToLocation {

	x, y, z := client.GetAccount().GetCurrentChar().GetXYZ()
	location := world.BackwardToLocation{
		TargetX: targetX,
		TargetY: targetY,
		TargetZ: targetZ,
		OriginX: x,
		OriginY: y,
		OriginZ: z,
	}
	return &location
}
