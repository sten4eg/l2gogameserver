package models

import (
	"fmt"
	"l2gogameserver/data"
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/race"
	"l2gogameserver/utils"

	"sync"
	"time"
)

type Character struct {
	Login         string
	ObjectId      int32
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
	Inventory              Inventory
	CursedWeaponEquippedId int
	BonusStats             []items.ItemBonusStat
	F                      chan IUP
	InGame                 bool
<<<<<<< HEAD

	Macros []Macro
=======
	Target                 int32
	Macros                 []Macro
	CharInfoTo             chan []int32
	DeleteObjectTo         chan []int32
	NpcInfo                chan []Npc
	IsMoving               bool
	Sit                    bool
}

//Меняет положение персонажа от сидячего к стоячему и на оборот
//Возращает значение нового положения
func (c *Character) SetSitStandPose() int32 {
	if c.Sit == false {
		c.Sit = true
		return 0
	}
	c.Sit = false
	return 1
}

type ToSendInfo struct {
	To   []int32
	Info utils.PacketByte
>>>>>>> 67ebec2007b68bf2c47d3ecf1ae277e36cfd3071
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
	c.ShortCut = restoreMe(c.ObjectId, c.ClassId)
	c.LoadSkills()
	c.SkillQueue = make(chan SkillHolder)
	c.Inventory = GetMyItems(c.ObjectId)
	c.Paperdoll = RestoreVisibleInventory(c.ObjectId)
	c.LoadCharactersMacros()
	for _, v := range c.Paperdoll {
		if v.ObjId != 0 {
			c.AddBonusStat(v.BonusStats)
		}
	}
	c.Stats = AllStats[int(c.ClassId)].StaticData //todo а для чего BaseClass ??

	reg := GetRegion(c.Coordinates.X, c.Coordinates.Y, c.Coordinates.Z)
	c.CharInfoTo = make(chan []int32, 2)
	c.DeleteObjectTo = make(chan []int32, 2)
	c.NpcInfo = make(chan []Npc, 2)
	c.setWorldRegion(reg)

	reg.AddVisibleChar(c)

	go c.Shadow()
	go c.ListenSkillQueue()
	go c.CheckRegion()

}

type IUP struct {
	ObjId      int32
	UpdateType int16
}

func (c *Character) Shadow() {
	for {
		for i, v := range c.Inventory.Items {
			if v.Item.Durability > 0 && v.Loc == PaperdollLoc {
				var iup IUP
				iup.ObjId = v.ObjId
				switch c.Inventory.Items[i].Mana {

				case 0:
					iup.UpdateType = UpdateTypeRemove
					c.F <- iup
					DeleteItem(v, c)
				default:
					c.Inventory.Items[i].Mana -= 1
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

func (c *Character) setWorldRegion(newRegion *WorldRegion) {
	var oldAreas []*WorldRegion

	if c.CurrentRegion != nil {
		c.CurrentRegion.DeleteVisibleChar(c)
		oldAreas = c.CurrentRegion.getNeighbors()
	}

	var newAreas []*WorldRegion
	if newRegion != nil {
		newRegion.AddVisibleChar(c)
		newAreas = newRegion.getNeighbors()
	}

	// кому отправить charInfo
	deleteObjectPkgTo := make([]int32, 0, 64)
	for _, region := range oldAreas {
		if !Contains(newAreas, region) {
			region.CharsInRegion.Range(func(objId, char interface{}) bool {
				if char.(*Character).ObjectId == c.ObjectId {
					return true
				}
				deleteObjectPkgTo = append(deleteObjectPkgTo, char.(*Character).ObjectId)
				return true
			})
		}
	}
	if len(deleteObjectPkgTo) > 0 {
		c.DeleteObjectTo <- deleteObjectPkgTo
	}

	// кому отправить charInfo
	charInfoPkgTo := make([]int32, 0, 64)
	npcPkgTo := make([]Npc, 0, 64)
	for _, region := range newAreas {
		if !Contains(oldAreas, region) {
			region.CharsInRegion.Range(func(objId, char interface{}) bool {
				if char.(*Character).ObjectId == c.ObjectId {
					return true
				}
				charInfoPkgTo = append(charInfoPkgTo, char.(*Character).ObjectId)
				return true
			})

			region.NpcInRegion.Range(func(objId, npc interface{}) bool {
				npcPkgTo = append(npcPkgTo, npc.(Npc))
				return true
			})
		}
	}
	if len(charInfoPkgTo) > 0 {
		c.CharInfoTo <- charInfoPkgTo
	}
	c.CurrentRegion = newRegion

	if len(npcPkgTo) > 0 {
		c.NpcInfo <- npcPkgTo
	}

}

func (c *Character) CheckRegion() {
	for {
		time.Sleep(time.Second)
		if c.CurrentRegion != nil {
			curReg := c.CurrentRegion
			x, y, z := c.GetXYZ()
			ncurReg := GetRegion(x, y, z)
			if curReg != ncurReg {
				c.setWorldRegion(ncurReg)
			}
		}
	}

}
