package models

import (
	"encoding/json"
	"io/fs"
	"math/rand"
	"os"
	"strconv"
)

var AllStats map[int]Stats

type Stats struct {
	ClassId        int
	StaticData     StaticData
	CreationPoints []CreationPoint
	LvlUpgainData  []LvlUpgainData
	ClassTree      *ClassTree
}
type StaticData struct {
	INT              int
	STR              int
	CON              int
	MEN              int
	DEX              int
	WIT              int
	BasePAtk         int
	BaseCritRate     int
	BaseAtkType      string
	BasePAtkSpd      int
	BasePDef         BasePDef
	BaseMAtk         int
	BaseMDef         BaseMDef
	BaseCanPenetrate int
	BaseAtkRange     int
	BaseDamRange     BaseDamRange
	BaseRndDam       int
	BaseMoveSpd      BaseMoveSpd
	BaseBreath       int
	BaseSafeFall     int
	CollisionMale    Collision
	CollisionFemale  Collision
}
type Collision struct {
	Radius int
	Height int
}
type BaseMoveSpd struct {
	Walk     int
	Run      int
	SlowSwim int
	FastSwim int
}
type BaseDamRange struct {
	VerticalDirection   int
	HorizontalDirection int
	Distance            int
	Width               int
}
type BaseMDef struct {
	Rear    int
	Lear    int
	Rfinger int
	Lfinger int
	Neck    int
}
type BasePDef struct {
	Chest     int
	Legs      int
	Head      int
	Feet      int
	Gloves    int
	Underwear int
	Cloak     int
}
type CreationPoint struct {
	X int
	Y int
	Z int
}
type LvlUpgainData struct {
	Level   int
	Hp      float32
	Mp      float32
	Cp      float32
	HpRegen float32
	MpRegen float32
	CpRegen float32
}

type Class struct {
	ClassId       int
	Name          string
	ServerName    string
	ParentClassId int
}

func LoadStats() {
	file, err := os.Open("./data/stats/char/classList.json")
	if err != nil {
		panic(err)
	}
	classes := make([]Class, 0, 107)
	err = json.NewDecoder(file).Decode(&classes)
	if err != nil {
		panic(err)
	}

	AllStats = make(map[int]Stats)

	fss := os.DirFS("./data/stats/char/baseStats")
	dir, err := fs.ReadDir(fss, ".")
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		loadStat(entry, classes)
	}
}

func loadStat(entry fs.DirEntry, classes []Class) {
	file, err := os.Open("./data/stats/char/baseStats/" + entry.Name())
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}

	stats := make([]Stats, 0, 1)

	err = json.NewDecoder(file).Decode(&stats)
	if err != nil {
		panic("Failed to decode config file " + file.Name() + " " + err.Error())
	}

	stats[0].ClassTree = findClassTree(stats[0].ClassId, classes)

	AllStats[stats[0].ClassId] = stats[0]
}

type ClassTree struct {
	ClassId    int
	Name       string
	ServerName string
	Child      *ClassTree
}

func findClassTree(classId int, classes []Class) *ClassTree {
	var ct ClassTree

	for _, v := range classes {
		if v.ClassId == classId {
			ct.ServerName = v.ServerName
			ct.Name = v.Name
			ct.ClassId = classId

			if v.ParentClassId != -1 {
				ct.Child = findClassTree(v.ParentClassId, classes)
			}

			return &ct
		}

	}

	return nil
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
