package community

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/teleport"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
	"log"
)

func UserTeleport(client interfaces.ReciverAndSender, teleportID int) {
	log.Println("Телепортация юзера...")
	locx, ok := teleport.GetTeleportID(teleportID)
	if !ok {
		log.Println("Not find teleport ID")
		return
	}
	x, y, z, h := locx.X, locx.Y, locx.Z, 0

	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.TeleportToLocation(client, x, y, z, h)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	client.SSend(buffer.Bytes())
}
