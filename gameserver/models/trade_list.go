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
func (t *TradeList) Confirmed() (bool, needSendTradeConfirm bool, tradeDone bool, success bool) {
	needSendTradeConfirm = false
	tradeDone = false
	success = false
	if t.confirmed {
		return true, needSendTradeConfirm, tradeDone, success
	}

	partner := t.GetPartner()
	if partner != nil {
		partnerList := partner.GetActiveTradeList()
		if partnerList == nil {
			log.Println(partner.GetName() + ": Trading partner (" + partner.GetName() + ") is invalid in this trade!")
			return false, needSendTradeConfirm, tradeDone, success
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
				return false, needSendTradeConfirm, tradeDone, success
			}
			if !t.Validate() {
				return false, needSendTradeConfirm, tradeDone, success
			}

			success = t.doExchange(partnerList)
			tradeDone = true
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
	return t.confirmed, needSendTradeConfirm, tradeDone, success
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
	t.items = append(t.items, r)
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

func (t *TradeList) CalcItemsWeight() int {
	weight := 0.0

	for _, v := range t.items {
		if v == nil {
			continue
		}
		weight += float64(v.GetCount()) * float64(v.GetBaseItem().GetWeight())
	}

	return int(math.Min(weight, math.MaxInt32))
}

func (t *TradeList) doExchange(partnerList interfaces.TradeListInterface) bool {
	success := false
	owner := t.GetOwner()
	partner := partnerList.GetOwner()

	if !owner.GetInventory().ValidateWeight(partnerList.CalcItemsWeight()) || !partner.GetInventory().ValidateWeight(t.CalcItemsWeight()) {
		owner.EncryptAndSend(sysmsg.SystemMessage(sysmsg.WeightLimitExceeded))
		partner.EncryptAndSend(sysmsg.SystemMessage(sysmsg.WeightLimitExceeded))
	} else if !owner.GetInventory().ValidateCapacity(partnerList.CountItemSlots(owner), owner) || !partner.GetInventory().ValidateCapacity(t.CountItemSlots(partner), partner) {
		owner.EncryptAndSend(sysmsg.SystemMessage(sysmsg.SlotsFull))
		partner.EncryptAndSend(sysmsg.SystemMessage(sysmsg.SlotsFull))
	} else {
		partnerList.TransferItems()
		t.TransferItems()
		success = true
	}

	return success
}

func (t *TradeList) CountItemSlots(partner interfaces.CharacterI) int {
	var slots int

	for _, item := range t.items {
		if item == nil {
			continue
		}
		if !item.IsStackable() {
			slots += int(item.GetCount())
		} else if partner.GetInventory().GetItemByItemId(int(item.GetBaseItem().GetId())) == nil {
			slots++
		}
	}

	return slots
}

func (t *TradeList) TransferItems() bool {
	for _, tItem := range t.items {
		oldItem := t.GetOwner().GetInventory().GetItemByObjectId(tItem.GetObjectId())
		if oldItem == nil {
			return false
		}
		newItem := t.GetOwner().GetInventory().TransferItem(tItem.GetObjectId(), int(tItem.GetCount()), t.partner.GetInventory(), t.partner)
		if newItem == nil {
			return false
		}
	}
	return true
}
