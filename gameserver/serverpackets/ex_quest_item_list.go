package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ExQuestItemList(client *models.Client) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xC6)
	buffer.WriteH(0)
	buffer.WriteH(0)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
