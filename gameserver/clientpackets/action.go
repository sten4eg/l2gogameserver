package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/serverpackets"
	"l2gogameserver/packets"
)

func NewAction(data []byte, client *models.Client) {
	var packet = packets.NewReader(data)

	objectId := packet.ReadInt32() //Target
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	actionId := packet.ReadSingleByte() // Action identifier : 0-Simple click, 1-Shift click

	_, _, _, _, _ = objectId, originX, originY, originZ, actionId

	//Очень много Логика по action
	//serverpackets.NewSocialAction(client)
	//client.SimpleSend(client.Buffer.Bytes(), true) - simpleSend removed
	serverpackets.NewTargetSelected(client.CurrentChar.CharId, objectId, originX, originY, originZ, client)
}
