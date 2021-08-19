package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewSkillList(client *models.Client) {

	skills := models.GetMySkills(client.CurrentChar.CharId)

	client.Buffer.WriteSingleByte(0x5F)

	client.Buffer.WriteD(int32(len(skills))) // skill size

	for _, v := range skills {
		client.Buffer.WriteD(v.IsPassive())         // passiv ?
		client.Buffer.WriteD(int32(v.CurrentLevel)) // level
		client.Buffer.WriteD(int32(v.ID))           // id
		client.Buffer.WriteD(0)                     // disable?
		client.Buffer.WriteD(0)                     // enchant ?
	}

	client.SaveAndCryptDataInBufferToSend(true)
}
