package gs2ls

import (
	"l2gogameserver/config"
	"l2gogameserver/packets"
)

type serverStatusCodes = int32
type serverStatusValues = int32

const (
	ServerListStatus        serverStatusCodes = 1
	ServerType              serverStatusCodes = 2
	ServerListSquareBracket serverStatusCodes = 3
	MAX_PLAYERS             serverStatusCodes = 4
	ServerAge               serverStatusCodes = 6

	// Server Statuses
	StatusAuto    serverStatusValues = 0x00
	STATUS_GOOD   serverStatusValues = 0x01
	STATUS_NORMAL serverStatusValues = 0x02
	STATUS_FULL   serverStatusValues = 0x03
	STATUS_DOWN   serverStatusValues = 0x04
	StatusGmOnly  serverStatusValues = 0x05

	// Server Types
	ServerNormal             serverStatusValues = 0x01
	ServerRelax              serverStatusValues = 0x02
	ServerTest               serverStatusValues = 0x04
	ServerNoLabel            serverStatusValues = 0x08
	ServerCreationRestricted serverStatusValues = 0x10
	ServerEvent              serverStatusValues = 0x20
	ServerFree               serverStatusValues = 0x40

	// Server Ages
	ServerAgeAll serverStatusValues = 0x00
	ServerAge15  serverStatusValues = 0x0F
	ServerAge18  serverStatusValues = 0x12

	// Server BRACKET
	BracketOn  serverStatusValues = 0x01
	BracketOff serverStatusValues = 0x00
)

func ServerStatus() *packets.Buffer {
	type attr struct {
		code  int32
		value int32
	}

	brackets := config.GetServerListBrackets()
	listType := config.GetServerListType()
	isGmOnly := config.GetGMOnly()
	serverListAge := config.GetServerListAge()

	var res []attr
	if brackets {
		res = append(res, attr{code: ServerListSquareBracket, value: BracketOn})
	} else {
		res = append(res, attr{code: ServerListSquareBracket, value: BracketOff})
	}

	serverType := ServerNormal
	if listType == "Relax" {
		serverType = ServerRelax
	} else if listType == "Test" {
		serverType = ServerTest
	} else if listType == "NoLabel" {
		serverType = ServerNoLabel
	} else if listType == "Restricted" {
		serverType = ServerCreationRestricted
	} else if listType == "Event" {
		serverType = ServerEvent
	} else if listType == "Free" {
		serverType = ServerFree
	}

	res = append(res, attr{code: ServerType, value: serverType})

	if isGmOnly {
		res = append(res, attr{code: ServerListStatus, value: StatusGmOnly})
	} else {
		res = append(res, attr{code: ServerListStatus, value: StatusAuto})
	}

	if serverListAge == 0 {
		res = append(res, attr{code: ServerAge, value: ServerAgeAll})
	} else if serverType == 15 {
		res = append(res, attr{code: ServerAge, value: ServerAge15})
	} else {
		res = append(res, attr{code: ServerAge, value: ServerAge18})
	}

	buf := packets.Get()

	buf.WriteSingleByte(0x06)
	buf.WriteD(int32(len(res)))

	for i := range res {
		buf.WriteD(res[i].code)
		buf.WriteD(res[i].value)
	}
	return buf
}
