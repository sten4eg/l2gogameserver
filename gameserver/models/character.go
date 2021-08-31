package models

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/pgtype"
	"l2gogameserver/data"
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/race"
	"math/rand"
	"strconv"
	"sync"
)

type Character struct {
	Login           pgtype.Bytea
	CharId          int32
	Level           int32
	MaxHp           int32
	CurHp           int32
	MaxMp           int32
	CurMp           int32
	Face            int32
	HairStyle       int32
	HairColor       int32
	Sex             int32
	Coordinates     *Coordinates
	Exp             int32
	Sp              int32
	Karma           int32
	PvpKills        int32
	PkKills         int32
	ClanId          int32
	Race            race.Race
	ClassId         int32
	BaseClass       int32
	Title           sql.NullString
	OnlineTime      int32
	Nobless         int32
	Vitality        int32
	CharName        pgtype.Bytea
	CurrentRegion   *WorldRegion
	Conn            *Client
	AttackEndTime   int64
	Paperdoll       [31][3]int32
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
	Inventory              []items.MyItem
	CursedWeaponEquippedId int
}

func GetNewCharacterModel() *Character {
	character := new(Character)
	sk := make(map[int]Skill)
	character.Skills = sk
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

func (c *Character) SetSkillToQueue(skill Skill, ctrlPressed, shiftPressed bool) {
	s := SkillHolder{
		Skill:        skill,
		CtrlPressed:  ctrlPressed,
		ShiftPressed: shiftPressed,
	}
	c.SkillQueue <- s
}
func SetupStats(char *Character) {
	char.Stats = AllStats[int(char.ClassId)].StaticData //todo а для чего BaseClass ??
}

type Account struct {
	Char     []*Character
	CharSlot int32
	Login    string
}

type Coordinates struct {
	mu sync.Mutex
	X  int32
	Y  int32
	Z  int32
}

type PacketByte struct {
	B []byte
}

// IsActiveWeapon есть ли у персонажа оружие в руках
func (c *Character) IsActiveWeapon() bool {
	x := c.Paperdoll[items.PAPERDOLL_RHAND]
	return x[0] == 0
}

// GetB получение массива байт в PacketByte
func (i *PacketByte) GetB() []byte {
	cl := make([]byte, len(i.B))
	_ = copy(cl, i.B)
	return cl
}

// SetB копирует массив байт в PacketByte
func (i *PacketByte) SetB(v []byte) {
	cl := make([]byte, len(v))
	i.B = cl
	copy(i.B, v)
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

// GetCreationCoordinates получить рандомные координаты при создании
// персонажа, зависит от classId создаваемого персонажа
func GetCreationCoordinates(classId int32) (int, int, int) {

	e, ok := AllStats[int(classId)]
	if !ok {
		panic("не найдена информация в AllStats по classId: " + strconv.Itoa(int(classId)))
	}

	rnd := rand.Intn(len(e.CreationPoints))

	x := e.CreationPoints[rnd].X
	y := e.CreationPoints[rnd].Y
	z := e.CreationPoints[rnd].Z
	return x, y, z

}

// Load загрузка персонажа
func (c *Character) Load() {
	c.ShortCut = restoreMe(c.CharId, c.ClassId)
	c.LoadSkills()
	c.SkillQueue = make(chan SkillHolder)
	c.Inventory = items.GetMyItems(c.CharId)
	go c.ListenSkillQueue()
}

func (c *Character) checkSoulShot() {
	if len(c.ActiveSoulShots) == 0 {
		return
	}
}

func (c *Character) IsCursedWeaponEquipped() bool {
	return c.CursedWeaponEquippedId != 0
}
