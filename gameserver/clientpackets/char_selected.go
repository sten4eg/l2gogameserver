package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func CharSelected(data []byte, client *models.Client) []byte {

	var read = packets.NewReader(data)
	charSlot := read.ReadInt32()
	_ = read.ReadUInt16() // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?

	client.Account.CharSlot = charSlot

	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.SsqInfo(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	pkg2 := serverpackets.CharSelected(client.Account.Char[client.Account.CharSlot], client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	return buffer.Bytes()
}
