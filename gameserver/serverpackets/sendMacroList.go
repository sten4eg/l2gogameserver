package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func SendMacroList(rev int32, count uint8, macros interfaces.MacrosInterface) *packets.Buffer {
	buffer := packets.Get()
	commands := macros.GetCommands()

	buffer.WriteSingleByte(0xE8)
	buffer.WriteD(rev)
	buffer.WriteSingleByte(0x00)
	buffer.WriteSingleByte(count)
	buffer.WriteSingleByte(1)

	buffer.WriteD(macros.GetId())
	buffer.WriteS(macros.GetName())
	buffer.WriteS(macros.GetDescription())
	buffer.WriteS(macros.GetAcronym())
	buffer.WriteSingleByte(macros.GetIcon())

	buffer.WriteSingleByte(byte(len(commands)))

	for i, cmd := range commands {
		buffer.WriteSingleByte(byte(i + 1))
		buffer.WriteSingleByte(cmd.GetType())
		buffer.WriteD(cmd.GetSkillId())
		buffer.WriteSingleByte(cmd.GetShortcutId())
		buffer.WriteS(cmd.GetName())
	}

	return buffer
}
