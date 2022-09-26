package weaponType

import (
	"errors"
	"strings"
)

type WeaponType int16

const (
	SWORD WeaponType = iota
	BLUNT
	DAGGER
	BOW
	POLE
	NONE
	DUAL
	ETC
	FIST
	DUALFIST
	FISHINGROD
	RAPIER
	ANCIENTSWORD
	CROSSBOW
	FLAG
	OWNTHING
	DUALDAGGER
)

func (t *WeaponType) UnmarshalJSON(data []byte) error {

	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "none":
		*t = NONE
	case "ancientsword":
		*t = ANCIENTSWORD
	case "pole":
		*t = POLE
	case "sword":
		*t = SWORD
	case "blunt":
		*t = BLUNT
	case "dagger":
		*t = DAGGER
	case "bow":
		*t = BOW
	case "dual":
		*t = DUAL
	case "etc":
		*t = ETC
	case "fist":
		*t = FIST
	case "dualfist":
		*t = DUALFIST
	case "fishingrod":
		*t = FISHINGROD
	case "rapier":
		*t = RAPIER
	case "crossbow":
		*t = CROSSBOW
	case "flag":
		*t = FLAG
	case "ownthing":
		*t = OWNTHING
	case "dualdagger":
		*t = DUALDAGGER
	default:
		return errors.New("Неправильный WeaponType: " + sData)
	}
	return nil
}
