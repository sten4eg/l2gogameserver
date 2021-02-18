package serverpackets

import "l2gogameserver/gameserver/models"

func NewSetupGauge(client *models.Client) {
	client.Buffer.WriteH(0) // reserve
	client.Buffer.WriteSingleByte(0x6b)
	client.Buffer.WriteD(client.CurrentChar.CharId)
	client.Buffer.WriteD(3) // color 0-blue 1-red 2-cyan 3-green

	client.Buffer.WriteD(4132)
	client.Buffer.WriteD(4132)
	client.SimpleSend(client.Buffer.Bytes(), true)

}
