package clientpackets

import "l2gogameserver/packets"

func NewAction(data []byte) {
	var packet = packets.NewReader(data)

	objectId := packet.ReadInt32() //Target
	originX := packet.ReadInt32()
	originY := packet.ReadInt32()
	originZ := packet.ReadInt32()
	actionId := packet.ReadSingleByte() // Action identifier : 0-Simple click, 1-Shift click

	_, _, _, _, _ = objectId, originX, originY, originZ, actionId
	var i bool
	_ = i

}
