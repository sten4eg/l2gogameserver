package serverpackets

import (
	"l2gogameserver/packets"
)

func SkillCoolTime() []byte {

	buffer := packets.Get()

	buffer.WriteSingleByte(0xC7)
	buffer.WriteD(0)

	//buffer.WriteD(1216)
	//buffer.WriteD(1)
	//buffer.WriteD(0)
	//buffer.WriteD(0)
	//
	//buffer.WriteD(1184)
	//buffer.WriteD(1)
	//buffer.WriteD(0)
	//buffer.WriteD(0)

	return buffer.Bytes()
}
