package items

import (
	"encoding/json"
	"l2gogameserver/data/logger"
)

// ItemBonusStat возможные значения для Type :
// mp_bonus, magical_defense, physical_defense
// physical_damage, magical_damage, critical
// attack_speed
type ItemBonusStat struct {
	Type  string
	Val   float64
	Order int
}

type stat struct {
	Stat string
	Val  float64
}

func (b *ItemBonusStat) UnmarshalJSON(data []byte) error {
	var s stat
	err := json.Unmarshal(data, &s)
	if err != nil {
		logger.Error.Panicln("Неверный Stats, stat:" + string(data))
	}
	*b = ItemBonusStat{
		Val:   s.Val,
		Type:  s.Stat,
		Order: 0,
	}
	return nil
}
