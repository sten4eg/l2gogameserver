package models

import (
	"l2gogameserver/config"
	"l2gogameserver/gameserver/interfaces"
	items2 "l2gogameserver/gameserver/models/items"
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

func (t *TradeList) SetTitle(title string) {
	t.title = title
}

func (t *TradeList) GetTitle() string {
	return t.title
}

func (t *TradeList) SetPackaged(packaged bool) {
	t.packaged = packaged
}

func (t *TradeList) IsPackaged() bool {
	return t.packaged
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
			//partner.SendSysMsg(sysmsg.AlreadyTrading)
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

// AdjustAvailableItem если предмет не стакается, добовляет его в трейд лист, если предмет стакается, проверяет что в инвентаре достаточно данного предмета(количество).
func (t *TradeList) AdjustAvailableItem(item interfaces.MyItemInterface) interfaces.TradableItemInterface {
	if item.IsStackable() {
		for _, exclItem := range t.items {
			if exclItem.GetId() == item.GetId() {
				if item.GetCount() <= exclItem.GetCount() {
					return nil
				}
				return NewTradeItem(item, item.GetCount()-exclItem.GetCount(), int64(item.GetDefaultPrice()))
			}
		}
	}
	return NewTradeItem(item, item.GetCount(), int64(item.GetDefaultPrice()))
}

func (t *TradeList) GetItems() []interfaces.TradableItemInterface {
	return t.items
}

func (t *TradeList) Clear() {
	t.MuLock()
	t.items = nil
	t.locked = false
	t.MuUnlock()
}

// PrivateStoreBuy @return byte: результат трейда. 0 - ок, 1 - отменен (недостаточно адены), 2 - неудача (ошибка предмета)
func (t *TradeList) PrivateStoreBuy(character interfaces.CharacterI, items []interfaces.ItemRequestInterface) byte {
	t.MuLock()
	defer t.MuUnlock()

	if t.locked {
		return 1
	}

	if !t.Validate() {
		t.locked = true
		return 1
	}

	//TODO проверка что оба игрока онлайн

	var slots int32
	var weight int32
	var totalPrice int64

	ownerInventory := t.owner.GetInventory()
	playerInventory := character.GetInventory()

	for _, item := range items {
		found := false

		for _, tradeItem := range t.items {
			if tradeItem.GetObjectId() == item.GetObjectId() {
				if tradeItem.GetPrice() == item.GetPrice() {
					if tradeItem.GetCount() < item.GetCount() {
						item.SetCount(tradeItem.GetCount())
					}
					found = true
				}
				break
			}
		}
		if !found {
			if t.IsPackaged() {
				//TODO читер бан
				return 2
			}

			item.SetCount(0)
			continue
		}

		if config.MaxAdena/item.GetCount() < item.GetPrice() {
			t.locked = true
			return 1
		}

		totalPrice += item.GetCount() * item.GetPrice()

		if config.MaxAdena < totalPrice || totalPrice < 0 {
			t.locked = true
			return 1
		}

		oldItem := t.owner.CheckItemManipulation(item.GetObjectId(), item.GetCount())
		if oldItem == nil { //TODO проврека на трейдбл
			t.locked = true
			return 2
		}

		template, _ := items2.GetItemInfo(int(item.GetId()))
		if template == nil {
			continue
		}
		weight += int32(int(item.GetCount()) * template.GetWeight())
		if !template.IsStackable() {
			slots += int32(item.GetCount())
		} else if playerInventory.GetItemByItemId(int(item.GetId())) == nil {
			slots++
		}

	}

	if totalPrice > playerInventory.GetAdenaCount() {
		character.EncryptAndSend(sysmsg.SystemMessage(sysmsg.YouNotEnoughAdena))
		return 1
	}

	if !playerInventory.ValidateWeight(int(weight)) {
		character.EncryptAndSend(sysmsg.SystemMessage(sysmsg.WeightLimitExceeded))
		return 1
	}

	if !playerInventory.ValidateCapacity(int(slots), character) {
		character.EncryptAndSend(sysmsg.SystemMessage(sysmsg.SlotsFull))
		return 1
	}

	adenaItem := playerInventory.GetItemByItemId(config.AdenaId)
	if totalPrice > adenaItem.GetCount() {
		character.EncryptAndSend(sysmsg.SystemMessage(sysmsg.YouNotEnoughAdena))
		return 1
	}
	adenaItem.ChangeCount(int(-totalPrice))
	adenaItem.UpdateDB()
	if adenaItem.GetCount() <= 0 {
		playerInventory.RemoveItem(adenaItem)
	}

	ownerInventory.AddItem2(config.AdenaId, int(totalPrice), true)

	ok := true

	for _, item := range items {
		if item.GetCount() == 0 {
			continue
		}

		oldItem := t.owner.CheckItemManipulation(item.GetObjectId(), item.GetCount())
		if oldItem == nil {
			t.locked = true
			ok = false
			break
		}

		newItem := ownerInventory.TransferItem(item.GetObjectId(), int(item.GetCount()), playerInventory, nil)
		if newItem == nil {
			ok = false
			break
		}
		t.removeItem(item.GetObjectId(), -1, item.GetCount())

		//TODO updateInventory

		if newItem.IsStackable() {
			msg1 := sysmsg.C1PurchasedS3S2S
			msg1.AddString(character.GetName())
			msg1.AddItemName(newItem.GetId())
			msg1.AddLong(item.GetCount())
			t.owner.SendSysMsg(msg1)

			msg2 := sysmsg.PurchasedS3S2SFromC1
			msg2.AddString(t.owner.GetName())
			msg2.AddItemName(newItem.GetId())
			msg2.AddLong(item.GetCount())
			character.SendSysMsg(msg2)
		} else {
			msg1 := sysmsg.C1purchasedS2
			msg1.AddString(character.GetName())
			msg1.AddItemName(newItem.GetId())
			t.owner.SendSysMsg(msg1)

			msg2 := sysmsg.PurchasedS2FromC1
			msg2.AddString(t.owner.GetName())
			msg2.AddItemName(newItem.GetId())
			character.SendSysMsg(msg2)
		}

	}

	if ok {
		return 0
	}
	return 2
}

func (t *TradeList) removeItem(objectId, itemId int32, count int64) interfaces.TradableItemInterface {
	if t.IsLocked() {
		return nil
	}

	for i, _ := range t.items {
		if t.items[i].GetObjectId() == objectId || t.items[i].GetId() == itemId {
			if t.partner != nil {
				partnerList := t.partner.GetActiveTradeList()
				if partnerList == nil {
					return nil
				}
				partnerList.InvalidateConfirmation()
			}

			var item interfaces.TradableItemInterface
			if count != -1 && t.items[i].GetCount() > count {
				item = t.items[i]
				t.items[i].SetCount(t.items[i].GetCount() - count)
			} else {
				item = t.items[i]
				t.items = append(t.items[:i], t.items[i+1:]...)
			}

			return item
		}
	}

	return nil
}
