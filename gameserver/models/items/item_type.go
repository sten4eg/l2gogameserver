package items

import (
	"errors"
	"strings"
)

type ItemType1 int16

const (
	WeaponRingEarringNecklace = 0
	ShieldArmor               = 1
	ITEM_QUESTITEM_ADENA      = 4
)

type ItemType2 int16

const (
	Weapon = iota
	ShieldOrArmor
	Accessory
	Quest
	Money
	Other
)

func (t *ItemType2) UnmarshalJSON(data []byte) error {

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
		return errors.New("неправильный ItemType2: " + sData)
	}
	return nil
}
