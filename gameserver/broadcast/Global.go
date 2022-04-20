package broadcast

import (
	"l2gogameserver/gameserver/interfaces"
	"sync"
)

var OnlineCharacters OnlineCharactersStruct

type OnlineCharactersStruct struct {
	Char map[int32]interfaces.ReciverAndSender
	Mu   sync.Mutex
}
