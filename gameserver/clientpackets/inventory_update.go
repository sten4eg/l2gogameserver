package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func InventoryUpdate(client interfaces.ReciverAndSender, item *models.MyItem, updateType int16) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.InventoryUpdate(item, updateType)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
