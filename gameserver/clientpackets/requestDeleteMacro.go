package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestDeleteMacro(client interfaces.ReciverAndSender, data []byte) {
	reader := packets.NewReader(data)

	id := reader.ReadInt32()

	character := client.GetCurrentChar()
	if character == nil {
		return
	}

	updateShortcut := character.DeleteMacros(id)
	rev := character.GetMacrosRevision()
	count := character.GetMacrosCount()

	for _, macros := range character.GetMacrosList() {
		client.SendBuf(serverpackets.SendMacroList(rev, count, macros))
	}

	if updateShortcut {
		client.EncryptAndSend(serverpackets.ShortCutInit(character))
	}
}
