package serverpackets

import (
	"l2gogameserver/gameserver/models"
)

func NewShortCutInit(client *models.Client) {

	client.Buffer.WriteSingleByte(0x45)

	client.Buffer.WriteD(0)
	//todo уже есть шорткат надо доделать цэй инит)
	client.SaveAndCryptDataInBufferToSend(true)
}
