package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewHennaInfo(client *models.Client) {

	client.Buffer.WriteSingleByte(0xE5)

	client.Buffer.WriteSingleByte(0) // equip INT
	client.Buffer.WriteSingleByte(0) // equip STR
	client.Buffer.WriteSingleByte(0) //equip CON
	client.Buffer.WriteSingleByte(0) // equip MEN
	client.Buffer.WriteSingleByte(0) //equip DEX
	client.Buffer.WriteSingleByte(0) //equip WIT
	client.Buffer.WriteD(3)          //slots
	client.Buffer.WriteD(0)          //Size

	client.SaveAndCryptDataInBufferToSend(true)
}
