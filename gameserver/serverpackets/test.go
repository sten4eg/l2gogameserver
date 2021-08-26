package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"log"
)

//MagicSkillLaunched
func NewTest(client *models.Client) {

	client.Buffer.WriteSingleByte(0x54)
	client.Buffer.WriteD(client.CurrentChar.CharId)
	client.Buffer.WriteD(1216)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(1)
	client.Buffer.WriteD(client.CurrentChar.CharId)
	log.Println("TESST")
	client.SaveAndCryptDataInBufferToSend(true)
	client.SentToSend()

}
