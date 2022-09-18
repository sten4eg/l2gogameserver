package interfaces

import (
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/packets"
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
	CalculateDistanceTo(int32, int32, int32, bool, bool) float64
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
	Confirmed() (bool, bool, bool, bool)
	MuLock()
	MuUnlock()
	Validate() bool
	CalcItemsWeight() int
	CountItemSlots(CharacterI) int
	TransferItems() bool
}
type InventoryInterface interface {
	sync.Locker
	GetItemByObjectId(id int32) MyItemInterface
	GetItemByItemId(int) MyItemInterface
	CanManipulateWithItemId(id int32) bool
	GetItemsWithUpdatedType() []MyItemInterface
	SetAllItemsUpdatedTypeNone()
	ValidateWeight(int) bool
	ValidateCapacity(int, CharacterI) bool
	AddItem(item MyItemInterface, actor CharacterI) MyItemInterface
	RefreshWeight()
	TransferItem(int32, int, InventoryInterface, CharacterI) MyItemInterface
}

type MyItemInterface interface {
	sync.Locker
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
	ChangeCount(int)
	SetUpdateType(int16)
	SetCount(int64)
	UpdateDB(int32)
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
	SendSysMsg(q interface{}, options ...string)

	ClientInterface

	GetInventoryLimit() int16
	OnTradeFinish()

}
type ClientInterface interface {
	ReciverAndSender
	SetLogin(login string)
	RemoveCurrentChar()
	SetState(state clientStates.State)
	GetState() clientStates.State
	SetSessionKey(playOk1, playOk2, loginOk1, loginOk2 uint32)
	GetSessionKey() (playOk1, playOk2, loginOk1, loginOk2 uint32)
}
type ReciverAndSender interface {
	Receive() (opcode byte, data []byte, err error)
	AddLengthAndSand(data []byte)
	Send(data []byte)
	SendBuf(buffer *packets.Buffer) error
	EncryptAndSend(data []byte)
	CryptAndReturnPackageReadyToShip(data []byte) []byte
	GetCurrentChar() CharacterI
}
