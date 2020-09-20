package serverpackets

import (
	"l2gogameserver/packets"
)

func NewCharSelectionInfo() []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x09)
	buffer.WriteD(1) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	buffer.WriteD(7)          //char max number
	buffer.WriteSingleByte(0) // delim

	//todo блок который должен повторяться

	buffer.WriteS("test") // Pers name
	buffer.WriteSingleByte(0)
	buffer.WriteSingleByte(0)

	buffer.WriteD(1)    // objId
	buffer.WriteS("12") // loginName
	buffer.WriteSingleByte(0)
	buffer.WriteSingleByte(0)

	buffer.WriteD(12) //sessionId
	buffer.WriteD(0)  //clanId
	buffer.WriteD(0)  // Builder Level

	buffer.WriteD(1) //sex
	buffer.WriteD(0) // race
	buffer.WriteD(1) // baseclass

	buffer.WriteD(1) // active ??

	buffer.WriteD(82744)  //x
	buffer.WriteD(148536) //y
	buffer.WriteD(3434)   //z

	buffer.WriteF(1000) //currentHP
	buffer.WriteF(1000) //currentMP

	buffer.WriteD(0)  // SP
	buffer.WriteQ(0)  // EXP
	buffer.WriteF(10) // percent ?
	buffer.WriteD(1)  // level

	buffer.WriteD(0) // karma
	buffer.WriteD(0) // pk
	buffer.WriteD(0) //pvp

	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	buffer.WriteD(0)
	//
	//
	m := make([]byte, 104)
	buffer.WriteSlice(m)

	buffer.WriteD(1) //hairStyle
	buffer.WriteD(1) //hairColor
	buffer.WriteD(1) // face

	buffer.WriteF(1000) //max hp
	buffer.WriteF(1000) // max mp

	buffer.WriteD(0)  // days left before
	buffer.WriteD(25) //classId

	buffer.WriteD(1)          //auto-selected
	buffer.WriteSingleByte(0) // enchanted
	buffer.WriteD(0)          //augumented

	buffer.WriteD(0) // Currently on retail when you are on character select you don't see your transformation.

	// Implementing it will be waster of resources.
	buffer.WriteD(0) // Pet ID
	buffer.WriteD(0) // Pet Level
	buffer.WriteD(0) // Pet Max Food
	buffer.WriteD(0) // Pet Current Food
	buffer.WriteF(0) // Pet Max HP
	buffer.WriteF(0) // Pet Max MP
	buffer.WriteD(0) // H5 Vitality
	return buffer.Bytes()
}
