package clientpackets

import (
	"database/sql"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func CharSelected(data []byte, client interfaces.NewClientCtxInterface, db *sql.DB) {
	var read = packets.NewReader(data)
	charSlot := read.ReadInt32()
	_ = read.ReadUInt16() // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?

	client.GetAccount().SetSelectedChar(int(charSlot))

	pkg := serverpackets.SsqInfo()
	client.EncryptAndSend(pkg)

	pkg2 := serverpackets.CharSelected(client.GetAccount().GetCurrentChar(), db, client.GetAccount().GetLogin())
	client.EncryptAndSend(pkg2)

	client.SetState(clientStates.Joining)
}
