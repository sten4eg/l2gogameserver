package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func Attack(clientI interfaces.ReciverAndSender, targetObjId, targetX, targetY, targetZ int32) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return []byte{}
	}

	x, y, z := client.CurrentChar.GetXYZ()

	buffer.WriteSingleByte(0x33)

	buffer.WriteD(client.CurrentChar.ObjectId)

	buffer.WriteD(targetObjId)
	buffer.WriteD(4)
	buffer.WriteD(0)

	buffer.WriteD(x)
	buffer.WriteD(y)
	buffer.WriteD(z)

	buffer.WriteH(1)
	//for(int i = 1; i < hits.length; i++)
	//{
	//writeD(hits[i]._targetId);
	//writeD(hits[i]._damage);
	//writeC(hits[i]._flags);
	//}

	buffer.WriteD(targetObjId)
	buffer.WriteD(4)
	buffer.WriteD(0)

	buffer.WriteD(targetX)
	buffer.WriteD(targetY)
	buffer.WriteD(targetZ)

	return buffer.Bytes()
}
