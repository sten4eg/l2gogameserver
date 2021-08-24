package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewRequestShortCutDel(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)
	id := packet.ReadInt32()
	slot := id % 12
	page := id / 12

	if page > 10 || page < 0 {
		return
	}

	models.DeleteShortCut(slot, page, client.CurrentChar.CharId, client.CurrentChar.ClassId)
	serverpackets.NewShortCutInit(client)
}
