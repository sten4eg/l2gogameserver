package crystalType

import (
	"errors"
	"strings"
)

type CrystalType int32

const (
	None CrystalType = iota
	D
	C
	B
	A
	S
)

func (m *CrystalType) UnmarshalJSON(data []byte) error {
	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "s", "s84", "s80":
		*m = S
	case "b":
		*m = B
	case "none", "crystal_free", "event":
		*m = None
	case "d":
		*m = D
	case "a":
		*m = A
	case "c":
		*m = C
	default:
		return errors.New("неправильный MaterialType: " + sData)
	}
	return nil
}
