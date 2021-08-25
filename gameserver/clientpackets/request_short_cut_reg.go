package clientpackets

import (
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewRequestShortCutReg(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)

	typeId := packet.ReadInt32()

	if (typeId < 1) || (typeId > 6) {
		typeId = 0
	}
	shortType := dto.ShortTypes[typeId]

	slotFromRequest := packet.ReadInt32()
	slot := slotFromRequest % 12
	page := slotFromRequest / 12

	id := packet.ReadInt32() // obj_id (form DB)
	lvl := packet.ReadInt32()
	characterType := packet.ReadInt32()

	if page > 10 || page < 0 {
		return
	}
	sc := dto.GetShortCutDTO(slot, page, id, lvl, characterType, shortType)

	models.RegisterShortCut(sc, client)
	serverpackets.NewShortCutRegister(sc, client)

}
