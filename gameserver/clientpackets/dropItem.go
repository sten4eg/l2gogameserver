package clientpackets

import (
	"l2gogameserver/gameserver/broadcast"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func DropItem(client interfaces.ReciverAndSender, data []byte) {
	var read = packets.NewReader(data)
	objectId := read.ReadInt32()
	count := read.ReadInt64()
	x := read.ReadInt32()
	y := read.ReadInt32()
	z := read.ReadInt32()

	activeChar := client.GetCurrentChar()

	dropItem, updateItem := activeChar.DropItem(objectId, count)

	if dropItem != nil {
		dropItem.SetCoordinate(x, y, z)

		if updateItem != nil {
			items := []interfaces.MyItemInterface{updateItem}
			msg := serverpackets.InventoryUpdate(items)
			client.EncryptAndSend(msg)
		}

		buffer := serverpackets.DropItem(dropItem, activeChar.GetObjectId())

		activeChar.GetCurrentRegion().AddVisibleItems(dropItem)

		broadcast.BroadCastBufferToAroundPlayers(client, buffer)
	}

}
