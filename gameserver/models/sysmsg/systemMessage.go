package sysmsg

import (
	"l2gogameserver/packets"
)

func SystemMessage(msg SysMsg) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x62)
	buffer.WriteD(msg.Id)
	buffer.WriteD(int32(len(msg.Params)))

	for _, v := range msg.Params {
		buffer.WriteD(int32(v.GetType()))
		switch v.GetType() {
		case TypeText, TypePlayerName:
			buffer.WriteS(v.GetValueString())
		case TypeLongNumber:
			buffer.WriteQ(v.GetValueInt64())
		case TypeItemName, TypeCastleName, TypeIntNumber, TypeNpcName, TypeElementName, TypeSystemString, TypeInstanceName, TypeDoorName:
			buffer.WriteD(v.GetValueInt32())
		case TypeSkillName:
			p := v.GetTwoElementSlice()
			buffer.WriteD(p[0]) // SkillId
			buffer.WriteD(p[1]) // SkillLevel
		case TypeZoneName:
			p := v.GetThreeElementSlice()
			buffer.WriteD(p[0]) // x
			buffer.WriteD(p[1]) // y
			buffer.WriteD(p[2]) // z
		}
	}
	return buffer.Bytes()
}

func SendCustomSystemMessage(text string) []byte {
	if text == "" {
		return []byte{}
	}
	buffer := packets.Get()

	buffer.WriteSingleByte(0x62)
	buffer.WriteD(S1.Id)
	buffer.WriteD(1)

	buffer.WriteD(int32(TypeText))
	buffer.WriteS(text)
	return buffer.Bytes()
}
