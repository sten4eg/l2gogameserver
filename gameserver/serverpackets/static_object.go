package serverpackets

import "l2gogameserver/packets"

func NewStaticObject() []byte {
	buffer := new(packets.Buffer)
	buffer.WriteD(0)
	buffer.WriteD(0)
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
