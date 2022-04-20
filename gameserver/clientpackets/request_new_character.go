package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestNewCharacter(client interfaces.ReciverAndSender, data []byte) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.CharacterSuccess(client)

	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()

}
