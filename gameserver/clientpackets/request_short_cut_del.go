package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestShortCutDel(data []byte, clientI interfaces.ReciverAndSender) {
	client, ok := clientI.(*models.Client)
	if !ok {
		return
	}

	var packet = packets.NewReader(data)
	id := packet.ReadInt32()
	slot := id % 12
	page := id / 12

	if page > 10 || page < 0 {
		return
	}

	models.DeleteShortCut(slot, page, client)

	pkg := serverpackets.ShortCutInit(client)
	client.EncryptAndSend(pkg)

}
