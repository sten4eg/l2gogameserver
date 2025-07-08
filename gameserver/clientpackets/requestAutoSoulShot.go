package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func RequestAutoSoulShot(data []byte, client interfaces.NewClientCtxInterface) {

	var packet = packets.NewReader(data[2:])
	itemId := packet.ReadInt32()
	typee := packet.ReadInt32()

	client.GetAccount().GetCurrentChar().AddActiveSoulShots(itemId)
	//todo реализцая должна быть в serverPackets
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)

	buffer.WriteH(0x0c)
	buffer.WriteD(itemId)
	buffer.WriteD(typee)

	pkg := buffer.Bytes()
	client.EncryptAndSend(pkg)

}
