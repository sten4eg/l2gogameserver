package serverpackets

import "l2gogameserver/packets"

func NewSkillList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x5F)
	buffer.WriteD(1) // skill size

	// for
	buffer.WriteD(0)    // passiv ?
	buffer.WriteD(1)    // level
	buffer.WriteD(1216) // id
	buffer.WriteD(0)    // disable?
	buffer.WriteD(0)    // enchant ?
	// endfor
	return buffer.Bytes()

}
