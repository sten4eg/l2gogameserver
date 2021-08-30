package items

import (
	"errors"
	"strings"
)

type SlotBitType int32

var SlotNone SlotBitType = 0x0000
var SlotUnderwear SlotBitType = 0x0001
var SlotREar SlotBitType = 0x0002
var SlotLEar SlotBitType = 0x0004
var SlotLrEar SlotBitType = 0x00006
var SlotNeck SlotBitType = 0x0008
var SlotRFinger SlotBitType = 0x0010
var SlotLFinger SlotBitType = 0x0020
var SlotLrFinger SlotBitType = 0x0030
var SlotHead SlotBitType = 0x0040
var SlotRHand SlotBitType = 0x0080
var SlotLHand SlotBitType = 0x0100
var SlotGloves SlotBitType = 0x0200
var SlotChest SlotBitType = 0x0400
var SlotLegs SlotBitType = 0x0800
var SlotFeet SlotBitType = 0x1000
var SlotBack SlotBitType = 0x2000
var SlotLrHand SlotBitType = 0x4000
var SlotFullArmor SlotBitType = 0x8000
var SlotHair SlotBitType = 0x010000
var SlotAlldress SlotBitType = 0x020000
var SlotHair2 SlotBitType = 0x040000
var SlotHairall SlotBitType = 0x080000
var SlotRBracelet SlotBitType = 0x100000
var SlotLBracelet SlotBitType = 0x200000
var SlotDeco SlotBitType = 0x400000
var SlotBelt SlotBitType = 0x10000000
var SlotWolf SlotBitType = -100
var SlotHatchling SlotBitType = -101
var SlotStrider SlotBitType = -102
var SlotBabypet SlotBitType = -103
var SlotGreatwolf SlotBitType = -104

var SLOT_MULTI_ALLWEAPON SlotBitType = SlotLrHand | SlotRHand

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
