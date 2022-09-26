package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestRestart(data []byte, clientI interfaces.ReciverAndSender) {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return
	}

	client.SaveUser()
	//todo need save in db

	_ = data
	buffer := packets.Get()

	pkg := serverpackets.RestartResponse(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	pkg2 := serverpackets.CharSelectionInfo(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	client.SetState(clientStates.Authed)
	client.Send(buffer.Bytes())
	packets.Put(buffer)
}
