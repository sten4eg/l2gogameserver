package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func CharCreateOk(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x0f)
	buffer.WriteD(1)
	return buffer.Bytes()
}
