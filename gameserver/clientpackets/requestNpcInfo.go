package clientpackets

import (
	"l2gogameserver/packets"
)

func RequestNpcInfo() []byte {
	buffer := packets.Get()

	//
	//pkg1 := serverpackets.NpcInfo(client)
	//buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg1))

	return buffer.Bytes()
}
