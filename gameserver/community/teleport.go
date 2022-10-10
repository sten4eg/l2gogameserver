package community

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/teleport"
	"l2gogameserver/gameserver/serverpackets"
)

func UserTeleport(client interfaces.ReciverAndSender, teleportID int) []byte {
	logger.Info.Println("Телепортация юзера...")
	locx, ok := teleport.GetTeleportID(teleportID)
	if !ok {
		logger.Info.Println("Not find teleport ID")
		return []byte{}
	}
	x, y, z, h := locx.X, locx.Y, locx.Z, 0

	return serverpackets.TeleportToLocation(client.GetCurrentChar(), x, y, z, h)
}
