package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func TeleportToLocation(client *models.Client, targetObjId, x, y, z, heading int32) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x22)

	buffer.WriteD(client.CurrentChar.ObjectId)

	buffer.WriteD(targetObjId)
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)
	buffer.WriteD(0x00) // isValidation ??
	buffer.WriteD(heading)

	client.CurrentChar.SetXYZ(x, y, z)
	return buffer.Bytes()
}
