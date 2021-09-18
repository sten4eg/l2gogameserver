package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestTargetCancel(data []byte, client *models.Client) []byte {

	var packet = packets.NewReader(data)
	unselect := packet.ReadUInt16()
	_ = unselect

	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.TargetUnselected(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
