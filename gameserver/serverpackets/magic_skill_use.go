package serverpackets

import "l2gogameserver/gameserver/models"

func NewMagicSkillUse(client *models.Client, skill models.Skill, ctrlPressed, shiftPressed bool) {

	if skill.OperateType.IsPassive() {
		NewActionFailed(client)
		return
	}

	if client.CurrentChar.IsCastingNow {
		*client.CurrentChar.SkillQueue <- skill
		NewActionFailed(client)
		return
	}

	client.CurrentChar.IsCastingNow = true

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
