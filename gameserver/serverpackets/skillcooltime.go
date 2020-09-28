package serverpackets

import (
	"l2gogameserver/packets"
)

func NewSkillCoolTime() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xC7)
	buffer.WriteD(0)
	return buffer.Bytes()
}
