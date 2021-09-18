package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

// SetupGauge полоска над персонажем во время каста скила
func SetupGauge(client *models.Client) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x6b)
	buffer.WriteD(client.CurrentChar.CharId)
	buffer.WriteD(0) // color 0-blue 1-red 2-cyan 3-green

	buffer.WriteD(4132)
	buffer.WriteD(4132)

	return buffer.Bytes()

}
