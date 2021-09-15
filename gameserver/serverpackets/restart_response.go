package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func RestartResponse(client *models.Client) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x71)
	buffer.WriteD(1)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
