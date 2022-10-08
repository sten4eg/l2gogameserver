package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ExBrExtraUserInfo(character interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xDA)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(0) /** Event abnormal visual effects map. */
	buffer.WriteD(0) /** Lecture mark. */

	return buffer.Bytes()
}
