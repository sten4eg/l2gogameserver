package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/skills/targets"
)

func NewMagicSkillUse(client *models.Client, skill models.Skill, ctrlPressed, shiftPressed bool) {

	if skill.OperateType.IsPassive() {
		ActionFailed(client)
		return
	}

	if client.CurrentChar.IsCastingNow {
		currSkill := client.CurrentChar.CurrentSkill
		if currSkill != nil && skill.ID == currSkill.Skill.ID {
			ActionFailed(client)
			return
		} //todo тут еще есть elseif isSkillDisabled()

		client.CurrentChar.SetSkillToQueue(skill, ctrlPressed, shiftPressed)
		ActionFailed(client)
		return
	}

	client.CurrentChar.IsCastingNow = true
	client.CurrentChar.CurrentSkill = &models.SkillHolder{
		Skill:        skill,
		CtrlPressed:  ctrlPressed,
		ShiftPressed: shiftPressed,
	}

	var target int32
	switch skill.TargetType {
	case targets.AURA, targets.FRONT_AURA, targets.BEHIND_AURA, targets.GROUND, targets.SELF, targets.AURA_CORPSE_MOB, targets.COMMAND_CHANNEL, targets.AURA_FRIENDLY, targets.AURA_UNDEAD_ENEMY:
		target = 0
	default:
		target = client.CurrentChar.CurrentTargetId
	}

	// запускаем обработчик скилла
	_ = target

	/////////////////////////////////////////////////////////////////////////////////
	client.Buffer.WriteSingleByte(0x48)
	client.Buffer.WriteD(client.CurrentChar.CharId) // activeChar id
	client.Buffer.WriteD(client.CurrentChar.CharId) // targetChar id
	client.Buffer.WriteD(int32(skill.ID))           // skillId
	client.Buffer.WriteD(int32(skill.Levels))       // skillLevel
	client.Buffer.WriteD(int32(skill.HitTime))      // hitTime
	client.Buffer.WriteD(int32(skill.ReuseDelay))   // reuseDelay

	x, y, z := client.CurrentChar.GetXYZ()
	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)

	client.Buffer.WriteH(0) //size???
	// for  by size ???

	client.Buffer.WriteH(0) // _groundLocations.size()
	// for by _groundLocations.size()

	//location target
	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)
	client.SaveAndCryptDataInBufferToSend(true)

}
