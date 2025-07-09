package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/chat"
	"l2gogameserver/packets"
)

func CreatureSay(say chat.Say, character interfaces.CharacterI) []byte {

	buffer := packets.Get()

	buffer.WriteSingleByte(0x4a)
	buffer.WriteD(character.GetObjectId()) //objId
	buffer.WriteD(say.Type)                //

	buffer.WriteS(character.GetName())

	buffer.WriteD(-1) // High Five NPCString ID
	buffer.WriteS(say.Text)

	return buffer.Bytes()
}
