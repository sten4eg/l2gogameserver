package models

import (
	"context"
	"encoding/json"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/models/skills"
	"log"
	"os"
)

type Skill struct {
	ID           int                `json:"id"`
	Levels       []int              `json:"levels"`
	Name         string             `json:"name"`
	Power        []int              `json:"power"`
	CastRange    int                `json:"castRange"`
	CoolTime     int                `json:"coolTime"`
	HitTime      int                `json:"hitTime"`
	OverHit      bool               `json:"overHit"`
	ReuseDelay   int                `json:"reuseDelay"`
	OperateType  skills.OperateType `json:"operateType"`
	TargetType   string             `json:"targetType"`
	IsMagic      int                `json:"isMagic"`
	MagicLvl     int                `json:"magicLvl"`
	MpConsume1   int                `json:"mpConsume1"`
	MpConsume2   int                `json:"mpConsume2"`
	CurrentLevel int
}

var AllSkills map[int]Skill

func LoadSkills() {
	file, err := os.Open("./data/stats/skills/0-100.json")
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}

	decoder := json.NewDecoder(file)

	var skillsJson []Skill

	err = decoder.Decode(&skillsJson)
	if err != nil {
		panic("Failed to decode config file " + file.Name() + " " + err.Error())
	}
	AllSkills = make(map[int]Skill)

	for _, v := range skillsJson {
		AllSkills[v.ID] = v
	}

}

func GetMySkills(charId int32) []Skill {
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}

	rows, err := dbConn.Query(context.Background(), "SELECT skill_id, skill_level FROM character_skills WHERE char_id = $1", charId)
	if err != nil {
		panic(err)
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
			panic("not found Skill")
		}
		sk.CurrentLevel = itm.SkillLevel
		skills = append(skills, sk)
	}
	return skills
}
