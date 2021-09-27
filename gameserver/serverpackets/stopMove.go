package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func StopMove(client *models.Client, TargetObjectId, x, y, z, heading int32) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x47)
	buffer.WriteD(TargetObjectId)
	buffer.WriteD(x) // 0 идти, 1 бежать
	buffer.WriteD(y)
	buffer.WriteD(z)
	buffer.WriteD(heading)

	return buffer.Bytes()
}
