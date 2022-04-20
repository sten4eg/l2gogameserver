package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestMakeMacro(clientI interfaces.ReciverAndSender, data []byte) {
	client, ok := clientI.(*models.Client)
	if !ok {
		return
	}

	var packet = packets.NewReader(data)

	macro := models.Macro{
		Id:          packet.ReadInt32(),
		Name:        packet.ReadString(),
		Description: packet.ReadString(),
		Acronym:     packet.ReadString(),
		Icon:        packet.ReadSingleByte(),
		Count:       packet.ReadSingleByte(),
	}
	for i := 1; i <= int(macro.Count); i++ {
		macro.Commands = append(macro.Commands, models.MacroCommand{
			Index:      packet.ReadSingleByte(), // command count
			Type:       packet.ReadSingleByte(), // type 1 = skill, 3 = action, 4 = shortcut
			SkillID:    packet.ReadInt32(),      // skill id
			ShortcutID: packet.ReadSingleByte(), // shortcut id
			Name:       packet.ReadString(),     // command name
		})
	}

	client.CurrentChar.AddMacros(macro)
	count := client.CurrentChar.MacrosesCount()
	pkg := serverpackets.MacroMake(macro, count)
	client.EncryptAndSend(pkg)
}
