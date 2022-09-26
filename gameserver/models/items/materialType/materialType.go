package materialType

import (
	"errors"
	"strings"
)

type MaterialType int32

const (
	Steel MaterialType = iota
	FineSteel
	Cotton
	BloodSteel
	Bronze
	Silver
	Gold
	Mithril
	Oriharukon
	Paper
	Wood
	Cloth
	Leather
	Bone
	Horn
	Damascus
	Adamantaite
	Chrysolite
	Crystal
	Liquid
	ScaleOfDragon
	Dyestuff
	Cobweb
	SEED
	Fish
	RuneXp
	RuneSp
	RuneRemovePenalty
)

func (m *MaterialType) UnmarshalJSON(data []byte) error {
	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "steel":
		*m = Steel
	case "oriharukon":
		*m = Oriharukon
	case "silver":
		*m = Silver
	case "fish":
		*m = Fish
	case "fine_steel":
		*m = FineSteel
	case "wood":
		*m = Wood
	case "bone":
		*m = Bone
	case "bronze":
		*m = Bronze
	case "leather":
		*m = Leather
	case "cloth":
		*m = Cloth
	case "gold":
		*m = Gold
	case "mithril":
		*m = Mithril
	case "liquid":
		*m = Liquid
	case "damascus":
		*m = Damascus
	case "adamantaite":
		*m = Adamantaite
	case "blood_steel":
		*m = BloodSteel
	case "paper":
		*m = Paper
	case "chrysolite":
		*m = Chrysolite
	case "crystal":
		*m = Crystal
	case "horn":
		*m = Horn
	case "scale_of_dragon":
		*m = ScaleOfDragon
	case "cotton":
		*m = Cotton
	case "dyestuff":
		*m = Dyestuff
	case "cobweb":
		*m = Cobweb
	case "rune_xp":
		*m = RuneXp
	case "rune_sp":
		*m = RuneSp
	case "rune_remove_penalty":
		*m = RuneRemovePenalty
	default:
		return errors.New("неправильный MaterialType: " + sData)
	}
	return nil
}
