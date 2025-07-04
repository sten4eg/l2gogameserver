package interfaces

import (
	"database/sql"
	"l2gogameserver/gameserver/models/clientStates"
	"l2gogameserver/gameserver/models/items/attribute"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
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

type ItemRequestInterface interface {
	UniquerId
	Identifier
	SetCount(int64)
	GetCount() int64
	GetPrice() int64
}

type PartyInterface interface {
	GetMemberCount() int
	AddPartyMember(character CharacterI) bool
	GetLeaderObjectId() int32
	GetDistributionType() PartyDistributionTypeInterface
	SetMembers(members []CharacterI)
	GetMembers() []CharacterI
	GetLeader() CharacterI
	IsMemberInParty(character CharacterI) bool
	IsLeader(i CharacterI) bool
	IsDisbanding() bool
	SetDisbanding(bool)
	GetMemberIndex(CharacterI) int
	BroadcastParty([]byte)
	AddMember(member CharacterI)
	RemoveMember(member CharacterI)
}

type PartyDistributionTypeInterface interface {
	Identifier
	GetSysStringId() int32
}

type SkillInterface interface {
	Identifier
	IsPassive() bool
	GetLevel() int
}

type Positionable interface {
	GetObjectId() int32
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
	AddVisibleItems(MyItemInterface)
	GetItemsInRegion() []MyItemInterface
	DeleteVisibleItem(MyItemInterface)
	GetChar(int32) (CharacterI, bool)
	GetItem(int32) (MyItemInterface, bool)
	GetNpc(int32) (Npcer, bool)
	GetCharacterInRegions(int32) CharacterI
	GetX() int32
	GetY() int32
	GetZ() int32
	DropItemChecker() []int32
}
type Npcer interface {
	UniquerId
	Identifier
	IsTargetable() bool
	GetCoordinates() (x, y, z int32)
	IsAttackable() int32
	GetHeading() int32
	GetCollisionRadius() float64
	GetCollisionHeight() float64
	GetSlotRhand() int32
	GetSlotLhand() int32
	GetMaxHp() int32
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
	SetCount(count int64)
	GetLocData() int32
	IsEquipped() int16
	GetDefaultPrice() int
	GetPrice() int64
	WriteItem(buffer *packets.Buffer)
	SetStoreCount(int64)
	GetStoreCount() int64
	SetPrice(int64)
	SetObjectId(int32)
	SetEnchant(int16)
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
	UpdateDB(*sql.DB)
	GetOwnerId() int32
	SetOwnerId(ownerId int32)
	SetCoordinate(x, y, z int32)
	GetCoordinate() (x, y, z int32)
	GetDefaultPrice() int
	GetTime() int
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
	ClientInterface
	EncryptAndSend(data []byte) error
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
	SendSysMsg(q interface{}, options ...string) error
	GetActiveEnchantItemId() int32
	GetInventoryLimit() int16
	OnTradeFinish()
	GetAccountLogin() string
	DropItem(objectId int32, count int64, db *sql.DB) (MyItemInterface, MyItemInterface)
	GetSellList() TradeListInterface
	SetPrivateStoreType(value privateStoreType.PrivateStoreType)
	GetPrivateStoreType() privateStoreType.PrivateStoreType
	IsSittings() bool
	SetTarget(int32)
	GetTarget() int32
	GetBuyList() TradeListInterface
	IsinParty() bool
	SetPartyDistributionType(pdt PartyDistributionTypeInterface)
	GetParty() PartyInterface
	JoinParty(party PartyInterface) bool
	GetCurrentHp() int32
	GetMaxHp() int32
	GetCurrentMp() int32
	GetMaxMp() int32
	GetCurrentSp() int32
	GetCurrentExp() int32
	SetParty(party PartyInterface)
	GetPartyDistributionType() PartyDistributionTypeInterface
	GetSex() int32
	GetRace() race.Race
	GetBaseClass() int32
	GetLevel() int32
	GetKarma() int32
	GetPK() int32
	GetPVP() int32
	GetHairStyle() int32
	GetHairColor() int32
	GetFace() int32
	GetVitality() int32
	GetINT() int
	GetSTR() int
	GetCON() int
	GetMEN() int
	GetDEX() int
	GetWIT() int
	IsActiveWeapon() bool
	GetTitle() string
	GetClanId() int32
	GetPDef() int32
	GetPercentFromCurrentLevel(exp, level int32) float64
	GetPaperdoll() []MyItemInterface
	GetSkills() []SkillInterface
	SetSitStandPose() int32

	SetMultiSocialAction(id, targetId int32)
	GetMultiSocialAction() int32
	GetMultiSocialTarget() int32

	LoadCharactersMacros()
	GetMacrosRevision() int32
	DeleteMacros(int32) bool
	GetMacrosCount() uint8
	AddMacros(MacrosInterface)
	GetMacrosList() []MacrosInterface
}

type ClientInterface interface {
	ReciverAndSender
	SetLogin(login string)
	RemoveCurrentChar()
	SetState(state clientStates.State)
	GetState() clientStates.State
	SetSessionKey(playOk1, playOk2, loginOk1, loginOk2 uint32)
	GetSessionKey() (playOk1, playOk2, loginOk1, loginOk2 uint32)
	GetAccountLogin() string
	CloseConnection()
}

type ReciverAndSender interface {
	Receive() (opcode byte, data []byte, err error)
	AddLengthAndSand(data []byte)
	Send(data []byte)
	SendBuf(buffer *packets.Buffer) error
	EncryptAndSend(data []byte) error
	SendSysMsg(q interface{}, options ...string) error
	CryptAndReturnPackageReadyToShip(data []byte) []byte
	GetCurrentChar() CharacterI

	GetAccountLogin() string

	GetObjectIdForSlot(int32) int32
	MarkToDeleteChar(int32) int8
}

type MacrosInterface interface {
	Identifier
	GetName() string
	GetDescription() string
	GetAcronym() string
	GetIcon() byte
	GetCount() byte
	GetCommands() []MacrosCommandInterface
}

type MacrosCommandInterface interface {
	Identifier
	GetIndex() byte
	GetType() byte
	GetSkillId() int32
	GetShortcutId() byte
	GetName() string
}
