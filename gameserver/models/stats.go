package models

import (
	"encoding/json"
	"io/fs"
	"os"
)

var AllStats map[int]Stats

type Stats struct {
	ClassId        int
	StaticData     StaticData
	CreationPoints []CreationPoint
	LvlUpgainData  []LvlUpgainData
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

func LoadStats() {
	AllStats = make(map[int]Stats)

	fss := os.DirFS("./data/stats/char/baseStats")
	r, err := fs.ReadDir(fss, ".")
	if err != nil {
		panic(err)
	}

	for _, v := range r {
		loadStat(v)
	}
}

func loadStat(entry fs.DirEntry) {
	file, err := os.Open("./data/stats/char/baseStats/" + entry.Name())
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}

	decoder := json.NewDecoder(file)

	stats := make([]Stats, 0, 1)

	err = decoder.Decode(&stats)
	if err != nil {
		panic("Failed to decode config file " + file.Name() + " " + err.Error())
	}

	AllStats[stats[0].ClassId] = stats[0]
}
