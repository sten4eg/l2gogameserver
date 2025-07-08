package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"

	"l2gogameserver/gameserver/models/skills/targets"
	"l2gogameserver/packets"
)

func NewMagicSkillUse(client interfaces.NewClientCtxInterface, skill interfaces.SkillInterface, ctrlPressed, shiftPressed bool) []byte {

	buffer := packets.Get()
	currChar := client.GetAccount().GetCurrentChar()

	//currChar.IsCastingNow = true
	//client.CurrentChar.CurrentSkill = &models.SkillHolder{
	//	Skill:        skill,
	//	CtrlPressed:  ctrlPressed,
	//	ShiftPressed: shiftPressed,
	//}

	var target int32
	switch skill.GetTargetType() {
	case int(targets.AURA), int(targets.FRONT_AURA), int(targets.BEHIND_AURA), int(targets.GROUND), int(targets.SELF), int(targets.AURA_CORPSE_MOB), int(targets.COMMAND_CHANNEL), int(targets.AURA_FRIENDLY), int(targets.AURA_UNDEAD_ENEMY):
		target = 0
	default:
		target = currChar.GetTarget()
	}

	// запускаем обработчик скилла
	_ = target

	/////////////////////////////////////////////////////////////////////////////////
	buffer.WriteSingleByte(0x48)
	buffer.WriteD(currChar.GetObjectId())       // activeChar id
	buffer.WriteD(currChar.GetObjectId())       // targetChar id
	buffer.WriteD(int32(skill.GetId()))         // skillId
	buffer.WriteD(int32(skill.GetLevel()))      // skillLevel
	buffer.WriteD(int32(skill.GetHitTime()))    // hitTime
	buffer.WriteD(int32(skill.GetReuseDelay())) // reuseDelay

	x, y, z := currChar.GetXYZ()
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)

	buffer.WriteH(0) //size???
	// for  by size ???

	buffer.WriteH(0) // _groundLocations.size()
	// for by _groundLocations.size()

	//location target
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)

	return buffer.Bytes()
}
