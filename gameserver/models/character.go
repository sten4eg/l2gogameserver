package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/pgtype"
	"l2gogameserver/data"
	"l2gogameserver/gameserver/dto"
	"l2gogameserver/gameserver/models/items"
	"math/rand"
	"os"
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
	Race            int32
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
	Stats           Stats
	pvpFlag         bool
	ShortCut        map[int32]dto.ShortCutDTO
	ActiveSoulShots []int32
	IsDead          bool
	IsFakeDeath     bool
	// Skills todo: проверить слайс или мапа лучше для скилов
	Skills       map[int]Skill
	IsCastingNow bool
	SkillQueue   chan SkillHolder
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
		default:
		}
	}
}

type SkillHolder struct {
	skill        Skill
	ctrlPressed  bool
	shiftPressed bool
}

func (c *Character) SetSkillToQueue(skill Skill, ctrlPressed, shiftPressed bool) {
	s := SkillHolder{
		skill:        skill,
		ctrlPressed:  ctrlPressed,
		shiftPressed: shiftPressed,
	}
	c.SkillQueue <- s
}
func SetupStats(char *Character) {
	if char.BaseClass == 0 {
		char.Stats = AllStats["humF"]
	}
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

type StartLocation struct {
	ClassId int32
	Spawn   []Coordinates
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
func GetCreationCoordinates(classId int32) *Coordinates {

	var config []StartLocation
	file, err := os.Open("./data/stats/char/pcCreationPoint.json")
	if err != nil {
		panic("Failed to load config file")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic("Failed to decode config file")
	}

	var coordinates Coordinates
	for _, v := range config {
		if v.ClassId == classId {
			/* #nosec */
			rnd := rand.Intn(len(v.Spawn))
			v.Spawn[rnd].mu.Lock()
			coordinates.mu.Lock()

			coordinates.X = v.Spawn[rnd].X
			coordinates.Y = v.Spawn[rnd].Y
			coordinates.Z = v.Spawn[rnd].Z

			v.Spawn[rnd].mu.Unlock()
			coordinates.mu.Unlock()
		}
	}
	return &coordinates
}

// Load загрузка персонажа
func (c *Character) Load() {
	c.ShortCut = restoreMe(c.CharId, c.ClassId)
	c.LoadSkills()
	c.SkillQueue = make(chan SkillHolder)
	go c.ListenSkillQueue()
}

func (c *Character) checkSoulShot() {
	if len(c.ActiveSoulShots) == 0 {
		return
	}

}
