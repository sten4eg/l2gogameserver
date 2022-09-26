package skills

import (
	"errors"
	"strings"
)

type OperateType int

const (
	// A1  Active Skill with "Instant Effect" (for example damage skills heal/pdam/mdam/cpdam skills).
	A1 OperateType = iota
	// A2  Active Skill with "Continuous effect + Instant effect" (for example buff/debuff or damage/heal over time skills).
	A2
	// A3  Active Skill with "Instant effect + Continuous effect"
	A3
	// A4  Active Skill with "Instant effect + ?" used for special event herb.
	A4
	// CA1  Continuous Active Skill with "instant effect" (instant effect casted by ticks).
	CA1
	// CA5  Continuous Active Skill with "continuous effect" (continuous effect casted by ticks).
	CA5
	// DA1  Directional Active Skill with "Charge/Rush instant effect".
	DA1
	// DA2  Directional Active Skill with "Charge/Rush Continuous effect".
	DA2
	// P  Passive Skill.
	P
	// T  Toggle Skill.
	T
)

func (o *OperateType) IsActive() bool {
	switch *o {
	case A1, A2, A3, A4, CA1, CA5, DA1, DA2:
		return true
	default:
		return false
	}
}

func (o *OperateType) IsPassive() bool {
	return *o == P
}

func (o *OperateType) IsContinuous() bool {
	switch *o {
	case A2, A4, DA2:
		return true
	default:
		return false
	}
}

func (o *OperateType) IsSelfContinuous() bool {
	return *o == A3
}

func (o *OperateType) IsToggle() bool {
	return *o == T
}

func (o *OperateType) IsChanneling() bool {
	switch *o {
	case CA1, CA5:
		return true
	default:
		return false
	}
}
func (o *OperateType) IsFlyType() bool {
	switch *o {
	case DA1, DA2:
		return true
	default:
		return false
	}
}
func (o *OperateType) UnmarshalJSON(data []byte) error {
	sData := strings.Trim(string(data), "\"")
	switch sData {
	case "A1":
		*o = A1
	case "A2":
		*o = A2
	case "A3":
		*o = A3
	case "A4":
		*o = A4
	case "CA1":
		*o = CA1
	case "CA5":
		*o = CA5
	case "DA1":
		*o = DA1
	case "DA2":
		*o = DA2
	case "P":
		*o = P
	case "T":
		*o = T
	default:
		return errors.New("неправильный OperateType скила, OpType = " + string(data))
	}
	return nil
}
