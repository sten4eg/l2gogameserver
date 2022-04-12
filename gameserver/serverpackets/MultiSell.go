package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/packets"
)

//Отправка пакета
func MultiSell(client *models.Client, msdata multisell.MultiList) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xD0)

	return buffer.Bytes()
}
