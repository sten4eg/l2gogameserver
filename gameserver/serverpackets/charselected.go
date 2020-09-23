package serverpackets

import "l2gogameserver/packets"

func NewCharSelected() []byte {

	buffer := new(packets.Buffer)
	//cSdSddddddddddFFdQddddddddddddddddBd

	buffer.WriteSingleByte(0x0b) // 1
	buffer.WriteS("test")        // 11
	buffer.WriteD(268481286)     // 15
	buffer.WriteS("")            //title // 21 2 нуля
	buffer.WriteD(1)             //sessionId //25
	buffer.WriteD(0)             //clanId // 29
	buffer.WriteD(0)             // ? //33
	buffer.WriteD(1)             //sex// 37
	buffer.WriteD(0)             //race 41
	buffer.WriteD(0)             //classId 45
	buffer.WriteD(1)             // ? 49

	buffer.WriteD(82744)  //x 53
	buffer.WriteD(148536) //y 57
	buffer.WriteD(3400)   //z 61

	buffer.WriteF(126.0) //currentHP 69
	buffer.WriteF(38.0)  //currentMP 77

	buffer.WriteD(0) // SP 81
	buffer.WriteQ(0) // EXP 89
	buffer.WriteD(1) // level 93
	buffer.WriteD(0) // karma 97
	buffer.WriteD(0) // pk 101

	buffer.WriteD(45) //int 105
	buffer.WriteD(45) //str 109
	buffer.WriteD(45) //con 113
	buffer.WriteD(45) //men 117
	buffer.WriteD(45) //dex 121
	buffer.WriteD(45) //wit 125

	buffer.WriteD(0) //inGameTime 129
	buffer.WriteD(0) // 133 //tODO

	buffer.WriteD(0) // 137

	buffer.WriteD(0) // 141
	buffer.WriteD(0) // 145
	buffer.WriteD(0) // 149
	buffer.WriteD(0) // 153

	m := make([]byte, 64)
	buffer.WriteSlice(m) //217

	buffer.WriteD(20000) // 221

	return buffer.Bytes()

}
