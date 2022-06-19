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
//
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
func (sys *SysMsg) AddZone(x, y, z int32) {
	sys.Params = append(sys.Params, Params{tType: TypeZoneName, value: [3]int32{x, y, z}})
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

var CantMoveSitting = SysMsg{Params: []Params{}, Id: 31}
var CannotEquipItemDueToBadCondition = SysMsg{Params: []Params{}, Id: 1518}
var TargetIsIncorrect = SysMsg{Params: []Params{}, Id: 144}
var C1IsBusyTryLater = SysMsg{Params: []Params{}, Id: 153}
var AlreadyTrading = SysMsg{Params: []Params{}, Id: 142}
var TargetTooFar = SysMsg{Params: []Params{}, Id: 22}
var S1 = SysMsg{Params: []Params{}, Id: 1987}
var RequestC1ForTrade = SysMsg{Params: []Params{}, Id: 118}
var TargetIsNotFoundInTheGame = SysMsg{Params: []Params{}, Id: 145}
var C1DeniedTradeRequest = SysMsg{Params: []Params{}, Id: 119}
var BeginTradeWithC1 = SysMsg{Params: []Params{}, Id: 120}
var C1CanceledTrade = SysMsg{Params: []Params{}, Id: 124}
var CannotAdjustItemsAfterTradeConfirmed = SysMsg{Params: []Params{}, Id: 124}
var NothingHappened = SysMsg{Params: []Params{}, Id: 61}
var C1ConfirmedTrade = SysMsg{Params: []Params{}, Id: 121}

func FindById() {

}
