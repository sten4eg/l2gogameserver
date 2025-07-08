package clientpackets

import (
	"database/sql"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func RequestShortCutDel(data []byte, client interfaces.NewClientCtxInterface, db *sql.DB) {

	var packet = packets.NewReader(data)
	id := packet.ReadInt32()
	slot := id % 12
	page := id / 12

	if page > 10 || page < 0 {
		return
	}

	models.DeleteShortCut(slot, page, client, db)

	pkg := serverpackets.ShortCutInit(client.GetAccount().GetCurrentChar(), db)
	client.EncryptAndSend(pkg)

}
