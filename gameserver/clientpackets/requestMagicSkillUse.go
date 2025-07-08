package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestMagicSkillUse(data []byte, client interfaces.NewClientCtxInterface) {

	var packet = packets.NewReader(data)

	magicId := packet.ReadInt32()                // Identifier of the used skill
	ctrlPressed := packet.ReadInt32() != 0       // True if it's a ForceAttack : Ctrl pressed
	shiftPressed := packet.ReadSingleByte() != 0 // True if Shift pressed

	buffer := packets.Get()
	currChar := client.GetAccount().GetCurrentChar()
	if currChar.IsDead() {
		pkg := serverpackets.ActionFailed()
		buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
		client.Send(buffer.Bytes())
		return
	}

	if currChar.IsFakeDead() {
		pkg := sysmsg.SystemMessage(sysmsg.CantMoveSitting)
		buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
		pkg2 := serverpackets.ActionFailed()
		buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))
		client.Send(buffer.Bytes())
		return
	}

	skill, exist := currChar.GetSkillById(magicId)
	if !exist {
		// todo тут еще идут проверки, возможно это кастомный? скилл или скилл трансформы и если нет то фейл
		pkg := serverpackets.ActionFailed()
		buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
		client.Send(buffer.Bytes())
		return
	}
	_, _, _ = magicId, ctrlPressed, shiftPressed

	if skill.IsPassive() {
		pkg := serverpackets.ActionFailed()
		buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
		client.Send(buffer.Bytes())
		return
	}

	if currChar.IsCastingNow() {
		currSkill := currChar.GetCurrentSkill()
		if currSkill != nil && skill.GetId() == currSkill.GetSkill().GetId() {
			pkg := serverpackets.ActionFailed()
			buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
			client.Send(buffer.Bytes())
			return
		} //todo тут еще есть elseif isSkillDisabled()

		currChar.SetSkillToQueue(skill, ctrlPressed, shiftPressed)
		pkg := serverpackets.ActionFailed()
		buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
		client.Send(buffer.Bytes())
		return
	}

	pkg2 := serverpackets.NewMagicSkillUse(client, skill, ctrlPressed, shiftPressed)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg2))

	pkg := serverpackets.SetupGauge(currChar)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))
	client.Send(buffer.Bytes())
}
