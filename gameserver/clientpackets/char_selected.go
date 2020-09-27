package clientpackets

import (
	"l2gogameserver/packets"
)

func NewCharSelected(data []byte) {

	var buffer = packets.NewReader(data)
	_ = buffer.ReadInt32()
	_ = buffer.ReadUInt16()
	_ = buffer.ReadInt32()
	_ = buffer.ReadInt32()
	_ = buffer.ReadInt32()

}
