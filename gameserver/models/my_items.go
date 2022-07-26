package models

import (
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/items/attribute"
)

type MyItem struct {
	// встроенный "шаблон" предмета
	items.Item
	ObjectId            int32
	Enchant             int16
	LocData             int32
	Count               int64
	Location            string
	Time                int
	AttackAttributeType attribute.Attribute
	AttackAttributeVal  int16
	Mana                int32
	AttributeDefend     [6]int16
	EnchantedOption     [3]int32

	//UpdateType для обновления инвентаря
	LastChange int16
}

func (i *MyItem) GetObjectId() int32 {
	return i.ObjectId
}
func (i *MyItem) IsEquipped() int16 {
	if i.Location == InventoryLoc {
		return 0
	}
	return 1
}
func (i *MyItem) GetAttackElementType() attribute.Attribute {
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
func (i *MyItem) getBaseAttributeElement() attribute.Attribute {
	return i.BaseAttributeAttack.Type
}
func (i *MyItem) GetCount() int64 {
	return i.Count
}
func (i *MyItem) GetEnchant() int16 {
	return i.Enchant
}
func (i *MyItem) GetAttackElementPower() int16 {
	return i.AttackAttributeVal
}
func (i *MyItem) GetElementDefAttr() [6]int16 {
	return i.AttributeDefend
}
func (i *MyItem) GetEnchantedOption() [3]int32 {
	return i.EnchantedOption
}
func (i *MyItem) GetLocation() string {
	return i.Location
}

func (i *MyItem) GetUpdateType() int16 {
	return i.LastChange
}
func (i *MyItem) GetLocData() int32 {
	return i.LocData
}
func (i *MyItem) GetMana() int32 {
	return i.Mana
}
