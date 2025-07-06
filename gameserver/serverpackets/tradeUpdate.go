package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func TradeUpdate(character interfaces.CharacterI, item interfaces.TradableItemInterface) []byte {
	buffer := packets.Get()

	count := character.GetInventory().GetItemByObjectId(item.GetObjectId()).GetCount() - item.GetCount()
	a := int16(2)
	if count > 0 && item.GetBaseItem().IsStackable() {
		a = 3
	}

	buffer.WriteSingleByte(0x81)
	buffer.WriteH(1)
	buffer.WriteH(a)
	buffer.WriteH(int16(item.GetItemType1()))
	buffer.WriteD(item.GetObjectId())
	buffer.WriteD(item.GetBaseItem().GetId())
	buffer.WriteQ(item.GetCount())
	buffer.WriteSingleByte(byte(item.GetItemType2()))
	buffer.WriteSingleByte(byte(item.GetItemType1()))
	buffer.WriteQ(int64(item.GetBodyPart()))
	buffer.WriteH(item.GetEnchant())
	buffer.WriteH(0x00)
	buffer.WriteH((item.GetItemType2()))
	item.WriteItem(buffer)


	return buffer.Bytes()
}