package serverpackets

import (
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/packets"
)

func TradeOtherAdd(item interfaces.TradableItemInterface) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x1B)
	buffer.WriteH(1) // static?item added count
	buffer.WriteH(0) // static?
	buffer.WriteD(item.GetObjectId())
	buffer.WriteD(item.GetId())
	buffer.WriteQ(item.GetCount())
	buffer.WriteH(int16(item.GetItemType2()))
	buffer.WriteH(0)                  // GetCustomType1 ?
	buffer.WriteD(item.GetBodyPart()) // bodypart
	buffer.WriteH(int16(item.GetEnchant()))
	buffer.WriteH(0x00) // ?
	buffer.WriteH(0)    //GetCustomType2 ??

	buffer.WriteH(int16(item.GetAttackElementType()))
	buffer.WriteH(int16(item.GetAttackElementPower()))

	for _, v := range item.GetElementDefAttr() {
		buffer.WriteH(v)
	}
	for _, v := range item.GetEnchantedOption() {
		buffer.WriteH(int16(v))
	}
	return buffer.Bytes()
}
