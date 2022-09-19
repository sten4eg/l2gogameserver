package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func TargetUnselected(clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return []byte{}
	}

	client.CurrentChar.Target = 0
	buffer := packets.Get()
	defer packets.Put(buffer)

	x, y, z := client.CurrentChar.GetXYZ()

	buffer.WriteSingleByte(0x24)
	buffer.WriteD(client.CurrentChar.ObjectId)
	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)
	buffer.WriteD(0)

	return buffer.Bytes()
}
