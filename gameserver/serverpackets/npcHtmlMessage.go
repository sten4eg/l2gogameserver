package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
	"strconv"
)

func NpcHtmlMessage(client interfaces.ReciverAndSender, npcid int32) []byte {

	buffer := packets.Get()
	defer packets.Put(buffer)

	buffer.WriteSingleByte(0x19)

	buffer.WriteD(33)
	buffer.WriteS("<html><title>Info</title><body>\n<center>\n" + strconv.Itoa(int(npcid)) + "<br1>\n</center>\n</body></html>")
	buffer.WriteD(0)

	return buffer.Bytes()
}

func NpcHtmlMessage2(npcObjId int32, html string, itemId int32) *packets.Buffer {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x19)

	buffer.WriteD(npcObjId)
	buffer.WriteS(html)
	buffer.WriteD(itemId)

	return buffer
}
