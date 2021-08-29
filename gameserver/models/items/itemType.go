package items

import (
	"errors"
	"strings"
)

type ItemType int

const (
	Weapon = iota
	ShieldOrArmor
	Accessory
	Quest
	Money
	Other
)

func (t *ItemType) UnmarshalJSON(data []byte) error {

	sData := strings.ReplaceAll(string(data), "[", "")
	sData = strings.ReplaceAll(sData, "", "")
	sData = strings.ReplaceAll(sData, "\n", "")
	sData = strings.ReplaceAll(sData, " ", "")
	sData = strings.ReplaceAll(sData, "]", "")
	sData = strings.ReplaceAll(sData, "\"", "")

	switch sData {
	case "lhand":
		*t = ShieldOrArmor
	case "none":
		*t = Other
	case "lrhand":
		*t = ShieldOrArmor
	case "rfinger,lfinger":
		*t = ShieldOrArmor
	default:
		return errors.New("неправильный ItemType: " + sData)
	}
	return nil
}
