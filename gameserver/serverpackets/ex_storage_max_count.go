package serverpackets

import "l2gogameserver/packets"

func NewExStorageMaxCount() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0x2F)

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
