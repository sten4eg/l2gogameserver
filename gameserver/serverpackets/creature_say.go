package serverpackets

import (
	"l2gogameserver/gameserver/clientpackets"
	"l2gogameserver/packets"
)

func NewCreatureSay(say *clientpackets.Say) []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x4a)
	buffer.WriteD(1)
	buffer.WriteD(say.Type)

	buffer.WriteS("test")

	buffer.WriteD(-1) // High Five NPCString ID
	buffer.WriteS(say.Text)
	return buffer.Bytes()
}
