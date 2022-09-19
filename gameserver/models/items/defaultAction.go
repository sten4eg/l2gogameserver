package items

import (
	"errors"
	"strings"
)

type DefaultAction int32

const (
	Calc DefaultAction = iota
	CallSkill
	Capsule
	CreateMpcc
	Dice
	Equip
	Fishingshot
	Harvest
	HideName
	KeepExp
	NickColor
	None
	Peel
	Recipe
	Seed
	ShowAdventurerGuideBook
	ShowHtml
	ShowSsqStatus
	SkillMaintain
	SkillReduce
	Soulshot
	Spiritshot
	StartQuest
	SummonSoulshot
	SummonSpiritshot
	XmasOpen
)

func (a *DefaultAction) UnmarshalJSON(data []byte) error {
	sData := strings.ReplaceAll(string(data), "\"", "")
	switch sData {
	case "action_equip":
		*a = Equip
	case "action_peel":
		*a = Peel
	case "action_none":
		*a = None
	case "action_skill_reduce":
		*a = SkillReduce
	case "action_soulshot":
		*a = Soulshot
	case "action_recipe":
		*a = Recipe
	case "action_skill_maintain":
		*a = SkillMaintain
	case "action_spiritshot":
		*a = Spiritshot
	case "action_dice":
		*a = Dice
	case "action_calc":
		*a = Calc
	case "action_seed":
		*a = Seed
	case "action_harvest":
		*a = Harvest
	case "action_capsule":
		*a = Capsule
	case "action_xmas_open":
		*a = XmasOpen
	case "action_show_html":
		*a = ShowHtml
	case "action_show_ssq_status":
		*a = ShowSsqStatus
	case "action_fishingshot":
		*a = Fishingshot
	case "action_summon_soulshot":
		*a = SummonSoulshot
	case "action_summon_spiritshot":
		*a = SummonSpiritshot
	case "action_call_skill":
		*a = CallSkill
	case "action_show_adventurer_guide_book":
		*a = ShowAdventurerGuideBook
	case "action_keep_exp":
		*a = KeepExp
	case "action_create_mpcc":
		*a = CreateMpcc
	case "action_nick_color":
		*a = NickColor
	case "action_hide_name":
		*a = HideName
	case "action_start_quest":
		*a = StartQuest
	default:
		return errors.New("неправильный DefaultAction: " + sData)
	}
	return nil
}
