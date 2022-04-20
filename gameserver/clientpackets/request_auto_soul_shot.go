package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func RequestAutoSoulShot(data []byte, clientI interfaces.ReciverAndSender) {
	client, ok := clientI.(*models.Client)
	if !ok {
		return
	}

	var packet = packets.NewReader(data[2:])
	itemId := packet.ReadInt32()
	typee := packet.ReadInt32()

	client.CurrentChar.ActiveSoulShots = append(client.CurrentChar.ActiveSoulShots, itemId)
	//todo реализцая должна быть в serverPackets
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xFE)

	buffer.WriteH(0x0c)
	buffer.WriteD(itemId)
	buffer.WriteD(typee)

	pkg := buffer.Bytes()
	client.EncryptAndSend(pkg)

}
