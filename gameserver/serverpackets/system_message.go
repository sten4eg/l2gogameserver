package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"l2gogameserver/packets"
)

func SystemMessage(msg sysmsg.SysMsg, client interfaces.ReciverAndSender) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x62) // 062 для всех сис мессаджей
	buffer.WriteD(msg.Id)
	buffer.WriteD(1) //params.len

	return buffer.Bytes()
	//buffer.WriteD(0)
	//buffer.WriteS("fuck")
	//return buffer.Bytes()
}
