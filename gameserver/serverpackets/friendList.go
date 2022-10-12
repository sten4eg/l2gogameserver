package serverpackets

import "l2gogameserver/packets"

func FriendList() []byte {

	buffer := packets.Get()

	buffer.WriteSingleByte(0x75)
	buffer.WriteD(0)

	return buffer.Bytes()
}
