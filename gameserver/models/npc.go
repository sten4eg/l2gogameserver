package models

import (
	"encoding/json"
	"errors"
	"l2gogameserver/config"
	"l2gogameserver/gameserver/idfactory"
	"log"
	"os"
)

type Npc struct {
	Type                string              `json:"type"`
	NpcId               int32               `json:"npcid"`
	Name                string              `json:"name"`
	Level               int                 `json:"level"`
	Exp                 int                 `json:"exp"`
	ExCrtEffect         int                 `json:"ex_crt_effect"`
	Unique              int                 `json:"unique"`
	SNpcPropHpRate      int                 `json:"s_npc_prop_hp_rate"`
	Race                string              `json:"race"`
	Sex                 string              `json:"sex"`
	SkillList           []string            `json:"skill_list"`
	SlotRhand           int32               `json:"slot_rhand"`
	SlotLhand           int32               `json:"slot_lhand"`
	CollisionRadius     float64             `json:"collision_radius"`
	CollisionHeight     float64             `json:"collision_height"`
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
	EventFlag               int                       `json:"event_flag"`
	Unsowing                int                       `json:"unsowing"`
	PrivateRespawnLog       int                       `json:"private_respawn_log"`
	AcquireExpRate          float64                   `json:"acquire_exp_rate"`
	AcquireSp               int                       `json:"acquire_sp"`
	AcquireRp               int                       `json:"acquire_rp"`
	CorpseMakeList          []CorpseMakeList          `json:"corpse_make_list"`
	AdditionalMakeList      []AdditionalMakeList      `json:"additional_make_list"`
	AdditionalMakeMultiList []AdditionalMakeMultiList `json:"additional_make_multi_list"`
	ExItemDropList          []ExItemDropList          `json:"ex_item_drop_list"`
	FakeClassId             int                       `json:"fake_class_id"`
	Locations               []Locations               `json:"locations"`
	ObjId                   int32
	Spawn                   Locations
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
type AdditionalMakeList struct {
}
type Locations struct {
	NpcId         int32 `json:"npc_id"`
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

// Npcs Список всех NPC map[NPCID]map[OBJECTID]Npc
var Npcs map[int32]map[int32]Npc

//	Список объектов, map[OBJECTID]map[Location]
var NpcObject map[int32]Locations

//Временное функция подгрузки листа с спаунами NPC
func LoadNpc() {
	if config.Get().Debug.EnableNPC == false {
		return
	}
	Npcs = make(map[int32]map[int32]Npc)
	NpcObject = make(map[int32]Locations)

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
		if len(p.Locations) == 0 {
			continue
		}
		tmp := make(map[int32]Npc)
		for _, vv := range p.Locations {
			objId := idfactory.GetNext()
			p.ObjId = objId
			p.Spawn = vv
			tmp[objId] = p

			NpcObject[objId] = Locations{
				NpcId:         p.NpcId,
				Locx:          vv.Locx,
				Locy:          vv.Locy,
				Locz:          vv.Locz,
				Randomx:       vv.Randomx,
				Randomy:       vv.Randomy,
				Heading:       vv.Heading,
				RespawnDelay:  vv.RespawnDelay,
				RespawnRandom: vv.RespawnRandom,
				LocId:         vv.LocId,
				PeriodOfDay:   vv.PeriodOfDay,
			}

		}

		Npcs[p.NpcId] = tmp
	}

	log.Printf("Загружено %d Npc", len(Npcs))
	log.Printf("Загружено %d Npc Object", len(NpcObject))

	if config.Get().Debug.EnabledSpawnlist == false {
		return
	}
	file, err = os.Open("./data/stats/npcdata/spawnlist.json")
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}
	var npcSpawn []Locations
	jsonParser = json.NewDecoder(file)
	if err = jsonParser.Decode(&npcSpawn); err != nil {
		panic("parsing config file" + err.Error())
	}

	for _, v := range Npcs {
		for _, vv := range v {
			reg := GetRegion(vv.Spawn.Locx, vv.Spawn.Locy, vv.Spawn.Locz)
			reg.AddVisibleNpc(vv)
		}
	}

}

// 0 - нпц с диалогом
// 1 - нпц монстер/рб...
//Необходимо для того чтоб понимать, это моб, или NPC диалога
func GetDialogNPC(npctype string) int32 {
	//Список не полный
	//типы нпц которые при обращении открывают HTML диалоги
	npcDialogs := []string{"citizen", "guild_coach", "guild_master", "teleporter", "merchant", "guard"}
	for _, dialog := range npcDialogs {
		if npctype == dialog {
			return 0
		}
	}
	return 1
}

//Информация о объекте
func GetNpcObject(objectId int32) (Npc, int32, int32, int32, error) {
	for npcObjId, npc := range NpcObject {
		if objectId == npcObjId {
			npcInfo, err := GetNpc(npc.NpcId, objectId)
			if err != nil {
				log.Println(err)
				return Npc{}, 0, 0, 0, err
			}
			return npcInfo, npc.Locx, npc.Locy, npc.Locz, nil
		}
	}
	return Npc{}, 0, 0, 0, errors.New("Not find object")
}

// Информация о NPC
func GetNpc(getNpcID, objectId int32) (Npc, error) {
	for npcId, npcinfo := range Npcs {
		if npcId == getNpcID {
			for _, npc := range npcinfo {
				if npc.ObjId == objectId {
					return npc, nil
				}
			}
			return Npc{}, errors.New("Not find npc#2")
		}
	}
	return Npc{}, errors.New("Not find npc#1")
}
