package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func InventoryUpdate(client *models.Client, item *models.MyItem, updateType int16) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.InventoryUpdate(client, item, updateType)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
