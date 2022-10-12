package serverpackets

import "l2gogameserver/packets"

func EtcStatusUpdate() []byte {

	buffer := packets.Get()

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
