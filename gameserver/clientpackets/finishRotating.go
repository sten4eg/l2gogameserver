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

	activeChar := client.GetCurrentChar()
	if activeChar == nil {
		return
	}

	if false { //TODO if (activeChar.isInAirShip() && activeChar.getAirShip().isCaptain(activeChar))
		// activeChar.getAirShip().setHeading(_degree)
		//pkg := serverpackets.StopRotation(0, degree, 0) // TODO activeChar.getAirShip().getObjectId()
		// activeChar.getAirShip().broadcastPacket(sr)
	} else {
		pkg := serverpackets.StopRotation(activeChar.GetObjectId(), degree, 0)
		broadcast.BroadCastBufferToAroundPlayers(client, pkg)
	}

}
