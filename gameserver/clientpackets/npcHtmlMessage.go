package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NpcHtmlMessage(client *models.Client, npcid int32) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.NpcHtmlMessage(client, npcid)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
