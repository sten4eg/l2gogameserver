package sysmsg

type SysMsg struct {
	Id     int32
	Params []Params
}
type Params struct {
	tType ParamType
	value interface{}
}
type ParamType = byte

const (
	TypeSystemString ParamType = 13
	TypePlayerName   ParamType = 12
	TypeDoorName     ParamType = 11
	TypeInstanceName ParamType = 10
	TypeElementName  ParamType = 9
	// id 8 - same as 3
	TypeZoneName   ParamType = 7
	TypeLongNumber ParamType = 6
	TypeCastleName ParamType = 5
	TypeSkillName  ParamType = 4
	TypeItemName   ParamType = 3
	TypeNpcName    ParamType = 2
	TypeIntNumber  ParamType = 1
	TypeText       ParamType = 0
)

func (sys *Params) GetType() ParamType {
	return sys.tType
}

// AddCastleId
// Appends a Castle name parameter type, the name will be read from CastleName-e.dat.<br>
// <ul>
// <li>1-9 Castle names</li>
// <li>21 Fortress of Resistance</li>
// <li>22-33 Clan Hall names</li>
// <li>34 Devastated Castle</li>
// <li>35 Bandit Stronghold</li>
// <li>36-61 Clan Hall names</li>
// <li>62 Rainbow Springs</li>
// <li>63 Wild Beast Reserve</li>
// <li>64 Fortress of the Dead</li>
// <li>81-89 Territory names</li>
// <li>90-100 null</li>
// <li>101-121 Fortress names</li>
func (sys *SysMsg) AddCastleId(number int32) {
	sys.Params = append(sys.Params, Params{tType: TypeCastleName, value: number})
}

func (sys *SysMsg) AddInt(number int32) {
	sys.Params = append(sys.Params, Params{tType: TypeIntNumber, value: number})
}

func (sys *SysMsg) AddLong(number int64) {
	sys.Params = append(sys.Params, Params{tType: TypeLongNumber, value: number})
}

func (sys *SysMsg) AddString(str string) {
	sys.Params = append(sys.Params, Params{tType: TypeText, value: str})
}
func (sys *SysMsg) AddItemName(id int32) {
	sys.Params = append(sys.Params, Params{tType: TypeItemName, value: id})
}

func (sys *SysMsg) AddZone(x, y, z int32) {
	sys.Params = append(sys.Params, Params{tType: TypeZoneName, value: [3]int32{x, y, z}})
}
func (sys *SysMsg) AddCharacterName(name string) {
	sys.Params = append(sys.Params, Params{tType: TypePlayerName, value: name})
}
func (sys *Params) GetValueString() string {
	return sys.value.(string)
}
func (sys *Params) GetValueInt64() int64 {
	return sys.value.(int64)
}
func (sys *Params) GetValueInt32() int32 {
	return sys.value.(int32)
}
func (sys *Params) GetTwoElementSlice() [2]int32 {
	return sys.value.([2]int32)
}
func (sys *Params) GetThreeElementSlice() [3]int32 {
	return sys.value.([3]int32)
}

var TargetTooFar = SysMsg{Params: []Params{}, Id: 22}
var CantMoveSitting = SysMsg{Params: []Params{}, Id: 31}
var NothingHappened = SysMsg{Params: []Params{}, Id: 61}
var C1InvitedToParty = SysMsg{Params: []Params{}, Id: 105}
var YouJoinedS1Party = SysMsg{Params: []Params{}, Id: 106}
var C1JoinedParty = SysMsg{Params: []Params{}, Id: 107}
var C1LeftParty = SysMsg{Params: []Params{}, Id: 108}
var IncorrectTarget = SysMsg{Params: []Params{}, Id: 109}
var RequestC1ForTrade = SysMsg{Params: []Params{}, Id: 118}
var C1DeniedTradeRequest = SysMsg{Params: []Params{}, Id: 119}
var BeginTradeWithC1 = SysMsg{Params: []Params{}, Id: 120}
var C1ConfirmedTrade = SysMsg{Params: []Params{}, Id: 121}
var CannotAdjustItemsAfterTradeConfirmed = SysMsg{Params: []Params{}, Id: 122}
var TradeSuccessful = SysMsg{Params: []Params{}, Id: 123}
var C1CanceledTrade = SysMsg{Params: []Params{}, Id: 124}
var SlotsFull = SysMsg{Params: []Params{}, Id: 129}
var AlreadyTrading = SysMsg{Params: []Params{}, Id: 142}
var TargetIsIncorrect = SysMsg{Params: []Params{}, Id: 144}
var TargetIsNotFoundInTheGame = SysMsg{Params: []Params{}, Id: 145}
var YouHaveInvitedTheWrongTarget = SysMsg{Params: []Params{}, Id: 152}
var C1IsBusyTryLater = SysMsg{Params: []Params{}, Id: 153}
var OnlyLeaderCanInvite = SysMsg{Params: []Params{}, Id: 154}
var PartyFull = SysMsg{Params: []Params{}, Id: 155}
var C1IsAlreadyInParty = SysMsg{Params: []Params{}, Id: 160}
var FirstSelectUserToInviteToParty = SysMsg{Params: []Params{}, Id: 185}
var YouLeftParty = SysMsg{Params: []Params{}, Id: 200}
var C1WasExpelledFromParty = SysMsg{Params: []Params{}, Id: 201}
var HaveBeenExpelledFromParty = SysMsg{Params: []Params{}, Id: 202}
var PartyDispersed = SysMsg{Params: []Params{}, Id: 203}
var YouNotEnoughAdena = SysMsg{Params: []Params{}, Id: 279}
var IncorrectItemCount = SysMsg{Params: []Params{}, Id: 347}
var C1purchasedS2 = SysMsg{Params: []Params{}, Id: 378}
var C1PurchasedS3S2S = SysMsg{Params: []Params{}, Id: 380}
var AnotherLoginWithAccount = SysMsg{Params: []Params{}, Id: 421}
var WeightLimitExceeded = SysMsg{Params: []Params{}, Id: 422}
var PurchasedS2FromC1 = SysMsg{Params: []Params{}, Id: 559}
var PurchasedS3S2SFromC1 = SysMsg{Params: []Params{}, Id: 561}
var ThePurchasePriceIsHigherThanMoney = SysMsg{Params: []Params{}, Id: 720}
var YouMayCreateUpTo48Macros = SysMsg{Params: []Params{}, Id: 797}
var InvalidMacro = SysMsg{Params: []Params{}, Id: 810}
var MacroDescriptionMax32Chars = SysMsg{Params: []Params{}, Id: 837}
var EnterTheMacroName = SysMsg{Params: []Params{}, Id: 838}
var YouHaveExceededQuantityThatCanBeInputted = SysMsg{Params: []Params{}, Id: 1036}
var NoPrivateStoreHere = SysMsg{Params: []Params{}, Id: 1296}
var C1HasBecomeAPartyLeader = SysMsg{Params: []Params{}, Id: 1384}
var CannotEquipItemDueToBadCondition = SysMsg{Params: []Params{}, Id: 1518}
var S1 = SysMsg{Params: []Params{}, Id: 1987}
var CoupleActionDenied = SysMsg{Params: []Params{}, Id: 3119}
var TargetDoNotMeetLocRequirements = SysMsg{Params: []Params{}, Id: 3120}
var C1IsInPrivateShopModeOrInABattleAndCannotBeRequestedForACoupleAction = SysMsg{Params: []Params{}, Id: 3123}
var C1IsAlreadyParticipatingInACoupleActionAndCannotBeRequestedForAnotherCoupleAction = SysMsg{Params: []Params{}, Id: 3126}
var C1IsInAChaoticStateAndCannotBeRequestedForACoupleAction = SysMsg{Params: []Params{}, Id: 3127}
var YouHaveRequestedCoupleActionC1 = SysMsg{Params: []Params{}, Id: 3150}
var C1IsSetToRefuseCoupleActions = SysMsg{Params: []Params{}, Id: 3164}
var C1IsSetToRefusePartyRequest = SysMsg{Params: []Params{}, Id: 3168}

func FindById() {

}
