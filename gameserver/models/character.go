package models

import (
	"fmt"
	"l2gogameserver/data"
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/race"
	"sync"
	"time"
)

type Character struct {
	Login         string
	CharId        int32
	Level         int32
	MaxHp         int32
	CurHp         int32
	MaxMp         int32
	CurMp         int32
	Face          int32
	HairStyle     int32
	HairColor     int32
	Sex           int32
	Coordinates   *Coordinates
	Exp           int32
	Sp            int32
	Karma         int32
	PvpKills      int32
	PkKills       int32
	ClanId        int32
	Race          race.Race
	ClassId       int32
	BaseClass     int32
	Title         string
	OnlineTime    int32
	Nobless       int32
	Vitality      int32
	CharName      string
	CurrentRegion *WorldRegion
	Conn          *Client
	AttackEndTime int64
	// Paperdoll - массив всех слотов которые можно одеть
	Paperdoll       [26]MyItem
	Stats           StaticData
	pvpFlag         bool
	ShortCut        map[int32]dto.ShortCutDTO
	ActiveSoulShots []int32
	IsDead          bool
	IsFakeDeath     bool
	// Skills todo: проверить слайс или мапа лучше для скилов
	Skills                 map[int]Skill
	IsCastingNow           bool
	SkillQueue             chan SkillHolder
	CurrentSkill           *SkillHolder // todo А может быть без * попробовать?
	CurrentTargetId        int32
	Inventory              []MyItem
	CursedWeaponEquippedId int
	BonusStats             []items.ItemBonusStat
	F                      chan IUP
	InGame                 bool
}

func GetNewCharacterModel() *Character {
	character := new(Character)
	sk := make(map[int]Skill)
	character.Skills = sk
	character.F = make(chan IUP, 10)
	character.InGame = false
	return character
}

func (c *Character) ListenSkillQueue() {
	for {
		select {
		case res := <-c.SkillQueue:
			fmt.Println("SKILL V QUEUE")
			fmt.Println(res)
		}
	}
}

type SkillHolder struct {
	Skill        Skill
	CtrlPressed  bool
	ShiftPressed bool
}
type Coordinates struct {
	mu sync.Mutex
	X  int32
	Y  int32
	Z  int32
}

func (c *Character) SetSkillToQueue(skill Skill, ctrlPressed, shiftPressed bool) {
	s := SkillHolder{
		Skill:        skill,
		CtrlPressed:  ctrlPressed,
		ShiftPressed: shiftPressed,
	}
	c.SkillQueue <- s
}

// IsActiveWeapon есть ли у персонажа оружие в руках
func (c *Character) IsActiveWeapon() bool {
	x := c.Paperdoll[PAPERDOLL_RHAND]
	//todo Еще есть кастеты
	return x.ObjId != 0
	//todo ?
}

// GetPercentFromCurrentLevel получить % опыта на текущем уровне
func (c *Character) GetPercentFromCurrentLevel(exp, level int32) float64 {
	expPerLevel, expPerLevel2 := data.GetExpData(level)
	return float64(int64(exp)-expPerLevel) / float64(expPerLevel2-expPerLevel)
}

// SetXYZ установить координаты для персонажа
func (c *Character) SetXYZ(x, y, z int32) {
	c.Coordinates.mu.Lock()
	c.Coordinates.X = x
	c.Coordinates.Y = y
	c.Coordinates.Z = z
	c.Coordinates.mu.Unlock()
}

// GetXYZ получить координаты персонажа
func (c *Character) GetXYZ() (x, y, z int32) {
	return c.Coordinates.X, c.Coordinates.Y, c.Coordinates.Z
}

// Load загрузка персонажа
func (c *Character) Load() {
	c.InGame = true
	c.ShortCut = restoreMe(c.CharId, c.ClassId)
	c.LoadSkills()
	c.SkillQueue = make(chan SkillHolder)
	c.Inventory = GetMyItems(c.CharId)
	c.Paperdoll = RestoreVisibleInventory(c.CharId)

	for _, v := range c.Paperdoll {
		if v.ObjId != 0 {
			c.AddBonusStat(v.BonusStats)
		}
	}

	c.Stats = AllStats[int(c.ClassId)].StaticData //todo а для чего BaseClass ??

	reg := GetRegion(c.Coordinates.X, c.Coordinates.Y)
	reg.AddVisibleObject(c)
	c.CurrentRegion = reg
	go c.Shadow()
	go c.ListenSkillQueue()

}

type IUP struct {
	ObjId      int32
	UpdateType int16
}

func (c *Character) Shadow() {
	for {
		for i, v := range c.Inventory {
			if v.Item.Durability > 0 && v.Loc == Paperdoll {
				var iup IUP
				iup.ObjId = v.ObjId
				switch c.Inventory[i].Mana {

				case 0:
					iup.UpdateType = UpdateTypeRemove
					c.F <- iup
					DeleteItem(v, c)
				default:
					c.Inventory[i].Mana -= 1
					iup.UpdateType = UpdateTypeModify
					c.F <- iup
				}

			}
		}

		time.Sleep(time.Second)
	}

}

func (c *Character) checkSoulShot() {
	if len(c.ActiveSoulShots) == 0 {
		return
	}
}

func (c *Character) IsCursedWeaponEquipped() bool {
	return c.CursedWeaponEquippedId != 0
}

func (c *Character) AddBonusStat(s []items.ItemBonusStat) {
	c.BonusStats = append(c.BonusStats, s...)
}

func (c *Character) RemoveBonusStat(s []items.ItemBonusStat) {
	//for i,v := range c.BonusStats {
	//	for _,vv := range s {
	//		if v == vv {
	//			c.BonusStats[i] = c.BonusStats[len(c.BonusStats)-1] //todo переделать на безопасный вариант ) или еще что нить придумать
	//			c.BonusStats = c.BonusStats[:len(c.BonusStats)-1]
	//		}
	//	}
	//
	//}

	news := make([]items.ItemBonusStat, 0, len(c.BonusStats))
	for _, v := range c.BonusStats {
		flag := false
		for _, vv := range s {
			if v == vv {
				flag = true
				break
			}
		}
		if !flag {
			news = append(news, v)
		}
	}
	c.BonusStats = news
}

func (c *Character) GetPDef() int32 {
	var base float64
	if c.Paperdoll[PAPERDOLL_FEET].ObjId == 0 {
		base = float64(c.Stats.BasePDef.Feet)
	}
	if c.Paperdoll[PAPERDOLL_CHEST].ObjId == 0 {
		base += float64(c.Stats.BasePDef.Chest)
	}
	if c.Paperdoll[PAPERDOLL_CLOAK].ObjId == 0 {
		base += float64(c.Stats.BasePDef.Cloak)
	}
	if c.Paperdoll[PAPERDOLL_HEAD].ObjId == 0 {
		base += float64(c.Stats.BasePDef.Head)
	}
	if c.Paperdoll[PAPERDOLL_GLOVES].ObjId == 0 {
		base += float64(c.Stats.BasePDef.Gloves)
	}
	if c.Paperdoll[PAPERDOLL_LEGS].ObjId == 0 {
		base += float64(c.Stats.BasePDef.Legs)
	}
	if c.Paperdoll[PAPERDOLL_UNDER].ObjId == 0 {
		base += float64(c.Stats.BasePDef.Underwear)
	}

	for _, v := range c.BonusStats {
		if v.Type == "physical_defense" {
			base += v.Val
		}
	}
	base *= float64(c.Level+89) / 100

	return int32(base)
}

func (c *Character) GetInventoryLimit() int16 {
	if c.Race == race.DWARF {
		return 100
	}
	return 80
}
