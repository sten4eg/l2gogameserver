package serverpackets

import "l2gogameserver/packets"

// TODO принимает на вход персонажа
func MyTargetSelected(targetId int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xB9)
	buffer.WriteD(targetId)
	buffer.WriteH(0) // color character.level - target.level
	buffer.WriteD(0x00)

	return buffer
}
