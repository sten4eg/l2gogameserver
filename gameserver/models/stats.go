package models

import (
	"encoding/json"
	"io/fs"
	"l2gogameserver/data/logger"
	"math/rand"
	"os"
	"strconv"
)

type (
	Stats struct {
		ClassId        int
		StaticData     StaticData
		CreationPoints []CreationPoint
		LvlUpgainData  []LvlUpgainData
		ClassTree      *ClassTree
	}

	StaticData struct {
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
	Collision struct {
		Radius int
		Height int
	}
	BaseMoveSpd struct {
		Walk     int
		Run      int
		SlowSwim int
		FastSwim int
	}
	BaseDamRange struct {
		VerticalDirection   int
		HorizontalDirection int
		Distance            int
		Width               int
	}
	BaseMDef struct {
		Rear    int
		Lear    int
		Rfinger int
		Lfinger int
		Neck    int
	}
	BasePDef struct {
		Chest     int
		Legs      int
		Head      int
		Feet      int
		Gloves    int
		Underwear int
		Cloak     int
	}
	CreationPoint struct {
		X int
		Y int
		Z int
	}
	LvlUpgainData struct {
		Level   int
		Hp      float32
		Mp      float32
		Cp      float32
		HpRegen float32
		MpRegen float32
		CpRegen float32
	}
	Class struct {
		ClassId       int
		Name          string
		ServerName    string
		ParentClassId int
	}
	ClassTree struct {
		ClassId    int
		Name       string
		ServerName string
		Child      *ClassTree
	}
)

var AllStats map[int]Stats

func LoadStats() {
	file, err := os.Open("./datapack/data/stats/char/classList.json")
	if err != nil {
		logger.Error.Panicln(err)
	}
	classes := make([]Class, 0, 107)
	err = json.NewDecoder(file).Decode(&classes)
	if err != nil {
		logger.Error.Panicln(err)
	}

	AllStats = make(map[int]Stats)

	fss := os.DirFS("./datapack/data/stats/char/baseStats")
	dir, err := fs.ReadDir(fss, ".")
	if err != nil {
		logger.Error.Panicln(err)
	}

	for _, entry := range dir {
		loadStat(entry, classes)
	}
}

func loadStat(entry fs.DirEntry, classes []Class) {
	file, err := os.Open("./datapack/data/stats/char/baseStats/" + entry.Name())
	if err != nil {
		logger.Error.Panicln("Failed to load config file " + err.Error())
	}

	stats := make([]Stats, 0, 1)

	err = json.NewDecoder(file).Decode(&stats)
	if err != nil {
		logger.Error.Panicln("Failed to decode config file " + file.Name() + " " + err.Error())
	}

	stats[0].ClassTree = findClassTree(stats[0].ClassId, classes)

	AllStats[stats[0].ClassId] = stats[0]
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
		logger.Error.Panicln("не найдена информация в AllStats по classId: " + strconv.Itoa(int(classId)))
	}

	/* #nosec */
	rnd := rand.Intn(len(e.CreationPoints))

	x := e.CreationPoints[rnd].X
	y := e.CreationPoints[rnd].Y
	z := e.CreationPoints[rnd].Z
	return x, y, z

}

func (s *Stats) GetINT() int {
	return s.StaticData.INT
}

func (s *Stats) GetSTR() int {
	return s.StaticData.STR
}

func (s *Stats) GetCON() int {
	return s.StaticData.CON
}

func (s *Stats) GetMEN() int {
	return s.StaticData.MEN
}

func (s *Stats) GetDEX() int {
	return s.StaticData.DEX
}

func (s *Stats) GetWIT() int {
	return s.StaticData.WIT
}
