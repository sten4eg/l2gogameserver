package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func RequestNpcInfo(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	//
	//pkg1 := serverpackets.NpcInfo(client)
	//buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg1))

	return buffer.Bytes()
}
