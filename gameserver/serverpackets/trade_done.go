package serverpackets

import (
	"l2gogameserver/packets"
)

func TradeDone(num int32) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x1c)
	buffer.WriteD(num)

	return buffer.Bytes()
}