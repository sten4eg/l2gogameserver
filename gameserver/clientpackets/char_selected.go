package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func CharSelected(data []byte, clientI interfaces.ReciverAndSender) {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return
	}
	var read = packets.NewReader(data)
	charSlot := read.ReadInt32()
	_ = read.ReadUInt16() // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?
	_ = read.ReadInt32()  // unused, remove ?

	client.Account.CharSlot = charSlot

	pkg := serverpackets.SsqInfo(client)
	client.EncryptAndSend(pkg)

	pkg2 := serverpackets.CharSelected(client.Account.Char[client.Account.CharSlot], client)
	client.EncryptAndSend(pkg2)

	client.SetState(clientStates.Joining)
}
