package serverpackets

import "l2gogameserver/gameserver/models"

type Attack struct {
	TargetId int32
	Damage   int32
	X        int32
	Y        int32
	Z        int32
}

func NewAttack(client *models.Client, attack *Attack) {

	x, y, z := client.CurrentChar.GetXYZ()

	client.Buffer.WriteSingleByte(0x33)

	client.Buffer.WriteD(client.CurrentChar.CharId)

	client.Buffer.WriteD(attack.TargetId)
	client.Buffer.WriteD(4)
	client.Buffer.WriteD(0)

	client.Buffer.WriteD(x)
	client.Buffer.WriteD(y)
	client.Buffer.WriteD(z)

	client.Buffer.WriteH(1)
	//for(int i = 1; i < hits.length; i++)
	//{
	//writeD(hits[i]._targetId);
	//writeD(hits[i]._damage);
	//writeC(hits[i]._flags);
	//}

	client.Buffer.WriteD(attack.TargetId)
	client.Buffer.WriteD(4)
	client.Buffer.WriteD(0)

	client.Buffer.WriteD(attack.X)
	client.Buffer.WriteD(attack.Y)
	client.Buffer.WriteD(attack.Z)

	client.SaveAndCryptDataInBufferToSend(true)
}
