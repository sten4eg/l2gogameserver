package serverpackets

import (
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func TradeOtherAdd(client *models.Client, item *models.MyItem, count uint64) []byte {
	buffer := packets.Get()
	defer packets.Put(buffer)
	buffer.WriteSingleByte(0x1B)
	buffer.WriteH(1) // item count
	buffer.WriteH(0)
	buffer.WriteD(item.ObjId)
	buffer.WriteD(int32(item.Id))
	buffer.WriteQ(int64(count))
	buffer.WriteH(int16(item.ItemType))
	buffer.WriteH(0)
	buffer.WriteD(int32(item.SlotBitType)) // bodypart
	buffer.WriteH(int16(item.Enchant))
	buffer.WriteH(0x00)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	buffer.WriteH(0)
	return buffer.Bytes()
}
