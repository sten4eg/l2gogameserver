package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NpcHtmlMessage(client interfaces.NewClientCtxInterface, npcid int32) []byte {
	buffer := packets.Get()

	pkg := serverpackets.NpcHtmlMessage(npcid)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
