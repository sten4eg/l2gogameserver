package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NpcHtmlMessage(client interfaces.ReciverAndSender, npcid int32) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.NpcHtmlMessage(client, npcid)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
