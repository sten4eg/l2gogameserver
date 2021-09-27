package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func ChangeMoveType(client *models.Client, targetObjectId int32) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg1 := serverpackets.ChangeMoveType(client, targetObjectId)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg1))

	return buffer.Bytes()

}
