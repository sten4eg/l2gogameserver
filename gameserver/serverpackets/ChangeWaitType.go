package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func ChangeWaitType(client *models.Client) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	locx, locy, locz := client.CurrentChar.GetXYZ()

	buffer.WriteSingleByte(0x29)
	buffer.WriteD(client.CurrentChar.CharId)
	buffer.WriteD(client.CurrentChar.SetSitStandPose())
	buffer.WriteD(locx)
	buffer.WriteD(locy)
	buffer.WriteD(locz)

	return buffer.Bytes()
}
