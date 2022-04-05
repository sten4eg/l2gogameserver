package sysmsg

type SysMsg struct {
	Id  int32
	Val string
}

var YouHaveInvitedTheWrongTarget = SysMsg{Val: "", Id: 152}
var CantMoveSitting = SysMsg{Val: "", Id: 31}
var CannotEquipItemDueToBadCondition = SysMsg{Val: "", Id: 1518}
