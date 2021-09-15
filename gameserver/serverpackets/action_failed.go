package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

//TODO нужен ли тут буффер ?
func ActionFailed(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x1f)
	return buffer.Bytes()
}
