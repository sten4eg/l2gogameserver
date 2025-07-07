package interfaces

import (
	"database/sql"
	"sync"
)

type InventoryInterface interface {
	sync.Locker
	GetItemByObjectId(id int32) MyItemInterface
	GetItemByItemId(int) MyItemInterface
	CanManipulateWithItemId(id int32) bool
	GetItemsWithUpdatedType() []MyItemInterface
	SetAllItemsUpdatedTypeNone()
	ValidateWeight(int) bool
	ValidateCapacity(int, CharacterI) bool
	AddItem(item MyItemInterface, db *sql.DB) MyItemInterface
	AddItem2(itemId int32, count int, stackable bool, db *sql.DB) MyItemInterface
	RefreshWeight()
	TransferItem(int32, int, InventoryInterface, CharacterI, *sql.DB) MyItemInterface
	RemoveItem(MyItemInterface) bool
	DestroyItem(MyItemInterface, int, *sql.DB) MyItemInterface
	GetAdenaCount() int64
	GetAvailableItems(tradeList TradeListInterface, char CharacterI) []TradableItemInterface
	GetUniqueItems(character CharacterI, allowAdena, allowAncientAdena, onlyAvailable bool) []MyItemInterface
	GetItemsByItemId(int32) []MyItemInterface
	AdjustAvailableItem(TradableItemInterface)
	GetItems() []MyItemInterface
}
