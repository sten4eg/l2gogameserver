package etcItemType

import (
	"errors"
	"strings"
)

type EtcItemType int16

const (
	NONE EtcItemType = iota
	ARROW
	POTION
	SCRL_ENCHANT_WP
	SCRL_ENCHANT_AM
	SCROLL
	RECIPE
	MATERIAL
	PET_COLLAR
	CASTLE_GUARD
	LOTTO
	RACE_TICKET
	DYE
	SEED
	CROP
	MATURECROP
	HARVEST
	SEED2
	TICKET_OF_LORD
	LURE
	BLESS_SCRL_ENCHANT_WP
	BLESS_SCRL_ENCHANT_AM
	COUPON
	ELIXIR
	SCRL_ENCHANT_ATTR
	BOLT
	SCRL_INC_ENCHANT_PROP_WP
	SCRL_INC_ENCHANT_PROP_AM
	ANCIENT_CRYSTAL_ENCHANT_WP
	ANCIENT_CRYSTAL_ENCHANT_AM
	RUNE_SELECT
	RUNE
)

func (t *EtcItemType) UnmarshalJSON(data []byte) error {

	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "arrow":
		*t = ARROW
	case "none":
		*t = NONE
	case "potion":
		*t = POTION
	case "scrl_enchant_wp":
		*t = SCRL_ENCHANT_WP
	case "scrl_enchant_am":
		*t = SCRL_ENCHANT_AM
	case "scroll":
		*t = SCROLL
	case "recipe":
		*t = RECIPE
	case "material":
		*t = MATERIAL
	case "pet_collar":
		*t = PET_COLLAR
	case "castle_guard":
		*t = CASTLE_GUARD
	case "lotto":
		*t = LOTTO
	case "race_ticket":
		*t = RACE_TICKET
	case "dye":
		*t = DYE
	case "seed":
		*t = SEED
	case "crop":
		*t = CROP
	case "maturecrop":
		*t = MATURECROP
	case "harvest":
		*t = HARVEST
	case "seed2":
		*t = SEED2
	case "ticket_of_lord":
		*t = TICKET_OF_LORD
	case "lure":
		*t = LURE
	case "bless_scrl_enchant_wp":
		*t = BLESS_SCRL_ENCHANT_WP
	case "bless_scrl_enchant_am":
		*t = BLESS_SCRL_ENCHANT_AM
	case "coupon":
		*t = COUPON
	case "elixir":
		*t = ELIXIR
	case "scrl_enchant_attr":
		*t = SCRL_ENCHANT_ATTR
	case "bolt":
		*t = BOLT
	case "scrl_inc_enchant_prop_wp":
		*t = SCRL_INC_ENCHANT_PROP_WP
	case "scrl_inc_enchant_prop_am":
		*t = SCRL_INC_ENCHANT_PROP_AM
	case "teleportbookmark":
		*t = NONE //todo в l2j такого нету! проверить
	case "ancient_crystal_enchant_wp":
		*t = ANCIENT_CRYSTAL_ENCHANT_WP
	case "ancient_crystal_enchant_am":
		*t = ANCIENT_CRYSTAL_ENCHANT_AM
	case "rune_select":
		*t = RUNE_SELECT
	case "rune":
		*t = RUNE
	default:
		return errors.New("Неправильный EtcItemType: " + sData)
	}
	return nil
}
