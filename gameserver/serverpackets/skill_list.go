package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func SkillList(clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.Client)
	if !ok {
		return []byte{}
	}

	buffer := packets.Get()

	skills := client.CurrentChar.Skills

	buffer.WriteSingleByte(0x5F)
	l := int32(len(skills))
	buffer.WriteD(l) // skill size

	for _, skill := range skills {
		isPassive := int32(0)
		if skill.OperateType.IsPassive() {
			isPassive = 1
		}
		buffer.WriteD(isPassive)           // passiv ?
		buffer.WriteD(int32(skill.Levels)) // level
		buffer.WriteD(int32(skill.ID))     // id
		buffer.WriteSingleByte(0)          // disable?
		buffer.WriteSingleByte(0)          // enchant ?
	}

	defer packets.Put(buffer)
	return buffer.Bytes()
}
