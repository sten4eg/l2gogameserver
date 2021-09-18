package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestSkillList(client *models.Client, data []byte) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg1 := serverpackets.SkillList(client)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg1))

	return buffer.Bytes()
}
