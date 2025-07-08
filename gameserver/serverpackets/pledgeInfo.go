package serverpackets

import (
	"l2gogameserver/packets"
)

func PledgeInfo() []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x89)
	buffer.WriteD(0)
	buffer.WriteS("")
	buffer.WriteS("")

	return buffer.Bytes()
}
