package models

import (
	"l2gogameserver/config"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/sysmsg"
	"log"
	"math"
	"sync"
)

type TradeList struct {
	owner     interfaces.CharacterI
	partner   interfaces.CharacterI
	items     []interfaces.TradableItemInterface
	title     string
	packaged  bool
	confirmed bool
	locked    bool
	mu        sync.Mutex
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
func (t *TradeList) Confirmed() (bool, needSendTradeConfirm bool) {
	needSendTradeConfirm = false
	if t.confirmed {
		return true, needSendTradeConfirm
	}

	partner := t.GetPartner()
	if partner != nil {
		partnerList := partner.GetActiveTradeList()
		if partnerList == nil {
			log.Println(partner.GetName() + ": Trading partner (" + partner.GetName() + ") is invalid in this trade!")
			return false, needSendTradeConfirm
		}

		var sync1, sync2 interfaces.TradeListInterface
		if t.owner.GetObjectId() > partnerList.GetOwner().GetObjectId() {
			sync1 = partnerList
			sync2 = t
		} else {
			sync1 = t
			sync2 = partnerList
		}

		sync1.MuLock()
		defer sync1.MuUnlock()
		sync2.MuLock()
		defer sync2.MuUnlock()
		t.confirmed = true
		if partnerList.IsConfirmed() {
			partnerList.Lock()
			t.Lock()
			if !partnerList.Validate() {
				return false, needSendTradeConfirm
			}
			if !t.Validate() {
				return false, needSendTradeConfirm
			}

			//	doExchange(partnerList)
		} else {
			partner.SendSysMsg(sysmsg.AlreadyTrading)
			needSendTradeConfirm = true
		}

		//TOdo тут много проверок
		t.confirmed = true
		///
	} else {
		t.confirmed = true
	}
	return t.confirmed, needSendTradeConfirm
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

func (t *TradeList) MuLock() {
	t.mu.Lock()
}

func (t *TradeList) MuUnlock() {
	t.mu.Unlock()
}

func (t *TradeList) Validate() bool {
	if t.owner == nil {
		log.Println("Invalid owner of TradeList")
		return false
	}

	for _, v := range t.items {
		item := t.GetOwner().CheckItemManipulation(v.GetObjectId(), v.GetCount())
		if item == nil || item.GetCount() < 1 {
			log.Println(t.GetOwner().GetName() + ": Invalid Item in TradeList")
			return false
		}
	}
	return true
}

func (t *TradeList) CalcItemsWeight() int32 {
	weight := 0.0

	for _, v := range t.items {
		if v == nil {
			continue
		}
		weight += float64(v.GetCount()) * float64(v.GetBaseItem().GetWeight())
	}

	return int32(math.Min(weight, math.MaxInt32))
}
