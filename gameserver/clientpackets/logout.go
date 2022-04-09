package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func Logout(data []byte, client *models.Client) []byte {
	pkg := serverpackets.LogoutToClient(data, client)
	buffer := packets.Get()
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	return buffer.Bytes()
}
