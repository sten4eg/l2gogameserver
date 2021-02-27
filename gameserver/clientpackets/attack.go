package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewAttack(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	attackId := packet.ReadSingleByte() // 0 for simple click 1 for shift-click

	var A serverpackets.Attack
	A.TargetId = objId
	A.X = originX
	A.Z = originZ
	A.Y = originY
	serverpackets.NewAttack(client, &A)
	_, _, _, _, _ = objId, originZ, originX, originY, attackId
}
