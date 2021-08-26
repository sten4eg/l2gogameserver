package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewSkillList(client *models.Client) {

	skills := models.GetMySkills(client.CurrentChar.CharId)

	client.Buffer.WriteSingleByte(0x5F)

	client.Buffer.WriteD(int32(len(skills))) // skill size

	for _, skill := range skills {
		isPassive := int32(0)
		if skill.OperateType.IsPassive() {
			isPassive = 1
		}
		client.Buffer.WriteD(isPassive)                 // passiv ?
		client.Buffer.WriteD(int32(skill.CurrentLevel)) // level
		client.Buffer.WriteD(int32(skill.ID))           // id
		client.Buffer.WriteD(0)                         // disable?
		client.Buffer.WriteD(0)                         // enchant ?
	}

	client.SaveAndCryptDataInBufferToSend(true)
}
