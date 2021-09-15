package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ShortBuffStatusUpdate(client *models.Client) []byte {
	buffer := packets.Get()
	buffer.WriteSingleByte(0xfa)
	buffer.WriteD(1242)
	buffer.WriteD(1)
	buffer.WriteD(20)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
