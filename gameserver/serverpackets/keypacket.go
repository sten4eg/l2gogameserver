package serverpackets

import (
	"l2gogameserver/gameserver/models"
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
	200,
	39,
	147,
	1,
	161,
	108,
	49,
	151,
}

func KeyPacket(client *models.Client) []byte {
	//TODO Данные с этого пакета НЕ надо шифровать
	buffer := packets.Get()

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

	defer packets.Put(buffer)
	return buffer.Bytes()
}
