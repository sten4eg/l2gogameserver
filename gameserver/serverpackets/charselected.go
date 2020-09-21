package serverpackets

import "l2gogameserver/packets"

func NewCharSelected() []byte {

	buffer := new(packets.Buffer)

	buffer.WriteSingleByte(0x0b)
	buffer.WriteS("test")
	buffer.WriteD(1)
	buffer.WriteS("hello") //title
	buffer.WriteD(1)       //sessionId
	buffer.WriteD(0)       //clanId
	buffer.WriteD(0)       // ?
	buffer.WriteD(1)       //sex
	buffer.WriteD(0)       //race
	buffer.WriteD(0)       //classId
	buffer.WriteD(1)       // ?

	buffer.WriteD(82744)  //x
	buffer.WriteD(148536) //y
	buffer.WriteD(3400)   //z

	buffer.WriteF(444.0)  //currentHP
	buffer.WriteF(1000.0) //currentMP

	buffer.WriteD(0)   // SP
	buffer.WriteQ(500) // EXP
	buffer.WriteD(3)   // level
	buffer.WriteD(0)   // karma
	buffer.WriteD(1)   // pk

	buffer.WriteD(5) //int
	buffer.WriteD(5) //str
	buffer.WriteD(5) //con
	buffer.WriteD(5) //men
	buffer.WriteD(5) //dex
	buffer.WriteD(5) //wit

	buffer.WriteD(1) //inGameTime
	buffer.WriteD(0) //

	buffer.WriteD(0)

	buffer.WriteD(0) //
	buffer.WriteD(0) //
	buffer.WriteD(0) //
	buffer.WriteD(0) //

	m := make([]byte, 64)
	buffer.WriteSlice(m)

	buffer.WriteD(0)

	return buffer.Bytes()

}
