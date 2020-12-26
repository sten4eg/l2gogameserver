package clientpackets

import (
	"github.com/jackc/pgx"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
)

func NewRequestRestart(data []byte, client *models.Client, db *pgx.Conn) {

	//todo need save in db

	_ = data
	serverpackets.NewRestartResponse(client)
	client.SimpleSend(client.Buffer.Bytes(), true)

	serverpackets.NewCharSelectionInfo(db, client)
}
