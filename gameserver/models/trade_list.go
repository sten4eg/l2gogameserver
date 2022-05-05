package models

import (
	"l2gogameserver/config"
	"l2gogameserver/gameserver/interfaces"
)

type TradeList struct {
	owner     interfaces.CharacterI
	partner   interfaces.CharacterI
	items     []interfaces.TradableItemInterface
	title     string
	packaged  bool
	confirmed bool
	locked    bool
}

func NewTradeList(owner interfaces.CharacterI) *TradeList {
	t := new(TradeList)
	t.owner = owner
	return t
}

func (t *TradeList) SetPartner(partner interfaces.CharacterI) {
	t.partner = partner
}

func (t *TradeList) GetPartner() interfaces.CharacterI {
	return t.partner
}

func (t *TradeList) GetOwner() interfaces.CharacterI {
	return t.owner
}

func (t *TradeList) Lock() {
	t.locked = true
}
func (t *TradeList) IsLocked() bool {
	return t.locked
}
func (t *TradeList) IsConfirmed() bool {
	return t.confirmed
}
func (t *TradeList) InvalidateConfirmation() {
	t.confirmed = false
}
func (t *TradeList) Confirmed() bool {
	if t.confirmed {
		return true
	}

	partner := t.GetPartner()
	if partner != nil {
		//TOdo тут много проверок
		t.confirmed = true
		///
	} else {
		t.confirmed = true
	}
	return t.confirmed
}
func (t *TradeList) AddItem(objectId int32, count int64, char interfaces.CharacterI, price int64) interfaces.TradableItemInterface {
	if t.IsLocked() {
		return nil
	}
	item := char.GetInventory().GetItemByObjectId(objectId)
	if item == nil {
		return nil
	}

	if !t.GetOwner().GetInventory().CanManipulateWithItemId(int32(item.GetId())) {
		return nil
	}

	if count <= 0 || count > item.GetCount() {
		return nil
	}

	if !item.IsStackable() && count > 1 {
		return nil
	}

	if (config.MaxAdena / count) < price {
		return nil
	}

	for i := range t.items {
		if t.items[i].GetObjectId() == objectId {
			return nil
		}
	}
	r := NewTradeItem(item, count, price)
	t.InvalidateConfirmation()
	return r
}
