package serverpackets

import "l2gogameserver/packets"

func NewSkillList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x5F)
	buffer.WriteD(1)

	buffer.WriteD(0)
	buffer.WriteD(1)
	buffer.WriteD(1216)
	buffer.WriteD(0)
	buffer.WriteD(0)

	return buffer.Bytes()

}
