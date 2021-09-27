package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

//НЕ РАБОТАЕТ :(
func Logout(client *models.Client) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x03)
	buffer.WriteS(client.CurrentChar.CharName)

	return buffer.Bytes()
}
