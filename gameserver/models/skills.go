package models

import (
	"context"
	"encoding/json"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/models/skills"
	"l2gogameserver/gameserver/models/skills/targets"
	"log"
	"os"
	"strconv"
)

type Skill struct {
	ID          int                `json:"id"`
	Levels      int                `json:"levels"`
	Name        string             `json:"name"`
	Power       int                `json:"power"`
	CastRange   int                `json:"castRange"`
	CoolTime    int                `json:"coolTime"`
	HitTime     int                `json:"hitTime"`
	OverHit     bool               `json:"overHit"`
	ReuseDelay  int                `json:"reuseDelay"`
	OperateType skills.OperateType `json:"operateType"`
	TargetType  targets.TargetType `json:"targetType"`
	IsMagic     int                `json:"isMagic"`
	MagicLvl    int                `json:"magicLvl"`
	MpConsume1  int                `json:"mpConsume1"`
	MpConsume2  int                `json:"mpConsume2"`
}

type SkillForParseJSON struct {
	ID          int                `json:"id"`
	Levels      int                `json:"levels"`
	Name        string             `json:"name"`
	Power       []int              `json:"power"`
	CastRange   int                `json:"castRange"`
	CoolTime    int                `json:"coolTime"`
	HitTime     int                `json:"hitTime"`
	OverHit     bool               `json:"overHit"`
	ReuseDelay  int                `json:"reuseDelay"`
	OperateType skills.OperateType `json:"operateType"`
	TargetType  targets.TargetType `json:"targetType"`
	IsMagic     int                `json:"isMagic"`
	MagicLvl    []int              `json:"magicLvl"`
	MpConsume1  []int              `json:"mpConsume1"`
	MpConsume2  []int              `json:"mpConsume2"`
}

var AllSkills map[Tuple]Skill

type Tuple struct {
	Id  int
	Lvl int
}

func LoadSkills() {
	if config.Get().Debug.EnabledSkills == false {
		return
	}
	logger.Info.Println("Загрузка скиллов")
	file, err := os.Open("./datapack/data/stats/skills/0-100.json")
	if err != nil {
		logger.Error.Panicln("Failed to load config file " + err.Error())
	}

	decoder := json.NewDecoder(file)

	var skillsJson []SkillForParseJSON

	err = decoder.Decode(&skillsJson)
	if err != nil {
		logger.Error.Panicln("Failed to decode config file " + file.Name() + " " + err.Error())
	}
	AllSkills = make(map[Tuple]Skill)

	for _, v := range skillsJson {
		fSkill := Skill{
			ID:          v.ID,
			Levels:      1,
			Name:        v.Name,
			Power:       v.Power[0],
			CastRange:   v.CastRange,
			CoolTime:    v.CoolTime,
			HitTime:     v.HitTime,
			OverHit:     v.OverHit,
			ReuseDelay:  v.ReuseDelay,
			OperateType: v.OperateType,
			TargetType:  v.TargetType,
			IsMagic:     v.IsMagic,
			MagicLvl:    v.MagicLvl[0],
			MpConsume1:  v.MpConsume1[0],
			MpConsume2:  v.MpConsume2[0],
		}

		if v.Levels > 1 {
			for i := 0; i < v.Levels; i++ {
				fSkill.Levels = i
				fSkill.Power = v.Power[i]
				AllSkills[Tuple{v.ID, i}] = fSkill
			}
		} else {
			AllSkills[Tuple{v.ID, v.Levels}] = fSkill
		}
	}
	qw := AllSkills
	_ = qw
}

func GetMySkills(charId int32) []Skill {
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT skill_id, skill_level FROM character_skills WHERE char_id = $1", charId)
	if err != nil {
		logger.Error.Panicln(err)
	}

	var skills []Skill
	for rows.Next() {
		var skl Tuple

		err = rows.Scan(&skl.Id, &skl.Lvl)
		if err != nil {
			log.Println(err)
		}
		sk, ok := AllSkills[skl]
		if !ok {
			logger.Error.Panicln("not found Skill")
		}
		skills = append(skills, sk)
	}
	return skills
}

func (c *Character) LoadSkills() {
	c.Skills = map[int]Skill{}
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT skill_id,skill_level FROM character_skills WHERE char_id=$1 AND class_id=$2", c.ObjectId, c.ClassId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var t Tuple
		err = rows.Scan(&t.Id, &t.Lvl)
		if err != nil {
			panic(err)
		}

		sk, ok := AllSkills[t]
		if !ok {
			logger.Error.Panicln("Скилл персонажа " + c.CharName + " не найден в мапе скиллов id: " + strconv.Itoa(t.Id) + " Level: " + strconv.Itoa(t.Lvl))
		}
		c.Skills[sk.ID] = sk //= append(c.Skills, sk)
	}

}
