package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func SendMacroList(client *models.Client) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xE8)
	buffer.WriteD(0)
	buffer.WriteSingleByte(0)
	buffer.WriteSingleByte(0)
	buffer.WriteSingleByte(0)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
