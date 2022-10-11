package serverpackets

import "l2gogameserver/packets"

func CharDeleteFail(reason int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x1E)
	buffer.WriteD(reason)

	return buffer
}
