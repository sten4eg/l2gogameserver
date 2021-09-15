package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func CreatureSay(say *models.Say, user *models.Character) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x4a)
	buffer.WriteD(user.CharId) //objId
	buffer.WriteD(say.Type)

	buffer.WriteS(user.CharName)

	buffer.WriteD(-1) // High Five NPCString ID
	buffer.WriteS(say.Text)

	return buffer.Bytes()
}
