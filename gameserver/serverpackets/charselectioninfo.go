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

	buffer.WriteD(1)    // objId
	buffer.WriteS("12") // loginName

	buffer.WriteD(1) //sessionId
	buffer.WriteD(0) //clanId
	buffer.WriteD(0) // Builder Level

	buffer.WriteD(1) //sex
	buffer.WriteD(0) // race
	buffer.WriteD(1) // baseclass

	buffer.WriteD(0) // active ??

	buffer.WriteD(-71549) //x 53
	buffer.WriteD(258198) //y 57
	buffer.WriteD(-3108)  //z 61

	buffer.WriteF(444.0)  //currentHP
	buffer.WriteF(1000.0) //currentMP

	buffer.WriteD(0)      // SP
	buffer.WriteQ(500)    // EXP
	buffer.WriteF(0.1234) // percent ?
	buffer.WriteD(3)      // level

	buffer.WriteD(0) // karma
	buffer.WriteD(1) // pk
	buffer.WriteD(1) //pvp

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

	buffer.WriteD(0) //hairStyle
	buffer.WriteD(1) //hairColor
	buffer.WriteD(0) // face

	buffer.WriteF(999)  //max hp
	buffer.WriteF(1000) // max mp

	buffer.WriteD(0) // days left before
	buffer.WriteD(1) //classId

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
