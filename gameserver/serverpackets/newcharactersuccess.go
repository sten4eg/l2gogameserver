package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func CharacterSuccess(client *models.Client) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x0D)
	buffer.WriteD(1)
	buffer.WriteD(1)
	buffer.WriteD(1)
	buffer.WriteD(0x46)
	buffer.WriteD(1)
	buffer.WriteD(0x0A)
	buffer.WriteD(0x46)
	buffer.WriteD(1)
	buffer.WriteD(0x0A)
	buffer.WriteD(0x46)
	buffer.WriteD(1)
	buffer.WriteD(0x0A)
	buffer.WriteD(0x46)
	buffer.WriteD(1)
	buffer.WriteD(0x0A)
	buffer.WriteD(0x46)
	buffer.WriteD(1)
	buffer.WriteD(0x0A)
	buffer.WriteD(0x46)
	buffer.WriteD(1)
	buffer.WriteD(0x0A)

	defer packets.Put(buffer)
	return buffer.Bytes()
}
