package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ValidateLocation(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	x, y, z := client.CurrentChar.GetXYZ()

	buffer.WriteSingleByte(0x79)
	buffer.WriteD(client.CurrentChar.ObjectId)
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)
	buffer.WriteD(z)
	return buffer.Bytes()

}
