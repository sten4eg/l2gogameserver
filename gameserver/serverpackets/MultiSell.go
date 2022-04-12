package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/packets"
)

//Отправка пакета на открытие мультиселла
func MultisellShow(client *models.Client, msdata multisell.MultiList) {
	buffer := packets.Get()
	defer packets.Put(buffer)
	pkg := multiSell(client, msdata)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	client.SSend(buffer.Bytes())
}

//Отправка пакета
func multiSell(client *models.Client, msdata multisell.MultiList) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xD0)

	return buffer.Bytes()
}
