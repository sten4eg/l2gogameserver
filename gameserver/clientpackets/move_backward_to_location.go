package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func MoveBackwardToLocation(client *models.Client, data []byte) []byte {

	var location models.BackwardToLocation
	var packet = packets.NewReader(data)

	location.TargetX = packet.ReadInt32()
	location.TargetY = packet.ReadInt32()
	location.TargetZ = packet.ReadInt32()
	location.OriginX = packet.ReadInt32()
	location.OriginY = packet.ReadInt32()
	location.OriginZ = packet.ReadInt32()

	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.MoveToLocation(&location, client)

	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
