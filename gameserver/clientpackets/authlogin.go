package clientpackets

import (
	"github.com/jackc/pgx"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewAuthLogin(data []byte, client *models.Client, db *pgx.Conn) {

	var packet = packets.NewReader(data)

	login := packet.ReadString()
	client.Account.Login = login
	client.Account = serverpackets.NewCharSelectionInfo(db, client)
	client.Account.Login = login
	client.SimpleSend(client.Buffer.Bytes(), true)

}
