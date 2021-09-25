package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
	"log"
)

func MoveToPawn(client *models.Client, data []byte) {
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

	log.Println(charId, targetId, distance, X, Y, Z, tX, tY, tZ)

	//buffer := packets.Get()
	//defer packets.Put(buffer)
	//pkg := serverpackets.MoveToPawn(client.CurrentChar)
	//buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

}
