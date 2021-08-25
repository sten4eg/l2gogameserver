package clientpackets

import (
	"fmt"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func NewRequestAutoSoulShot(data []byte, client *models.Client) {
	var packet = packets.NewReader(data[2:])
	itemId := packet.ReadInt32()
	typee := packet.ReadInt32()

	client.CurrentChar.ActiveSoulShots = append(client.CurrentChar.ActiveSoulShots, itemId)

	client.Buffer.WriteSingleByte(0xFE)

	client.Buffer.WriteH(0x0c)
	client.Buffer.WriteD(itemId)
	client.Buffer.WriteD(typee)
	client.SaveAndCryptDataInBufferToSend(true)
	fmt.Println(itemId, typee)
}
