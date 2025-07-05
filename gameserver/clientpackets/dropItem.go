package clientpackets

import (
	"database/sql"
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func DropItem(client interfaces.ReciverAndSender, data []byte, db *sql.DB) {
	var read = packets.NewReader(data)
	objectId := read.ReadInt32()
	count := read.ReadInt64()
	x := read.ReadInt32()
	y := read.ReadInt32()
	z := read.ReadInt32()

	character := client.GetCurrentChar()

	dropItem, updateItem := character.DropItem(objectId, count, db)

	if dropItem != nil {
		dropItem.SetCoordinate(x, y, z)

		if updateItem != nil {
			items := []interfaces.MyItemInterface{updateItem}
			msg := serverpackets.InventoryUpdate(items)
			client.EncryptAndSend(msg)
		}

		buffer := serverpackets.DropItem(dropItem, character.GetObjectId())

		character.GetCurrentRegion().AddVisibleItems(dropItem)

		broadcast.BroadCastBufferToAroundPlayers(client, buffer)
	}

}
