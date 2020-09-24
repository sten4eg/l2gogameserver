package serverpackets

import "l2gogameserver/packets"

func NewActionFailed() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x1f)

	return buffer.Bytes()
}
