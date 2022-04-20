package community

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/teleport"
	"l2gogameserver/gameserver/serverpackets"
	"log"
)

func UserTeleport(client interfaces.ReciverAndSender, teleportID int) []byte {
	log.Println("Телепортация юзера...")
	locx, ok := teleport.GetTeleportID(teleportID)
	if !ok {
		log.Println("Not find teleport ID")
		return []byte{}
	}
	x, y, z, h := locx.X, locx.Y, locx.Z, 0

	return serverpackets.TeleportToLocation(client, x, y, z, h)
}
