package clientpackets

import (
	"l2gogameserver/packets"
)

func NewCharSelected(data []byte) {

	var buffer = packets.NewReader(data)
	_ = buffer.ReadUInt32()
	_ = buffer.ReadUInt16()
	_ = buffer.ReadUInt32()
	_ = buffer.ReadUInt32()
	_ = buffer.ReadUInt32()

}
