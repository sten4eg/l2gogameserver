package serverpackets

import (
	"l2gogameserver/packets"
)

func TradeOtherDone() []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x82)
	return buffer.Bytes()
}
