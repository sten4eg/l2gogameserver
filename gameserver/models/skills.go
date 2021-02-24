package models

import (
	"encoding/json"
	"github.com/jackc/pgx"
	"log"
	"os"
)

type Skill struct {
	ID           int    `json:"id"`
	Levels       []int  `json:"levels"`
	Name         string `json:"name"`
	Power        []int  `json:"power"`
	CastRange    int    `json:"castRange"`
	CoolTime     int    `json:"coolTime"`
	HitTime      int    `json:"hitTime"`
	OverHit      bool   `json:"overHit"`
	ReuseDelay   int    `json:"reuseDelay"`
	OperateType  int    `json:"operateType"`
	TargetType   string `json:"targetType"`
	IsMagic      int    `json:"isMagic"`
	MagicLvl     int    `json:"magicLvl"`
	MpConsume1   int    `json:"mpConsume1"`
	MpConsume2   int    `json:"mpConsume2"`
	CurrentLevel int
}

var AllSkills map[int]Skill

func LoadSkills() {
	file, err := os.Open("./data/stats/skills/0-100.json")
	if err != nil {
		log.Fatal("Failed to load config file " + err.Error())
	}

	decoder := json.NewDecoder(file)

	var skillsJson []Skill

	err = decoder.Decode(&skillsJson)
	if err != nil {
		log.Fatal("Failed to decode config file " + file.Name() + " " + err.Error())
	}
	AllSkills = make(map[int]Skill)

	for _, v := range skillsJson {
		AllSkills[v.ID] = v
	}

}

func GetMySkills(charId int32, db *pgx.Conn) []Skill {
	rows, err := db.Query("SELECT skill_id, skill_level FROM character_skills WHERE char_id = $1", charId)
	if err != nil {
		log.Fatal(err)
	}

	type tempSkillFromDB struct {
		SkillId    int
		SkillLevel int
	}

	var skills []Skill
	for rows.Next() {
		var itm tempSkillFromDB

		err = rows.Scan(&itm.SkillId, &itm.SkillLevel)
		if err != nil {
			log.Println(err)
		}
		sk, ok := AllSkills[itm.SkillId]
		if !ok {
			log.Fatal("not found Skill")
		}
		sk.CurrentLevel = itm.SkillLevel
		skills = append(skills, sk)
	}
	return skills
}

func (s *Skill) IsPassive() int32 {
	if s.OperateType == 3 {
		return 1
	}
	return 0
}
