package clientpackets

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func MoveToPawn(client interfaces.ReciverAndSender, data []byte) {
	var packet = packets.NewReader(data)
	charId := packet.ReadInt32()
	targetId := packet.ReadInt32()
	distance := packet.ReadInt32()
	X := packet.ReadInt32()
	Y := packet.ReadInt32()
	Z := packet.ReadInt32()
	tX := packet.ReadInt32()
	tY := packet.ReadInt32()
	tZ := packet.ReadInt32()

	logger.Info.Println(charId, targetId, distance, X, Y, Z, tX, tY, tZ)

	//buffer := packets.Get()
	//defer packets.Put(buffer)
	//pkg := serverpackets.MoveToPawn(client.CurrentChar)
	//buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

}
