package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func SocialAction(clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.Client)
	if !ok {
		return []byte{}
	}

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x27)
	buffer.WriteD(client.CurrentChar.ObjectId)
	buffer.WriteD(3)

	return buffer.Bytes()
}
