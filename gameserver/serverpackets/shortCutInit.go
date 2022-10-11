package serverpackets

import (
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

// TODO убрать модель
func ShortCutInit(character interfaces.CharacterI) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x45)

  shortCuts := dto.GetAllShortCuts(character.GetObjectId(), character.GetClassId())
	buffer.WriteD(int32(len(shortCuts)))

	for _, v := range shortCuts {
		buffer.WriteD(v.ShortcutType)
		buffer.WriteD(v.Slot + (v.Page * models.MaxShortcutsPerBar))

		buffer.WriteD(v.Id)

		shortCutsType := dto.ShortTypes[v.ShortcutType]
		switch shortCutsType {
		case "ITEM":
			buffer.WriteD(0x01)
			buffer.WriteD(0) //sc.getSharedReuseGroup()
			buffer.WriteD(0)
			buffer.WriteD(0)
			buffer.WriteH(0)
			buffer.WriteH(0)
		case "SKILL":
			buffer.WriteD(v.Level)
			buffer.WriteSingleByte(0) // C5
			buffer.WriteD(0x01)       // C6
		case "ACTION", "MACRO", "RECIPE", "BOOKMARK":
			buffer.WriteD(0x01) // C6
		}
	}

	return buffer.Bytes()
}
