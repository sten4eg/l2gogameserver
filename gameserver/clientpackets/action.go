package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func Action(data []byte, client *models.Client) ([]byte, int32, byte, bool) {
	reAppeal := false // повторное обращение к объекту
	var packet = packets.NewReader(data)
	objectId := packet.ReadInt32() //Target
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	actionId := packet.ReadSingleByte() // Action identifier : 0-Simple click, 1-Shift click
	if objectId == client.CurrentChar.Target {
		reAppeal = true
	} else {
		client.CurrentChar.Target = objectId
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	pkg := serverpackets.TargetSelected(client.CurrentChar.ObjectId, objectId, originX, originY, originZ)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes(), objectId, actionId, reAppeal
}
