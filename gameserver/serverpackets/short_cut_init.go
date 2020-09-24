package serverpackets

import "l2gogameserver/packets"

func NewShortCutInit() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x45)

	buffer.WriteD(0)

	return buffer.Bytes()
}
