package serverpackets

import (
	"l2gogameserver/packets"
)

func TradeOtherDone() []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x82)
	return buffer.Bytes()
}
