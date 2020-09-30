package serverpackets

import "l2gogameserver/gameserver/models"

func NewPledgeInfo(client *models.Client) {

	client.Buffer.WriteH(0) //reserve
	client.Buffer.WriteSingleByte(0x89)
	client.Buffer.WriteD(0)
	client.Buffer.WriteS("")
	client.Buffer.WriteS("")

}
