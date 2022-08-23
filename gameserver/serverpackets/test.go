package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

// MagicSkillLaunched
func NewTest(clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return []byte{}
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x54)
	buffer.WriteD(client.CurrentChar.ObjectId)
	buffer.WriteD(1216)
	buffer.WriteD(1)
	buffer.WriteD(1)
	buffer.WriteD(client.CurrentChar.ObjectId)

	return buffer.Bytes()
}
