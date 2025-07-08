package models

import (
	"database/sql"
	"encoding/json"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/skills"
	"l2gogameserver/gameserver/models/skills/targets"
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
type SkillHolder struct {
	Skill        Skill
	CtrlPressed  bool
	ShiftPressed bool
}

func (sh *SkillHolder) GetSkill() interfaces.SkillInterface {
	return &sh.Skill
}
func LoadSkills() {
	if !config.IsEnableSkills() {
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

func GetMySkills(charId int32, db *sql.DB) []Skill {
	rows, err := db.Query("SELECT skill_id, skill_level FROM character_skills WHERE char_id = $1", charId)
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer rows.Close()
	var skills []Skill
	for rows.Next() {
		var skl Tuple

		err = rows.Scan(&skl.Id, &skl.Lvl)
		if err != nil {
			logger.Info.Println(err)
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

	rows, err := c.Conn.db.Query("SELECT skill_id,skill_level FROM character_skills WHERE char_id=$1 AND class_id=$2", c.ObjectId, c.ClassId)
	if err != nil {
		logger.Error.Panicln(err)
	}

	for rows.Next() {
		var t Tuple
		err = rows.Scan(&t.Id, &t.Lvl)
		if err != nil {
			logger.Error.Panicln(err)
		}

		sk, ok := AllSkills[t]
		if !ok {
			logger.Error.Panicln("Скилл персонажа " + c.CharName + " не найден в мапе скиллов id: " + strconv.Itoa(t.Id) + " Level: " + strconv.Itoa(t.Lvl))
		}
		c.Skills[sk.ID] = sk //= append(c.Skills, sk)
	}

}

// Геттеры
func (s *Skill) GetId() int32 {
	return int32(s.ID)
}

func (s *Skill) IsPassive() bool {
	return s.OperateType.IsPassive()
}

func (s *Skill) GetLevel() int {
	return s.Levels
}

func (s *Skill) GetName() string {
	return s.Name
}

func (s *Skill) GetPower() int {
	return s.Power
}

func (s *Skill) GetCastRange() int {
	return s.CastRange
}

func (s *Skill) GetCoolTime() int {
	return s.CoolTime
}

func (s *Skill) GetHitTime() int {
	return s.HitTime
}

func (s *Skill) GetOverHit() bool {
	return s.OverHit
}

func (s *Skill) GetReuseDelay() int {
	return s.ReuseDelay
}

func (s *Skill) GetOperateType() int {
	return int(s.OperateType)
}

func (s *Skill) GetTargetType() int {
	return int(s.TargetType)
}

func (s *Skill) GetIsMagic() int {
	return s.IsMagic
}

func (s *Skill) GetMagicLvl() int {
	return s.MagicLvl
}

func (s *Skill) GetMpConsume1() int {
	return s.MpConsume1
}

func (s *Skill) GetMpConsume2() int {
	return s.MpConsume2
}

func (s *Skill) SetLevels(levels int) {
	s.Levels = levels
}

func (s *Skill) SetName(name string) {
	s.Name = name
}

func (s *Skill) SetPower(power int) {
	s.Power = power
}

func (s *Skill) SetCastRange(castRange int) {
	s.CastRange = castRange
}

func (s *Skill) SetCoolTime(coolTime int) {
	s.CoolTime = coolTime
}

func (s *Skill) SetHitTime(hitTime int) {
	s.HitTime = hitTime
}

func (s *Skill) SetOverHit(overHit bool) {
	s.OverHit = overHit
}

func (s *Skill) SetReuseDelay(reuseDelay int) {
	s.ReuseDelay = reuseDelay
}

func (s *Skill) SetOperateType(opType skills.OperateType) {
	s.OperateType = opType
}

func (s *Skill) SetTargetType(targetType targets.TargetType) {
	s.TargetType = targetType
}

func (s *Skill) SetIsMagic(isMagic int) {
	s.IsMagic = isMagic
}

func (s *Skill) SetMagicLvl(magicLvl int) {
	s.MagicLvl = magicLvl
}

func (s *Skill) SetMpConsume1(mpConsume int) {
	s.MpConsume1 = mpConsume
}

func (s *Skill) SetMpConsume2(mpConsume int) {
	s.MpConsume2 = mpConsume
}
