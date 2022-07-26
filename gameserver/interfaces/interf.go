package interfaces

import (
	"l2gogameserver/gameserver/models/items/attribute"
	"sync"
)

type Identifier interface {
	GetId() int32
}
type UniquerId interface {
	GetObjectId() int32
}
type Namer interface {
	GetName() string
}
type Positionable interface {
	SetX(int32)
	SetY(int32)
	SetZ(int32)
	SetXYZ(int32, int32, int32)
	SetHeading(int32)
	SetInstanceId(int32)
	GetX() int32
	GetY() int32
	GetZ() int32
	GetXYZ() (int32, int32, int32)
	GetCurrentRegion() WorldRegioner
	CalculateDistanceTo(Positionable, bool, bool) float64
	//setLocation(Location)
	//setXYZByLoc(ILocational)
}
type WorldRegioner interface {
	GetNeighbors() []WorldRegioner
	GetCharsInRegion() []CharacterI
	AddVisibleChar(CharacterI)
	GetNpcInRegion() []Npcer
	DeleteVisibleChar(CharacterI)
}
type Npcer interface {
	UniquerId
	Identifier
}
type TradableItemInterface interface {
	UniquerId
	Identifier
	BaseItemInterface
	GetBodyPart() int32
	GetEnchant() int16
	GetAttackElementType() attribute.Attribute
	GetAttackElementPower() int16
	GetElementDefAttr() [6]int16
	GetEnchantedOption() [3]int32
	GetCount() int64
}
type TradeListInterface interface {
	SetPartner(CharacterI)
	GetPartner() CharacterI
	Lock()
	AddItem(int32, int64, CharacterI, int64) TradableItemInterface
	IsLocked() bool
	IsConfirmed() bool
	GetOwner() CharacterI
	InvalidateConfirmation()
	Confirmed() (bool, bool)
	MuLock()
	MuUnlock()
	Validate() bool
	CalcItemsWeight() int32
}
type InventoryInterface interface {
	GetItemByObjectId(id int32) MyItemInterface
	CanManipulateWithItemId(id int32) bool
	GetItemsWithUpdatedType() []MyItemInterface
	SetAllItemsUpdatedTypeNone()
	sync.Locker
}
type MyItemInterface interface {
	BaseItemInterface
	UniquerId
	TradableItemInterface
	IsEquipped() int16
	GetAttackElementType() attribute.Attribute
	GetCount() int64
	GetEnchant() int16
	GetLocation() string
	GetAttackElementPower() int16
	GetElementDefAttr() [6]int16
	GetEnchantedOption() [3]int32
	GetUpdateType() int16
	GetLocData() int32
	GetMana() int32
}

type BaseItemInterface interface {
	Identifier
	IsEquipable() bool
	IsHeavyArmor() bool
	IsMagicArmor() bool
	IsArmor() bool
	IsOnlyKamaelWeapon() bool
	IsWeapon() bool
	IsWeaponTypeNone() bool
	IsStackable() bool
	GetBaseItem() BaseItemInterface
	GetItemType1() int
	GetItemType2() int16
	GetWeight() int
}
type CharacterI interface {
	Positionable
	Namer
	UniquerId
	EncryptAndSend(data []byte)
	CloseChannels()
	GetClassId() int32
	StartTransactionRequest()
	IsProcessingRequest() bool
	IsProcessingTransaction() bool
	GetTradeRefusal() bool
	OnTransactionRequest(CharacterI)
	SetActiveRequester(CharacterI)
	GetActiveRequester() CharacterI
	OnTransactionResponse()
	StartTrade(CharacterI)
	OnTradeStart(CharacterI)
	IsRequestExpired() bool
	GetActiveTradeList() TradeListInterface
	CancelActiveTrade() (bool, bool)
	OnTradeCancel() bool
	ValidateItemManipulation(int32) bool
	GetInventory() InventoryInterface
	CheckItemManipulation(int32, int64) MyItemInterface
	ValidateWeight(int32) bool
	GetMaxLoad() int32
	SendSysMsg(num int32, options ...string)
}
type ReciverAndSender interface {
	Receive() (opcode byte, data []byte, e error)
	AddLengthAndSand(d []byte)
	Send(data []byte)
	EncryptAndSend(data []byte)
	CryptAndReturnPackageReadyToShip(data []byte) []byte
	GetCurrentChar() CharacterI
}
