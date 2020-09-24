package serverpackets

import "l2gogameserver/packets"

func NewQuestList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x86)
	buffer.WriteH(0)
	x := make([]byte, 128, 128)
	buffer.WriteSlice(x)

	return buffer.Bytes()
}
