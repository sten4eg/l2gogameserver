package serverpackets

import "l2gogameserver/packets"

func ExNevitAdventTimeChange() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xE1)
	buffer.WriteSingleByte(0)
	buffer.WriteD(1)
	return buffer.Bytes()
}
