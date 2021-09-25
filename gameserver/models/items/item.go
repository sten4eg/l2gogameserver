package items

import (
	"encoding/json"
	"l2gogameserver/gameserver/models/items/armorType"
	"l2gogameserver/gameserver/models/items/crystalType"
	"l2gogameserver/gameserver/models/items/etcItemType"
	"l2gogameserver/gameserver/models/items/materialType"
	"l2gogameserver/gameserver/models/items/weaponType"
	"os"
)

type Item struct {
	Id                     int                       `json:"id"`
	ItemType               ItemType                  `json:"itemType"`
	Name                   string                    `json:"name"`
	Icon                   string                    `json:"icon"`
	SlotBitType            SlotBitType               `json:"slot_bit_type"`
	ArmorType              armorType.ArmorType       `json:"armor_type"`
	EtcItemType            etcItemType.EtcItemType   `json:"etcitem_type"`
	ItemMultiSkillList     []string                  `json:"item_multi_skill_list"`
	RecipeId               int                       `json:"recipe_id"`
	Weight                 int                       `json:"weight"`
	ConsumeType            string                    `json:"consume_type"`
	SoulShotCount          int                       `json:"soulshot_count"`
	SpiritShotCount        int                       `json:"spiritshot_count"`
	DropPeriod             int                       `json:"drop_period"`
	DefaultPrice           int                       `json:"default_price"`
	ItemSkill              string                    `json:"item_skill"`
	CriticalAttackSkill    string                    `json:"critical_attack_skill"`
	AttackSkill            string                    `json:"attack_skill"`
	MagicSkill             string                    `json:"magic_skill"`
	ItemSkillEnchantedFour string                    `json:"item_skill_enchanted_four"`
	MaterialType           materialType.MaterialType `json:"material_type"`
	CrystalType            crystalType.CrystalType   `json:"crystal_type"`
	CrystalCount           int                       `json:"crystal_count"`
	IsTrade                bool                      `json:"is_trade"`
	IsDrop                 bool                      `json:"is_drop"`
	IsDestruct             bool                      `json:"is_destruct"`
	IsPrivateStore         bool                      `json:"is_private_store"`
	KeepType               int                       `json:"keep_type"`
	RandomDamage           int                       `json:"random_damage"`
	WeaponType             weaponType.WeaponType     `json:"weapon_type"`
	HitModify              int                       `json:"hit_modify"`
	AvoidModify            int                       `json:"avoid_modify"`
	ShieldDefense          int                       `json:"shield_defense"`
	ShieldDefenseRate      int                       `json:"shield_defense_rate"`
	AttackRange            int                       `json:"attack_range"`
	ReuseDelay             int                       `json:"reuse_delay"`
	MpConsume              int                       `json:"mp_consume"`
	Durability             int                       `json:"durability"`
	MagicWeapon            bool                      `json:"magic_weapon"`
	EnchantEnable          bool                      `json:"enchant_enable"`
	ElementalEnable        bool                      `json:"elemental_enable"`
	ForNpc                 bool                      `json:"for_npc"`
	IsOlympiadCanUse       bool                      `json:"is_olympiad_can_use"`
	IsPremium              bool                      `json:"is_premium"`
	BonusStats             []ItemBonusStat           `json:"stats"`
	DefaultAction          DefaultAction             `json:"default_action"`
	InitialCount           int                       `json:"initial_count"`
	ImmediateEffect        int                       `json:"immediate_effect"`
	CapsuledItems          []CapsuledItem            `json:"capsuled_items"`
	DualFhitRate           int                       `json:"dual_fhit_rate"`
	DamageRange            int                       `json:"damage_range"`
	Enchanted              int                       `json:"enchanted"`
	BaseAttributeAttack    BaseAttributeAttack       `json:"base_attribute_attack"`
	BaseAttributeDefend    BaseAttributeDefend       `json:"base_attribute_defend"`
	UnequipSkill           []string                  `json:"unequip_skill"`
	ItemEquipOption        []string                  `json:"item_equip_option"`
	CanMove                bool                      `json:"can_move"`
	DelayShareGroup        int                       `json:"delay_share_group"`
	Blessed                int                       `json:"blessed"`
	ReducedSoulshot        []string                  `json:"reduced_soulshot"`
	ExImmediateEffect      int                       `json:"ex_immediate_effect"`
	UseSkillDistime        int                       `json:"use_skill_distime"`
	Period                 int                       `json:"period"`
	EquipReuseDelay        int                       `json:"equip_reuse_delay"`
	Price                  int                       `json:"price"`
}

// AllItems - ONLY READ MAP, set in init server
var AllItems map[int]Item

func LoadItems() {
	AllItems = make(map[int]Item)
	loadItems()
}

func loadItems() {
	file, err := os.Open("./data/stats/items/items.json")
	if err != nil {
		panic("Failed to load config file")
	}

	var items []Item

	err = json.NewDecoder(file).Decode(&items)

	if err != nil {
		panic("Ошибка при чтении с файла items.json. " + err.Error())
	}

	for _, v := range items {
		v.removeEmptyStats()
		AllItems[v.Id] = v
	}

}
func (i *Item) removeEmptyStats() {
	var bStat []ItemBonusStat
	for _, v := range i.BonusStats {
		if v.Val != 0 {
			bStat = append(bStat, v)
		}
	}
	i.BonusStats = bStat
}

func GetItemFromStorage(itemId int) (item Item, ok bool) {
	item, ok = AllItems[itemId]
	return
}
