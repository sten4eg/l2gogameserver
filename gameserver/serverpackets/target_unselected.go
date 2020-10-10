package serverpackets

import "l2gogameserver/gameserver/models"

func NewTargetUnselected(client *models.Client) {

	client.Buffer.WriteH(0)
	client.Buffer.WriteSingleByte(0x24)
	client.Buffer.WriteD(client.CurrentChar.CharId)
	client.Buffer.WriteD(client.CurrentChar.Coordinates.X)
	client.Buffer.WriteD(client.CurrentChar.Coordinates.Y)
	client.Buffer.WriteD(client.CurrentChar.Coordinates.Z)
	client.Buffer.WriteD(0)

}
