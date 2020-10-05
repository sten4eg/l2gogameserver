package serverpackets

import (
	"l2gogameserver/gameserver/models"
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

func NewKeyPacket(client *models.Client) {

	client.Buffer.WriteH(0) //reserve
	client.Buffer.WriteSingleByte(0x2e)
	client.Buffer.WriteSingleByte(1) // protocolOk
	sk := StaticBlowfish

	for i := 0; i < 8; i++ {
		client.Buffer.WriteSingleByte(sk[i])
	}
	client.Buffer.WriteD(0x01)
	client.Buffer.WriteD(0x01) // server id
	client.Buffer.WriteSingleByte(0x01)
	client.Buffer.WriteD(0x00)

	client.SimpleSend(client.Buffer.Bytes(), true)

}
