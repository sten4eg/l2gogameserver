package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func MoveToPawn(client *models.Character) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x72)
	buffer.WriteD(client.ObjectId)
	buffer.WriteD(client.Target)
	buffer.WriteD(0)

	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	return buffer.Bytes()

}
