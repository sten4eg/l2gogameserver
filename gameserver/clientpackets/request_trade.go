package clientpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func TradeRequest(data []byte, client *models.Client) int32 {
	var packet = packets.NewReader(data)
	objId := packet.ReadInt32() // targetObjId
	return objId
}
