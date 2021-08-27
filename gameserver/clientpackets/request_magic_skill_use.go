package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewRequestMagicSkillUse(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)

	magicId := packet.ReadInt32()                // Identifier of the used skill
	ctrlPressed := packet.ReadInt32() != 0       // True if it's a ForceAttack : Ctrl pressed
	shiftPressed := packet.ReadSingleByte() != 0 // True if Shift pressed

	if client.CurrentChar.IsDead {
		serverpackets.NewActionFailed(client)
		return
	}

	if client.CurrentChar.IsFakeDeath {
		serverpackets.NewSystemMessage(models.CantMoveSitting, client)
		serverpackets.NewActionFailed(client)
		return
	}

	skill, exist := client.CurrentChar.Skills[int(magicId)]
	if !exist {
		// todo тут еще идут проверки, возможно это кастомный? скилл или скилл трансформы и если нет то фейл
		serverpackets.NewActionFailed(client)
		return
	}
	_, _, _ = magicId, ctrlPressed, shiftPressed

	serverpackets.NewSetupGauge(client)
	serverpackets.NewMagicSkillUse(client, skill, ctrlPressed, shiftPressed)

	//todo что это такое  / go serverpackets.NewTest(client)
}
