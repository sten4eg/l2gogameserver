package interfaces

import "database/sql"

type TradeListInterface interface {
	SetPartner(CharacterI)
	GetPartner() CharacterI
	Lock()
	AddItem(int32, int64, CharacterI, int64) TradableItemInterface
	IsLocked() bool
	IsConfirmed() bool
	GetOwner() CharacterI
	InvalidateConfirmation()
	Confirmed(*sql.DB) (bool, bool, bool, bool)
	MuLock()
	MuUnlock()
	Validate() bool
	CalcItemsWeight() int
	CountItemSlots(CharacterI) int
	TransferItems(*sql.DB) bool
	AdjustAvailableItem(item MyItemInterface) TradableItemInterface
	GetItems() []TradableItemInterface
	SetTitle(string)
	GetTitle() string
	Clear()
	SetPackaged(bool)
	IsPackaged() bool
	PrivateStoreBuy(CharacterI, []ItemRequestInterface, *sql.DB) byte
	AddItemByItemId(int32, int64, int64) TradableItemInterface
	GetAvailableItems(inventory InventoryInterface) []TradableItemInterface
	UpdateItems()
	PrivateStoreSell(character CharacterI, items []ItemRequestInterface, db *sql.DB) bool
}
