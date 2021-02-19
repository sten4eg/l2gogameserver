package models

import (
	"encoding/json"
	"log"
	"os"
)

type Skill struct {
	ID         int    `json:"id"`
	Levels     int    `json:"levels"`
	Name       string `json:"name"`
	Power      []int  `json:"power"`
	CastRange  int    `json:"castRange"`
	CoolTime   int    `json:"coolTime"`
	HitTime    int    `json:"hitTime"`
	OverHit    bool   `json:"overHit"`
	ReuseDelay int    `json:"reuseDelay"`
}

var AllSkills map[int]Skill

func LoadSkills() {
	file, err := os.Open("./data/stats/skills/0-100.json")
	if err != nil {
		log.Fatal("Failed to load config file")
	}

	decoder := json.NewDecoder(file)

	var skillsJson []Skill

	err = decoder.Decode(&skillsJson)
	if err != nil {
		log.Fatal("Failed to decode config file")
	}
	AllSkills = make(map[int]Skill)

	for _, v := range skillsJson {
		AllSkills[v.ID] = v
	}

}
