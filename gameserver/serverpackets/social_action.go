package serverpackets

import "l2gogameserver/gameserver/models"

func NewSocialAction(client *models.Client) {

	client.Buffer.WriteH(0) //reserve
	client.Buffer.WriteSingleByte(0x27)
	client.Buffer.WriteD(client.CurrentChar.CharId)
	client.Buffer.WriteD(3)

}
