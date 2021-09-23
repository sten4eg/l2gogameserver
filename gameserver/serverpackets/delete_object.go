package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func DeleteObject(character *models.Character) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x08)
	buffer.WriteD(character.ObjectId)
	buffer.WriteD(0)
	return buffer.Bytes()
}
