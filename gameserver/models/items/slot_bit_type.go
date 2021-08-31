package items

import (
	"errors"
	"strings"
)

type SlotBitType int32

const (
	SlotNone             SlotBitType = 0x0000
	SlotUnderwear        SlotBitType = 0x0001
	SlotREar             SlotBitType = 0x0002
	SlotLEar             SlotBitType = 0x0004
	SlotLrEar            SlotBitType = 0x00006
	SlotNeck             SlotBitType = 0x0008
	SlotRFinger          SlotBitType = 0x0010
	SlotLFinger          SlotBitType = 0x0020
	SlotLrFinger         SlotBitType = 0x0030
	SlotHead             SlotBitType = 0x0040
	SlotRHand            SlotBitType = 0x0080
	SlotLHand            SlotBitType = 0x0100
	SlotGloves           SlotBitType = 0x0200
	SlotChest            SlotBitType = 0x0400
	SlotLegs             SlotBitType = 0x0800
	SlotFeet             SlotBitType = 0x1000
	SlotBack             SlotBitType = 0x2000
	SlotLrHand           SlotBitType = 0x4000
	SlotFullArmor        SlotBitType = 0x8000
	SlotHair             SlotBitType = 0x010000
	SlotAlldress         SlotBitType = 0x020000
	SlotHair2            SlotBitType = 0x040000
	SlotHairall          SlotBitType = 0x080000
	SlotRBracelet        SlotBitType = 0x100000
	SlotLBracelet        SlotBitType = 0x200000
	SlotDeco             SlotBitType = 0x400000
	SlotBelt             SlotBitType = 0x10000000
	SlotWolf             SlotBitType = -100
	SlotHatchling        SlotBitType = -101
	SlotStrider          SlotBitType = -102
	SlotBabypet          SlotBitType = -103
	SlotGreatwolf        SlotBitType = -104
	SLOT_MULTI_ALLWEAPON SlotBitType = SlotLrHand | SlotRHand
)

func (t *SlotBitType) UnmarshalJSON(data []byte) error {

	sData := strings.ReplaceAll(string(data), "[", "")
	sData = strings.ReplaceAll(sData, "", "")
	sData = strings.ReplaceAll(sData, "\n", "")
	sData = strings.ReplaceAll(sData, " ", "")
	sData = strings.ReplaceAll(sData, "]", "")
	sData = strings.ReplaceAll(sData, "\"", "")
	sData = strings.ReplaceAll(sData, "\r", "")

	switch sData {
	case "lhand":
		*t = SlotLHand
	case "none":
		*t = SlotNone
	case "lrhand":
		*t = SlotLrHand
	case "rfinger,lfinger":
		*t = SlotLrFinger
	case "rhand":
		*t = SlotRHand
	case "chest":
		*t = SlotChest
	case "legs":
		*t = SlotLegs
	case "feet":
		*t = SlotFeet
	case "head":
		*t = SlotHead
	case "gloves":
		*t = SlotGloves
	case "onepiece":
		*t = SlotFullArmor
	case "rear,lear":
		*t = SlotLrEar
	case "neck":
		*t = SlotNeck
	case "back":
		*t = SlotBack
	case "underwear":
		*t = SlotUnderwear
	case "hair":
		*t = SlotHair
	case "alldress":
		*t = SlotAlldress
	case "hair2":
		*t = SlotHair2
	case "hairall":
		*t = SlotHairall
	case "rbracelet":
		*t = SlotRBracelet
	case "lbracelet":
		*t = SlotLBracelet
	case "deco1":
		*t = SlotDeco
	case "waist":
		*t = SlotBelt

	default:
		return errors.New("неправильный ItemType: " + sData)
	}
	return nil
}
