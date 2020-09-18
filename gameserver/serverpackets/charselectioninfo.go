package serverpackets

import "l2gogameserver/packets"

func NewCharSelectionInfo() []byte {

	buffer := new(packets.Buffer)
	buffer.WriteSingleByte(0x09)
	buffer.WriteD(0) //size

	buffer.WriteD(0) // Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	buffer.WriteSingleByte(0)

	return buffer.Bytes()
}
