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
	playKey1 := packet.ReadInt32()
	playKey2 := packet.ReadInt32()
	loginKey1 := packet.ReadInt32()
	loginKey2 := packet.ReadInt32()
	_, _, _, _ = playKey1, playKey2, loginKey1, loginKey2
	serverpackets.NewCharSelectionInfo(db, client)
	client.Account.Login = login

}
