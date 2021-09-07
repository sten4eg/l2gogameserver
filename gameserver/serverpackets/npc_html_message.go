package serverpackets

import "l2gogameserver/gameserver/models"

func NpcHtmlMessage(client *models.Client) {
	client.Buffer.WriteSingleByte(0x19)

	client.Buffer.WriteD(33)
	client.Buffer.WriteS("<html><title>Admin Help</title><body> Privet</body></html>")
	client.Buffer.WriteD(0)
	client.SaveAndCryptDataInBufferToSend(true)
}
