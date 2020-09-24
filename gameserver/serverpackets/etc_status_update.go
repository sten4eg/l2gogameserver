package serverpackets

import "l2gogameserver/packets"

func NewEtcStatusUpdate() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xf9)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)

	return buffer.Bytes()
}
