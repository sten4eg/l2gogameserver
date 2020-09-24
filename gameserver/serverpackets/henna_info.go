package serverpackets

import "l2gogameserver/packets"

func NewHennaInfo() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xE5)

	buffer.WriteSingleByte(0) //int
	buffer.WriteSingleByte(0) //int
	buffer.WriteSingleByte(0) //int
	buffer.WriteSingleByte(0) //int
	buffer.WriteSingleByte(0) //int
	buffer.WriteSingleByte(0) //int
	buffer.WriteD(3)          //slots
	buffer.WriteD(0)

	return buffer.Bytes()
}
