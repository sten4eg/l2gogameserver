package interfaces

import (
	"database/sql"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/gameserver/models/trade/privateStoreType"
)

type CharacterI interface {
	Positionable
	Namer
	UniquerId
	ClientCtxInterface
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
	GetPaperdollCharInfo() []MyItemInterface
}
