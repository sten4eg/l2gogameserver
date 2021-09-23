package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func CreatureSay(say *models.Say, character *models.Character) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x4a)
	buffer.WriteD(character.ObjectId) //objId
	buffer.WriteD(say.Type)           //

	buffer.WriteS(character.CharName)

	buffer.WriteD(-1) // High Five NPCString ID
	buffer.WriteS(say.Text)

	return buffer.Bytes()
}
