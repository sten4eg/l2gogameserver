package serverpackets

import "l2gogameserver/packets"

func NewObservationReturn() []byte {
	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xEC)
	buffer.WriteD(83306)  //x 53
	buffer.WriteD(148115) //y 57
	buffer.WriteD(-3405)  //z 61

	return buffer.Bytes()
}
