package serverpackets

import (
	"l2gogameserver/packets"
)

func NpcSay(objectId int32, textType int32, npcId int32, npcString string) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x30)
	buffer.WriteD(objectId)
	buffer.WriteD(textType)
	buffer.WriteD(npcId)
	buffer.WriteD(0) // -1 тогда свой текст
	buffer.WriteS(npcString)
	return buffer.Bytes()
}
