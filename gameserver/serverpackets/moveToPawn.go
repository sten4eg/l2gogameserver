package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func MoveToPawn(character interfaces.CharacterI) []byte {
	buffer := packets.Get()

	//writeC(0x72);
	//
	//writeD(_charObjId);
	//writeD(_targetId);
	//writeD(_distance);
	//
	//writeD(_x);
	//writeD(_y);
	//writeD(_z);
	//writeD(_tx);
	//writeD(_ty);
	//writeD(_tz);

	buffer.WriteSingleByte(0x72)
	buffer.WriteD(character.GetObjectId())
	buffer.WriteD(character.GetTarget())
	buffer.WriteD(0)

	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	return buffer.Bytes()

}
