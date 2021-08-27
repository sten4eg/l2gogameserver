package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewSkillList(client *models.Client) {

	skills := client.CurrentChar.Skills

	client.Buffer.WriteSingleByte(0x5F)
	l := int32(len(skills))
	client.Buffer.WriteD(l) // skill size

	for _, skill := range skills {
		isPassive := int32(0)
		if skill.OperateType.IsPassive() {
			isPassive = 1
		}
		client.Buffer.WriteD(isPassive)           // passiv ?
		client.Buffer.WriteD(int32(skill.Levels)) // level
		client.Buffer.WriteD(int32(skill.ID))     // id
		client.Buffer.WriteSingleByte(0)          // disable?
		client.Buffer.WriteSingleByte(0)          // enchant ?
	}

	client.SaveAndCryptDataInBufferToSend(true)
}
