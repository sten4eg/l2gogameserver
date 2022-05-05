package serverpackets

import (
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/packets"
)

func SystemMessage(msg sysmsg.SysMsg) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x62)
	buffer.WriteD(msg.Id)
	buffer.WriteD(int32(len(msg.Params)))

	for _, v := range msg.Params {
		buffer.WriteD(int32(v.GetType()))
		switch v.GetType() {
		case sysmsg.TypeText, sysmsg.TypePlayerName:
			buffer.WriteS(v.GetValueString())
		case sysmsg.TypeLongNumber:
			buffer.WriteQ(v.GetValueInt64())
		case sysmsg.TypeItemName, sysmsg.TypeCastleName, sysmsg.TypeIntNumber, sysmsg.TypeNpcName, sysmsg.TypeElementName, sysmsg.TypeSystemString, sysmsg.TypeInstanceName, sysmsg.TypeDoorName:
			buffer.WriteD(v.GetValueInt32())
		case sysmsg.TypeSkillName:
			p := v.GetTwoElementSlice()
			buffer.WriteD(p[0]) // SkillId
			buffer.WriteD(p[1]) // SkillLevel
		case sysmsg.TypeZoneName:
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
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x62)
	buffer.WriteD(sysmsg.S1.Id)
	buffer.WriteD(1)

	buffer.WriteD(int32(sysmsg.TypeText))
	buffer.WriteS(text)
	return buffer.Bytes()
}
