package consumeType

import (
	"errors"
	"strings"
)

type ConsumeType int32

const (
	// Stackable складывается
	Stackable ConsumeType = iota
	// Normal не может складываться
	Normal
	// Asset только золотые монеты ?
	Asset
)

func (m *ConsumeType) UnmarshalJSON(data []byte) error {
	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "consume_type_stackable":
		*m = Stackable
	case "consume_type_normal":
		*m = Normal
	case "consume_type_asset":
		*m = Asset
	default:
		return errors.New("неправильный ConsumeType: " + sData)
	}
	return nil
}
