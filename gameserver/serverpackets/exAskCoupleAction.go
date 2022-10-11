package serverpackets

import "l2gogameserver/packets"

func ExAskCoupleAction(objId, actionId int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0xFE)
	buffer.WriteH(0xBB)
	buffer.WriteD(actionId)
	buffer.WriteD(objId)

	return buffer
}
