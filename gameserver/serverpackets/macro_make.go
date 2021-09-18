package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func MacroMake(macro models.Macro, count uint8) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0xE8)

	buffer.WriteD(0)                  // macro change revision (changes after each macro edition)
	buffer.WriteSingleByte(0x00)      // unknown
	buffer.WriteSingleByte(count + 1) // count of Macros
	buffer.WriteSingleByte(1)         // unknown

	buffer.WriteD(macro.Id)            // Macro ID
	buffer.WriteS(macro.Name)          // Macro Name
	buffer.WriteS(macro.Desc)          // Desc
	buffer.WriteS(macro.Acronym)       // acronym
	buffer.WriteSingleByte(macro.Icon) // icon

	buffer.WriteSingleByte(uint8(len(macro.Commands))) // count

	for _, command := range macro.Commands {
		buffer.WriteSingleByte(command.Index)      // command count
		buffer.WriteSingleByte(command.Type)       // type 1 = skill, 3 = action, 4 = shortcut
		buffer.WriteD(command.SkillID)             // skill id
		buffer.WriteSingleByte(command.ShortcutID) // shortcut id
		buffer.WriteS(command.Name)                // command name
	}
	return buffer.Bytes()
}
