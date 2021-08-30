package items

import (
	"errors"
	"strings"
)

type ItemType int16

const (
	Weapon = iota
	ShieldOrArmor
	Accessory
	Quest
	Money
	Other
)

func (t *ItemType) UnmarshalJSON(data []byte) error {

	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "etcitem":
		*t = Other
	case "weapon":
		*t = Weapon
	case "accessary":
		*t = Accessory
	case "armor":
		*t = ShieldOrArmor
	case "asset":
		*t = Money
	case "questitem":
		*t = Quest
	default:
		return errors.New("неправильный ItemType: " + sData)
	}
	return nil
}
