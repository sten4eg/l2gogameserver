package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"l2gogameserver/utils"
)

func SkillList(character interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	skills := character.GetSkills()

	buffer.WriteSingleByte(0x5F)
	l := int32(len(skills))
	buffer.WriteD(l) // skill size

	for _, skill := range skills {

		buffer.WriteD(utils.BoolToInt32(skill.IsPassive())) // passiv ?
		buffer.WriteD(int32(skill.GetLevel()))              // level
		buffer.WriteD(skill.GetId())                        // id
		buffer.WriteSingleByte(0)                           // disable?
		buffer.WriteSingleByte(0)                           // enchant ?
	}

	return buffer.Bytes()
}
