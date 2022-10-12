package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func RequestNpcInfo(client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()

	//
	//pkg1 := serverpackets.NpcInfo(client)
	//buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg1))

	return buffer.Bytes()
}
