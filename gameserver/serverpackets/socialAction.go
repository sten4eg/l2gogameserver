package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func SocialAction(character interfaces.CharacterI, actionId int32) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x27)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(actionId)

	return buffer.Bytes()
}
