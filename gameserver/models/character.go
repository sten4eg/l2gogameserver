package models

import (
	"database/sql"
	"encoding/json"
	"github.com/jackc/pgx/pgtype"
	"l2gogameserver/data"
	"log"
	"math/rand"
	"os"
)

type Character struct {
	Login       pgtype.Bytea
	CharId      int32
	Level       int32
	MaxHp       int32
	CurHp       int32
	MaxMp       int32
	CurMp       int32
	Face        int32
	HairStyle   int32
	HairColor   int32
	Sex         int32
	Coordinates *Coordinates
	Exp         int32
	Sp          int32
	Karma       int32
	PvpKills    int32
	PkKills     int32
	ClanId      int32
	Race        int32
	ClassId     int32
	BaseClass   int32
	Title       sql.NullString
	OnlineTime  int32
	Nobless     int32
	Vitality    int32
	CharName    pgtype.Bytea
}

type Account struct {
	Char     []*Character
	CharSlot int32
	Login    string
}

type Coordinates struct {
	X int32
	Y int32
	Z int32
}

type StartLocation struct {
	ClassId int32
	Spawn   []Coordinates
}

func (c *Character) GetPercentFromCurrentLevel(exp, level int32) float64 {
	expPerLevel, expPerLevel2 := data.GetExpData(level)
	return float64(int64(exp)-expPerLevel) / float64(expPerLevel2-expPerLevel)
}
func GetCreationCoordinates(classId int32) *Coordinates {

	var config []StartLocation
	file, err := os.Open("./data/stats/char/pcCreationPoint.json")
	if err != nil {
		log.Fatal("Failed to load config file")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Failed to decode config file")
	}

	var coordinates Coordinates
	for _, v := range config {
		if v.ClassId == classId {
			coordinates = v.Spawn[rand.Intn(len(v.Spawn))]
		}
	}
	return &coordinates
}
