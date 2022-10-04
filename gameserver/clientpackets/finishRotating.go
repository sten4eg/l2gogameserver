package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func FinishRotating(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	degree := reader.ReadInt32()
	_ = reader.ReadInt32() // unknown

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	if false { //TODO if (character.isInAirShip() && character.getAirShip().isCaptain(character))
		// character.getAirShip().setHeading(_degree)
		//pkg := serverpackets.StopRotation(0, degree, 0) // TODO character.getAirShip().getObjectId()
		// character.getAirShip().broadcastPacket(sr)
	} else {
		pkg := serverpackets.StopRotation(character.GetObjectId(), degree, 0)
		broadcast.BroadCastBufferToAroundPlayers(client, pkg)
	}

}
