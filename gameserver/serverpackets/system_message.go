package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/sysmsg"
)

func NewSystemMessage(msg sysmsg.SysMsg, client *models.Client) {

	client.Buffer.WriteSingleByte(0x62) // 062 для всех сис мессаджей
	client.Buffer.WriteD(msg.Id)
	client.Buffer.WriteD(1) //params.len

	client.SaveAndCryptDataInBufferToSend(true)
	//buffer.WriteD(0)
	//buffer.WriteS("fuck")
	//return buffer.Bytes()
}
