package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func MoveBackwardToLocation(client interfaces.ReciverAndSender, data []byte) models.BackwardToLocation {

	var location models.BackwardToLocation
	var packet = packets.NewReader(data)

	location.TargetX = packet.ReadInt32()
	location.TargetY = packet.ReadInt32()
	location.TargetZ = packet.ReadInt32()
	location.OriginX = packet.ReadInt32()
	location.OriginY = packet.ReadInt32()
	location.OriginZ = packet.ReadInt32()

	return location

}

func MoveToLocation(clientI interfaces.ReciverAndSender, targetX, targetY, targetZ int32) *models.BackwardToLocation {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return nil
	}
	x, y, z := client.CurrentChar.GetXYZ()
	location := models.BackwardToLocation{
		TargetX: targetX,
		TargetY: targetY,
		TargetZ: targetZ,
		OriginX: x,
		OriginY: y,
		OriginZ: z,
	}
	return &location
}
