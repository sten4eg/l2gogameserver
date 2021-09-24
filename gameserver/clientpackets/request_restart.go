package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestRestart(data []byte, client *models.Client) []byte {

	client.SaveUser()
	//todo need save in db

	_ = data
	buffer := packets.Get()

	pkg := serverpackets.RestartResponse(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	pkg2 := serverpackets.CharSelectionInfo(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	defer packets.Put(buffer)
	return buffer.Bytes()

}
