package serverpackets

import "l2gogameserver/packets"

func NewObservationReturn(user *User) []byte {
	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0xEC)
	buffer.WriteD(user.X) //x 53
	buffer.WriteD(user.Y) //y 57
	buffer.WriteD(user.Z) //z 61

	return buffer.Bytes()
}
