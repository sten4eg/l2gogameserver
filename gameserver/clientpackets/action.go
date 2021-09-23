package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func Action(data []byte, client *models.Client) []byte {
	var packet = packets.NewReader(data)

	objectId := packet.ReadInt32() //Target
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	actionId := packet.ReadSingleByte() // Action identifier : 0-Simple click, 1-Shift click

	client.CurrentChar.CurrentTargetId = objectId
	_, _, _, _, _ = objectId, originX, originY, originZ, actionId

	buffer := packets.Get()
	defer packets.Put(buffer)

	client.CurrentChar.Target = objectId
	pkg := serverpackets.TargetSelected(client.CurrentChar.ObjectId, objectId, originX, originY, originZ)
	buffer.WriteSlice(client.CryptAndReturnPackageReadyToShip(pkg))

	return buffer.Bytes()
}
