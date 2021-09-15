package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func GameGuardQuery(client *models.Client) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x74)
	buffer.WriteD(0x27533DD9)
	buffer.WriteD(0x2E72A51D)
	buffer.WriteD(0x2017038B)
	buffer.WriteDU(0xC35B1EA3)

	defer packets.Put(buffer)
	return buffer.Bytes()

}
