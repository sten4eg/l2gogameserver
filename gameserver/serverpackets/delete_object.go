package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func DeleteObject(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x08)
	buffer.WriteD(client.CurrentChar.CharId)
	buffer.WriteD(0)
	return buffer.Bytes()
}
