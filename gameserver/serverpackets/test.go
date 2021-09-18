package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

//MagicSkillLaunched
func NewTest(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x54)
	buffer.WriteD(client.CurrentChar.CharId)
	buffer.WriteD(1216)
	buffer.WriteD(1)
	buffer.WriteD(1)
	buffer.WriteD(client.CurrentChar.CharId)

	return buffer.Bytes()
}
