package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ObservationReturn(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	x, y, z := client.CurrentChar.GetXYZ()

	buffer.WriteSingleByte(0xEC)
	buffer.WriteD(x) //x 53
	buffer.WriteD(y) //y 57
	buffer.WriteD(z) //z 61

	return buffer.Bytes()
}
