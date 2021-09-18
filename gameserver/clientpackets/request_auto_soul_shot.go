package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func RequestAutoSoulShot(data []byte, client *models.Client) []byte {
	var packet = packets.NewReader(data[2:])
	itemId := packet.ReadInt32()
	typee := packet.ReadInt32()

	client.CurrentChar.ActiveSoulShots = append(client.CurrentChar.ActiveSoulShots, itemId)

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xFE)

	buffer.WriteH(0x0c)
	buffer.WriteD(itemId)
	buffer.WriteD(typee)

	pkg := buffer.Bytes()

	return client.CryptAndReturnPackageReadyToShip(pkg)

}
