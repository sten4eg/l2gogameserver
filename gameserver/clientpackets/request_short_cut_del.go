package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestShortCutDel(data []byte, client *models.Client) []byte {
	var packet = packets.NewReader(data)
	id := packet.ReadInt32()
	slot := id % 12
	page := id / 12

	if page > 10 || page < 0 {
		return []byte{}
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	models.DeleteShortCut(slot, page, client)

	pkg := serverpackets.ShortCutInit(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
