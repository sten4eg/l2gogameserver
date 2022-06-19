package items

import (
	"encoding/json"
	"l2gogameserver/config"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models/items/armorType"
	"l2gogameserver/gameserver/models/items/consumeType"
	"l2gogameserver/gameserver/models/items/crystalType"
	"l2gogameserver/gameserver/models/items/etcItemType"
	"l2gogameserver/gameserver/models/items/materialType"
	"l2gogameserver/gameserver/models/items/weaponType"
	"os"
)

type Item struct {
	Id                     int `json:"id"`
	ItemType1              ItemType1
	ItemType2              ItemType2                 `json:"itemType"`
	Name                   string                    `json:"name"`
	Icon                   string                    `json:"icon"`
	SlotBitType            SlotBitType               `json:"slot_bit_type"`
	ArmorType              armorType.ArmorType       `json:"armor_type"`
	EtcItemType            etcItemType.EtcItemType   `json:"etcitem_type"`
	ItemMultiSkillList     []string                  `json:"item_multi_skill_list"`
	RecipeId               int                       `json:"recipe_id"`
	Weight                 int                       `json:"weight"`
	ConsumeType            consumeType.ConsumeType   `json:"consume_type"`
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

//var _ Item = interfaces.BaseItemInterface{}

// AllItems - ONLY READ MAP, set in init datapack
var AllItems map[int]Item

func LoadItems() {
	AllItems = make(map[int]Item)
	loadItems()
}

func GetItemInfo(id int) (Item, bool) {
	for _, item := range AllItems {
		if item.Id == id {
			return item, true
		}
	}
	return Item{}, false
}

func loadItems() {
	if config.Get().Debug.EnabledItems == false {
		return
	}
	logger.Info.Println("Загрузка предметов")
	file, err := os.Open("./datapack/data/stats/items/items.json")
	if err != nil {
		logger.Error.Panicln("Failed to load config file")
	}

	var items []Item

	err = json.NewDecoder(file).Decode(&items)

	if err != nil {
		logger.Error.Panicln("Ошибка при чтении с файла items.json. " + err.Error())
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
func (i *Item) setItemType1() {
	if (i.SlotBitType == SlotNeck) || ((i.SlotBitType & SlotLEar) != 0) || ((i.SlotBitType & SlotLFinger) != 0) || ((i.SlotBitType & SlotRBracelet) != 0) {
		i.ItemType1 = WeaponRingEarringNecklace
		i.ItemType2 = Accessory
	} else {
		if i.ArmorType == armorType.NONE && i.SlotBitType == SlotLHand {
			i.ArmorType = armorType.SHIELD
			i.ItemType1 = ShieldArmor
			i.ItemType2 = ShieldOrArmor
		}

	}

	if i.IsWeapon() {
		i.ItemType1 = WeaponRingEarringNecklace
		i.ItemType2 = Weapon
	}
}
func (i *Item) IsStackable() bool {
	return i.ConsumeType == 0
}
func (i *Item) GetId() int32 {
	return int32(i.Id)
}
func (i *Item) IsEquipable() bool {
	return !((i.SlotBitType == SlotNone) || (i.EtcItemType == etcItemType.ARROW) || (i.EtcItemType == etcItemType.BOLT) || (i.EtcItemType == etcItemType.LURE))
}
func (i *Item) IsHeavyArmor() bool {
	return i.ArmorType == armorType.HEAVY
}
func (i *Item) IsMagicArmor() bool {
	return i.ArmorType == armorType.MAGIC
}
func (i *Item) IsArmor() bool {
	return i.ItemType2 == ShieldOrArmor
}
func (i *Item) IsOnlyKamaelWeapon() bool {
	return i.WeaponType == weaponType.RAPIER || i.WeaponType == weaponType.CROSSBOW || i.WeaponType == weaponType.ANCIENTSWORD
}
func (i *Item) IsWeapon() bool {
	return i.ItemType2 == Weapon
}
func (i *Item) IsWeaponTypeNone() bool {
	return i.WeaponType == weaponType.NONE
}
func (i *Item) GetBaseItem() interfaces.BaseItemInterface {
	return i
}

func (i *Item) GetItemType1() int {
	return int(i.ItemType1)
}
func (i *Item) GetItemType2() int {
	return int(i.ItemType2)
}
func (i *Item) GetBodyPart() int32 {
	return int32(i.SlotBitType)
}
func GetItemFromStorage(itemId int) (item Item, ok bool) {
	item, ok = AllItems[itemId]
	return
}
func (i *Item) GetWeight() int {
	return i.Weight
}
