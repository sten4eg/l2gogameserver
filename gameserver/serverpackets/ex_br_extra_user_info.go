package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ExBrExtraUserInfo(client *models.Client) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xDA)
	buffer.WriteD(1)
	buffer.WriteD(0)
	buffer.WriteD(0)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
