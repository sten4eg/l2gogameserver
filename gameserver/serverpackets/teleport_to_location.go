package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

//Телепорт к локации
//TODO: в будущем можно будет сделать направление персонажа после ТП.
func TeleportToLocation(clientI interfaces.ReciverAndSender, x, y, z, h int) []byte {
	client, ok := clientI.(*models.Client)
	if !ok {
		return []byte{}
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x22)
	buffer.WriteD(client.CurrentChar.ObjectId)
	buffer.WriteD(int32(x))
	buffer.WriteD(int32(y))
	buffer.WriteD(int32(z))
	buffer.WriteD(0x00) //IsValidation
	buffer.WriteD(int32(h))
	return buffer.Bytes()
}
