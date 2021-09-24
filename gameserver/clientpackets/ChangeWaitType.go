package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func ChangeWaitType(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.ChangeWaitType(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
