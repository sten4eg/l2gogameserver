package models

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
)

type Npc struct {
	Type                string              `json:"type"`
	NpcId               int32                 `json:"npcid"`
	Name                string              `json:"name"`
	Level               int                 `json:"level"`
	Exp                 int                 `json:"exp"`
	ExCrtEffect         int                 `json:"ex_crt_effect"`
	Unique              int                 `json:"unique"`
	SNpcPropHpRate      int                 `json:"s_npc_prop_hp_rate"`
	Race                string              `json:"race"`
	Sex                 string              `json:"sex"`
	SkillList           []string            `json:"skill_list"`
	SlotRhand           string              `json:"slot_rhand"`
	SlotLhand           interface{}         `json:"slot_lhand"`
	CollisionRadius     float64              `json:"collision_radius"`
	CollisionHeight     float64              `json:"collision_height"`
	HitTimeFactor       float64             `json:"hit_time_factor"`
	HitTimeFactorSkill  int                 `json:"hit_time_factor_skill"`
	GroundHigh          int                 `json:"ground_high"`
	GroundLow           int                 `json:"ground_low"`
	Str                 int                 `json:"str"`
	Int                 int                 `json:"int"`
	Dex                 int                 `json:"dex"`
	Wit                 int                 `json:"wit"`
	Con                 int                 `json:"con"`
	Men                 int                 `json:"men"`
	OrgHp               float64             `json:"org_hp"`
	OrgHpRegen          float64             `json:"org_hp_regen"`
	OrgMp               float64             `json:"org_mp"`
	OrgMpRegen          float64             `json:"org_mp_regen"`
	BaseAttackType      string              `json:"base_attack_type"`
	BaseAttackRange     int                 `json:"base_attack_range"`
	BaseDamageRange     BaseDamageRange     `json:"base_damage_range"`
	BaseRandDam         int                 `json:"base_rand_dam"`
	BasePhysicalAttack  float64             `json:"base_physical_attack"`
	BaseCritical        int                 `json:"base_critical"`
	PhysicalHitModify   float64             `json:"physical_hit_modify"`
	BaseAttackSpeed     int                 `json:"base_attack_speed"`
	BaseReuseDelay      int                 `json:"base_reuse_delay"`
	BaseMagicAttack     float64             `json:"base_magic_attack"`
	BaseDefend          float64             `json:"base_defend"`
	BaseMagicDefend     float64             `json:"base_magic_defend"`
	BaseAttributeAttack BaseAttributeAttack `json:"base_attribute_attack"`
	BaseAttributeDefend BaseAttributeDefend `json:"base_attribute_defend"`
	PhysicalAvoidModify int                 `json:"physical_avoid_modify"`
	ShieldDefenseRate   int                 `json:"shield_defense_rate"`
	ShieldDefense       int                 `json:"shield_defense"`
	SafeHeight          int                 `json:"safe_height"`
	SoulshotCount       int                 `json:"soulshot_count"`
	SpiritshotCount     int                 `json:"spiritshot_count"`
	Clan                []string            `json:"clan"`
	ClanHelpRange       int                 `json:"clan_help_range"`
	Undying             int                 `json:"undying"`
	CanBeAttacked       int                 `json:"can_be_attacked"`
	CorpseTime          int                 `json:"corpse_time"`
	NoSleepMode         int                 `json:"no_sleep_mode"`
	AgroRange           int                 `json:"agro_range"`
	PassableDoor        int                 `json:"passable_door"`
	CanMove             int                 `json:"can_move"`
	Flying              int                 `json:"flying"`
	HasSummoner         int                 `json:"has_summoner"`
	Targetable          int                 `json:"targetable"`
	ShowNameTag         int                 `json:"show_name_tag"`
	NpcAi               struct {
		Name string `json:"name"`
		List []struct {
			Ai  string `json:"ai"`
			Val string `json:"val"`
		} `json:"list"`
	} `json:"npc_ai"`
	EventFlag               string                    `json:"event_flag"`
	Unsowing                int                       `json:"unsowing"`
	PrivateRespawnLog       int                       `json:"private_respawn_log"`
	AcquireExpRate          float64                   `json:"acquire_exp_rate"`
	AcquireSp               int                       `json:"acquire_sp"`
	AcquireRp               int                       `json:"acquire_rp"`
	CorpseMakeList          []CorpseMakeList          `json:"corpse_make_list"`
	AdditionalMakeList      string                    `json:"additional_make_list"`
	AdditionalMakeMultiList []AdditionalMakeMultiList `json:"additional_make_multi_list"`
	ExItemDropList          []ExItemDropList          `json:"ex_item_drop_list"`
	FakeClassId             string                    `json:"fake_class_id"`
	Locations               []Locations               `json:"locations"`
}
/**
Структуры NPC пока пусты ибо не подготовил их реализацию и содержимое в правильном формате
 */
type BaseDamageRange struct {
}
type BaseAttributeAttack struct {
}
type BaseAttributeDefend struct {
}
type AdditionalMakeMultiList struct {
}
type CorpseMakeList struct {
}
type ExItemDropList struct {
}

type Locations struct {
	NpcId          int32 `json:"npc_id"`
	Locx          int32 `json:"locx"`
	Locy          int32 `json:"locy"`
	Locz          int32 `json:"locz"`
	Randomx       int   `json:"randomx"`
	Randomy       int   `json:"randomy"`
	Heading       int32 `json:"heading"`
	RespawnDelay  int   `json:"respawn_delay"`
	RespawnRandom int   `json:"respawn_random"`
	LocId         int   `json:"loc_id"`
	PeriodOfDay   int   `json:"periodOfDay"`
}

//Список всех NPC
var Npcs []Npc
//Список всех локаций NPC
var Spawnlist []Locations

//Временное функция подгрузки листа с спаунами NPC
func LoadNpc() {
	log.Println("Загрузка NPC")
	file, err := os.Open("./data/stats/npcdata/npcdata.json")
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}
	var NpcData []Npc
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&NpcData); err != nil {
		panic("parsing config file" + err.Error())
	}
	for _, p := range NpcData {
		Npcs = append(Npcs, p)
	}
	log.Printf("Загружено %d Npc", len(Npcs))

	file, err = os.Open("./data/stats/npcdata/spawnlist.json")
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}
	var NpcSpawn []Locations
	jsonParser = json.NewDecoder(file)
	if err = jsonParser.Decode(&NpcSpawn); err != nil {
		panic("parsing config file" + err.Error())
	}
	for _, p := range NpcSpawn {
		Spawnlist = append(Spawnlist, p)
	}
	log.Printf("Загружено %d спаунов", len(Spawnlist))
}

//Функция возращает массив ближайших NPC к игроку
// maxDistance максимальное поинтов от NPC к игроку
func (c *Character) NpcDistancePoint(maxDistance float64) []Npc {
	playerX, playerY, _ := c.GetXYZ()
	var npcdata []Npc
	for _, npc := range Npcs {
		for _, location := range npc.Locations {
			NpcX := float64(location.Locx)
			NpcY := float64(location.Locy)
			distance := math.Sqrt(math.Pow(float64(playerX)-NpcX, 2) + math.Pow(float64(playerY)-NpcY, 2))
			if distance < maxDistance {
				npcdata = append(npcdata, npc)
			}
		}
	}
	return npcdata
}

//Получение информации о NPC
func GetNpcInfo(id int32) (Npc, error) {
	for _, npc := range Npcs {
		if npc.NpcId == id {
			return npc, nil
		}
	}
	return Npc{}, errors.New("Not find NPC " + strconv.Itoa(int(id)))
}

//Список ближайших NPC
//Параметр maxDistance указывается максимальный диапазон поиска NPC
//диапазона 3000 хватает чтоб получить список всех NPC в Talking Village (стоя в центре)
//диапазона 7000 хватает для прогрузки NPC аж до моста от Talking Village (стоя в центре)
//Если персонаж находится в центре города
func (c *Character) SpawnDistancePoint(maxDistance float64) []Locations {
	playerX, playerY, _ := c.GetXYZ()
	var npcdata []Locations
	for _, spawn := range Spawnlist {
			NpcX := float64(spawn.Locx)
			NpcY := float64(spawn.Locy)
			distance := math.Sqrt(math.Pow(float64(playerX)-NpcX, 2) + math.Pow(float64(playerY)-NpcY, 2))
			if distance < maxDistance {
				npcdata = append(npcdata, spawn)
			}
	}
	return npcdata
}
