package armorType

import (
	"errors"
	"strings"
)

type ArmorType int16

const (
	NONE ArmorType = iota
	LIGHT
	HEAVY
	MAGIC
	SIGIL
	// L2J CUSTOM
	SHIELD
)

func (t *ArmorType) UnmarshalJSON(data []byte) error {

	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "none":
		*t = NONE
	case "light":
		*t = LIGHT
	case "heavy":
		*t = HEAVY
	case "magic":
		*t = MAGIC
	case "sigil":
		*t = SIGIL
	default:
		return errors.New("Неправильный ArmorType: " + sData)
	}
	return nil
}
