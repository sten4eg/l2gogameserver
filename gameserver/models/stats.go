package models

import (
	"encoding/json"
	"log"
	"os"
)

var AllStats map[string]Stats

type Stats struct {
	Str int32 `json:"str"`
	Dex int32 `json:"dex"`
	Con int32 `json:"con"`
	Int int32 `json:"int"`
	Wit int32 `json:"wit"`
	Men int32 `json:"men"`
}

func LoadStats() {
	file, err := os.Open("./data/stats/char/baseStats/humanFighter.json")
	if err != nil {
		log.Fatal("Failed to load config file " + err.Error())
	}

	decoder := json.NewDecoder(file)

	var skillsJson Stats

	err = decoder.Decode(&skillsJson)
	if err != nil {
		log.Fatal("Failed to decode config file " + file.Name() + " " + err.Error())
	}
	AllStats = make(map[string]Stats)

	AllStats["humF"] = skillsJson

}
