package clientpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func Attack(data []byte, client interfaces.CharacterI) {
	var packet = packets.NewReader(data)

	objId := packet.ReadInt32() // targetObjId
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	attackId := packet.ReadSingleByte() // 0 for simple click 1 for shift-click

	_ = attackId

	pkg := serverpackets.Attack(client, objId, originX, originY, originZ)
	client.EncryptAndSend(pkg)
}
