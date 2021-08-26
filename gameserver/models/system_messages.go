package models

type SysMsg struct {
	Id  int32
	Val string
}

var YouHaveInvitedTheWrongTarget = SysMsg{Val: "", Id: 152}
var CantMoveSitting = SysMsg{Val: "", Id: 31}

func q() {
	w := CantMoveSitting
	_ = w
}
