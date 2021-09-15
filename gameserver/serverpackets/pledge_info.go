package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func PledgeInfo(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x89)
	buffer.WriteD(0)
	buffer.WriteS("")
	buffer.WriteS("")

	return buffer.Bytes()
}
