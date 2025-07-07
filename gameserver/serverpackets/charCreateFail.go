package serverpackets

import (
	"l2gogameserver/packets"
)

func CharCreateFail(reason int32) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x10)
	buffer.WriteD(reason)

	return buffer.Bytes()
}
