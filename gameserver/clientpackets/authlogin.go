package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func AuthLogin(data []byte, clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.Client)
	if !ok {
		return []byte{}
	}

	var packet = packets.NewReader(data)

	login := packet.ReadString()
	client.Account.Login = login
	playKey1 := packet.ReadInt32()
	playKey2 := packet.ReadInt32()
	loginKey1 := packet.ReadInt32()
	loginKey2 := packet.ReadInt32()
	_, _, _, _ = playKey1, playKey2, loginKey1, loginKey2

	buffer := packets.Get()
	pkg := serverpackets.CharSelectionInfo(client)
	client.Account.Login = login

	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	defer packets.Put(buffer)
	return buffer.Bytes()
}
