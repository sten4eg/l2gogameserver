package attribute

import (
	"errors"
	"strings"
)

type Attribute int32

const (
	Fire Attribute = iota
	Water
	Wind
	Earth
	Holy
	Unholy
	None Attribute = -2
)

func (t *Attribute) UnmarshalJSON(data []byte) error {

	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "fire":
		*t = Fire
	case "water":
		*t = Water
	case "wind":
		*t = Wind
	case "earth":
		*t = Earth
	case "holy":
		*t = Holy
	case "unholy":
		*t = Unholy
	case "none":
		*t = None
	default:
		return errors.New("Неправильный Attribute: " + sData)
	}
	return nil
}
