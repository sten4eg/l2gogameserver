package models

import (
	"encoding/json"

	"os"
)

var AllStats map[string]Stats

type Stats struct {
	BaseSTR      int32        `json:"baseSTR"`
	BaseDEX      int32        `json:"baseDEX"`
	BaseCON      int32        `json:"baseCON"`
	BaseINT      int32        `json:"baseINT"`
	BaseWIT      int32        `json:"baseWIT"`
	BaseMEN      int32        `json:"baseMEN"`
	BaseCritRate int32        `json:"baseCritRate"`
	BaseAtkType  string       `json:"baseAtkType"`
	BasePAtkSpd  int32        `json:"basePAtkSpd"`
	BasePDef     basePDef     `json:"basePDef"`
	BaseMAtk     int32        `json:"baseMAtk"`
	BaseMDef     baseMDef     `json:"baseMDef"`
	BaseAtkRange int32        `json:"baseAtkRange"`
	BaseDamRange baseDamRange `json:"baseDamRange"`
	BaseRndDam   int32        `json:"baseRndDam"`
	BaseMoveSpd  baseMoveSpd  `json:"baseMoveSpd"`
}
type basePDef struct {
	Chest     int32 `json:"chest"`
	Legs      int32 `json:"legs"`
	Head      int32 `json:"head"`
	Feet      int32 `json:"feet"`
	Gloves    int32 `json:"gloves"`
	Underwear int32 `json:"underwear"`
	Cloak     int32 `json:"cloak"`
}
type baseMoveSpd struct {
	Walk int32 `json:"walk"`
	Run  int32 `json:"run"`
	Swim int32 `json:"swim"`
}
type baseDamRange struct {
	VerticalDirection   int32 `json:"verticalDirection"`
	HorizontalDirection int32 `json:"horizontalDirection"`
	Distance            int32 `json:"distance"`
	Width               int32 `json:"width"`
}

type baseMDef struct {
	Rear    int32 `json:"rear"`
	Lear    int32 `json:"lear"`
	Rfinger int32 `json:"rfinger"`
	Lfinger int32 `json:"lfinger"`
	Neck    int32 `json:"neck"`
}

func LoadStats() {
	file, err := os.Open("./data/stats/char/baseStats/humanFighter.json")
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}

	decoder := json.NewDecoder(file)

	var statsJson Stats

	err = decoder.Decode(&statsJson)
	if err != nil {
		panic("Failed to decode config file " + file.Name() + " " + err.Error())
	}
	AllStats = make(map[string]Stats)

	AllStats["humF"] = statsJson

}
