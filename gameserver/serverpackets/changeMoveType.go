package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

//В интернете пишется что устарел
func ChangeMoveType(client *models.Client, TargetObjectId int32) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x28)
	buffer.WriteD(TargetObjectId)
	buffer.WriteD(1) // 0 идти, 1 бежать
	buffer.WriteD(0)

	return buffer.Bytes()
}
