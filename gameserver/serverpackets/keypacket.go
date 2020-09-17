package serverpackets

import (
	"l2gogameserver/packets"
)

var StaticBlowfish = []byte{
	0x6b,
	0x60,
	0xcb,
	0x5b,
	0x82,
	0xce,
	0x90,
	0xb1,
	0xcc,
	0x2b,
	0x6c,
	0x55,
	0x6c,
	0x6c,
	0x6c,
	0x6c,
}

func NewKeyPacket() []byte {
	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x2e)
	buffer.WriteSingleByte(1) // protocolOk
	sk := StaticBlowfish

	for i := 0; i < 8; i++ {
		buffer.WriteSingleByte(sk[i])
	}
	buffer.WriteD(0x01)
	buffer.WriteD(0x01) // server id
	buffer.WriteSingleByte(0x01)
	buffer.WriteD(0x00)
	return buffer.Bytes()
}
