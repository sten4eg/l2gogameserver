package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

// TODO убрать модель
func SendMacroList(rev int32, count uint8, macro models.Macro) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xE8)

	buffer.WriteD(rev)
	buffer.WriteSingleByte(0x00)
	buffer.WriteSingleByte(count)
	buffer.WriteSingleByte(1)

	buffer.WriteD(macro.Id)
	buffer.WriteS(macro.Name)
	buffer.WriteS(macro.Description)
	buffer.WriteS(macro.Acronym)
	buffer.WriteSingleByte(macro.Icon)

	buffer.WriteSingleByte(byte(len(macro.Commands)))

	for i, cmd := range macro.Commands {
		buffer.WriteSingleByte(byte(i + 1))
		buffer.WriteSingleByte(cmd.Type)
		buffer.WriteD(cmd.Id)
		buffer.WriteSingleByte(cmd.ShortcutID)
		buffer.WriteS(cmd.Name)
	}

	return buffer
}
