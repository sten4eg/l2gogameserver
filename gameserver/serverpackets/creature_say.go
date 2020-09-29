package serverpackets

import (
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/packets"
)

func NewCreatureSay(say *clientpackets.Say) []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x4a)
	buffer.WriteD(5)
	buffer.WriteD(say.Type)

	name := []byte{116, 0, 101, 0, 115, 0, 116, 0}
	buffer.WriteS(string(name))

	buffer.WriteD(-1) // High Five NPCString ID
	buffer.WriteS(say.Text)
	return buffer.Bytes()
}
