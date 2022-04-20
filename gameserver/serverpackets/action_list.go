package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func ActionList(client interfaces.ReciverAndSender) []byte {
	const (
		count1 = 74
		count2 = 99
		count3 = 16
	)
	n := count1 + count2 + count3 + 1
	DefaultActionList := make([]int, n, 2*n)
	for i := count1; i > 0; i-- {
		DefaultActionList[i] = i
	}
	for i := count2; i > 0; i-- {
		DefaultActionList[count1+i] = 1000 + i
	}
	for i := count3; i > 0; i-- {
		DefaultActionList[count1+count2+i] = 5000 + i
	}
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0xfe)
	buffer.WriteH(0x5f)
	buffer.WriteD(int32(len(DefaultActionList)))

	for _, index := range DefaultActionList {
		buffer.WriteD(int32(index))
	}

	return buffer.Bytes()
}
