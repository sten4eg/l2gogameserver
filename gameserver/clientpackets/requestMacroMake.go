package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestMakeMacro(client interfaces.ReciverAndSender, data []byte) {
	var reader = packets.NewReader(data)

	macros := models.Macro{
		Id:          reader.ReadInt32(),
		Name:        reader.ReadString(),
		Description: reader.ReadString(),
		Acronym:     reader.ReadString(),
		Icon:        reader.ReadSingleByte(),
		Count:       reader.ReadSingleByte(),
	}
	var commandsLength int
	for i := 1; i <= int(macros.Count); i++ {
		macros.Commands = append(macros.Commands, models.MacroCommand{
			Index:      reader.ReadSingleByte(), // command count
			Type:       reader.ReadSingleByte(), // type 1 = skill, 3 = action, 4 = shortcut
			SkillID:    reader.ReadInt32(),      // skill id
			ShortcutID: reader.ReadSingleByte(), // shortcut id
			//Name:       reader.ReadString(),     // command name
		})
		name := reader.ReadString()
		macros.Commands[i-1].Name = name
		commandsLength += len(name)
	}

	character := client.GetCurrentChar()

	if character == nil {
		return
	}

	if commandsLength > 255 {
		character.SendSysMsg(sysmsg.InvalidMacro)
		return
	}

	if character.GetMacrosCount() > 48 {
		character.SendSysMsg(sysmsg.YouMayCreateUpTo48Macros)
		return
	}

	if macros.GetName() == "" {
		character.SendSysMsg(sysmsg.EnterTheMacroName)
		return
	}

	if len(macros.GetDescription()) > 32 {
		character.SendSysMsg(sysmsg.MacroDescriptionMax32Chars)
		return
	}

	character.AddMacros(&macros)

	count := character.GetMacrosCount()
	rev := character.GetMacrosRevision()

	for _, macros := range character.GetMacrosList() {
		client.SendBuf(serverpackets.SendMacroList(rev, count, macros))
	}
}
