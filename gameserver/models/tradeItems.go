package models

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/attribute"
)

type TradeItem struct {
	items.Item
	ObjectId            int32
	Enchant             int16
	LocData             int32
	Price               int64
	Location            string
	Type1               int
	Type2               int16
	AttackAttributeType attribute.Attribute
	AttackAttributeVal  int16
	AttributeDefend     [6]int16
	StoreCount          int64
	Count               int64
	EnchantedOption     [3]int32
}

func NewTradeItem(item interfaces.MyItemInterface, count, price int64) interfaces.TradableItemInterface {
	baseItemInterface := item.GetBaseItem()
	baseItem, ok := baseItemInterface.(*items.Item)
	if !ok {
		logger.Error.Panicln("NewTradeItem baseItemInterface is not item.Item")
	}
	t := TradeItem{
		Item:                *baseItem,
		ObjectId:            item.GetObjectId(),
		Location:            item.GetLocation(),
		Enchant:             item.GetEnchant(),
		Type1:               item.GetItemType1(),
		Type2:               item.GetItemType2(),
		Count:               count,
		Price:               price,
		AttackAttributeType: item.GetAttackElementType(),
		AttackAttributeVal:  item.GetAttackElementPower(),
		AttributeDefend:     item.GetElementDefAttr(),
		EnchantedOption:     item.GetEnchantedOption(),
	}
	return &t
}

func (i *TradeItem) GetObjectId() int32 {
	return i.ObjectId
}
func (i *TradeItem) GetItemType() int {
	return int(i.SlotBitType)
}
func (i *TradeItem) GetBodyPart() int32 {
	return 0
}
func (i *TradeItem) GetItemType1() int {
	return int(i.ItemType1)
}
func (i *TradeItem) GetItemType2() int16 {
	return int16(i.ItemType2)
}
func (i *TradeItem) GetAttackElementType() attribute.Attribute {
	el := attribute.Attribute(-2) // none
	if i.IsWeapon() {
		el = i.AttackAttributeType
	}

	if el == attribute.None {
		if i.BaseAttributeAttack.Val > 0 {
			return i.getBaseAttributeElement()
		}
	}

	return el
}
func (i *TradeItem) getBaseAttributeElement() attribute.Attribute {
	return i.BaseAttributeAttack.Type
}
func (i *TradeItem) GetEnchant() int16 {
	return i.Enchant
}
func (i *TradeItem) GetAttackElementPower() int16 {
	return i.AttackAttributeVal
}

func (i *TradeItem) GetElementDefAttr() [6]int16 {
	return i.AttributeDefend
}
func (i *TradeItem) GetEnchantedOption() [3]int32 {
	return i.EnchantedOption
}
func (i *TradeItem) GetCount() int64 {
	return i.Count
}
