package clientpackets

import (
	"l2gogameserver/packets"
)

func NewCharSelected(data []byte) int32 {

	var buffer = packets.NewReader(data)
	charSlot := buffer.ReadInt32()
	_ = buffer.ReadUInt16()
	_ = buffer.ReadInt32()
	_ = buffer.ReadInt32()
	_ = buffer.ReadInt32()

	return charSlot
}
