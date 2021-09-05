package stats

import (
	"l2gogameserver/gameserver/models/items"
	"sync"
)

var AllBStats sync.Map

func AddStatTo(charId int, stat items.ItemBonusStat) {
	AllBStats.Store(charId, stat)
}
