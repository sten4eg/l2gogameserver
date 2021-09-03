package serverpackets

import "l2gogameserver/gameserver/models"

func TargetUnselected(client *models.Client) {

	x, y, z := client.CurrentChar.GetXYZ()

	client.Buffer.WriteSingleByte(0x24)
	client.Buffer.WriteD(client.CurrentChar.CharId)
	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)
	client.Buffer.WriteD(0)
	client.SaveAndCryptDataInBufferToSend(true)
}
