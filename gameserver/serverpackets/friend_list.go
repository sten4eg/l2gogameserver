package serverpackets

import "l2gogameserver/packets"

func NewFriendList() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x75)
	buffer.WriteD(0)

	return buffer.Bytes()
}
