package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestSkillList(client interfaces.ReciverAndSender, data []byte) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg1 := serverpackets.SkillList(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg1))

	return buffer.Bytes()
}
