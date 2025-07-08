package clientpackets

import (
	"database/sql"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
)

func RequestGoToLobby(client interfaces.NewClientCtxInterface, db *sql.DB) {
	client.SendBuf(serverpackets.CharSelectionInfo(client, db))
}
