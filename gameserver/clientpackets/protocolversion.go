package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func ProtocolVersion(data []byte, client *models.Client) []byte {

	var packet = packets.NewReader(data)
	protocolVersion := packet.ReadUInt16() //todo check !=273
	_ = protocolVersion
	buffer := packets.Get()

	pkg1 := serverpackets.KeyPacket(client)
	buffer.WriteSlice(client.ReturnPackageReadyToShip(pkg1))

	defer packets.Put(buffer)
	return buffer.Bytes()

}
