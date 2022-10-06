package serverpackets

import "l2gogameserver/packets"

func JoinParty(response int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x3a)
	buffer.WriteD(response)

	return buffer
}
