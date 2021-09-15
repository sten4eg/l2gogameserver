package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func QuestList(client *models.Client) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x86)
	buffer.WriteH(0)
	x := make([]byte, 128)
	buffer.WriteSlice(x)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
