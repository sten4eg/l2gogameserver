package serverpackets

import (
	"l2gogameserver/packets"
)

//func DeleteObject(character *models.Character) []byte {
//	buffer := packets.Get()
//	defer packets.Put(buffer)
//
//	buffer.WriteSingleByte(0x08)
//	buffer.WriteD(character.ObjectId)
//	buffer.WriteD(0)
//	return buffer.Bytes()
//}

//func DeleteObject(objectId int32) []byte {
//	buffer := packets.Get()
//	defer packets.Put(buffer)
//
//	buffer.WriteSingleByte(0x08)
//	buffer.WriteD(objectId)
//	buffer.WriteD(0)
//	return buffer.Bytes()
//}

func DeleteObject(objectId int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x08)
	buffer.WriteD(objectId)
	buffer.WriteD(0x00)

	return buffer
}
