package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
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

func KeyPacket(client interfaces.ReciverAndSender) []byte {
	//TODO Данные с этого пакета НЕ надо шифровать
	buffer := packets.Get()

	buffer.WriteSingleByte(0x2e)
	buffer.WriteSingleByte(1) // protocolOk
	sk := StaticBlowfish

	for i := 0; i < 8; i++ {
		buffer.WriteSingleByte(sk[i])
	}
	buffer.WriteD(0x01)
	buffer.WriteD(0x01) // datapack id
	buffer.WriteSingleByte(0x01)
	buffer.WriteD(0x00)

	return buffer.Bytes()
}
