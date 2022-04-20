package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

//TeleportToLocation Телепорт к локации
//TODO: в будущем можно будет сделать направление персонажа после ТП.
func TeleportToLocation(client interfaces.ReciverAndSender, x, y, z, h int) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x22)
	buffer.WriteD(client.GetCurrentChar().GetObjectId())
	buffer.WriteD(int32(x))
	buffer.WriteD(int32(y))
	buffer.WriteD(int32(z))
	buffer.WriteD(0x00) //IsValidation
	buffer.WriteD(int32(h))
	return buffer.Bytes()
}
