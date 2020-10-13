package main

import (
	"encoding/json"
	"l2gogameserver/data"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
	"log"
	"os"
)

type CrystalType struct {
	Id                        int
	CrystalId                 int
	CrystalEnchantBonusArmor  int
	CrystalEnchantBonusWeapon int32
}
type Item struct {
	Id              int
	DisplayId       int
	Name            string
	Icon            string
	Weight          int
	MaterialType    int
	EquipReuseDelay int
	Duration        int
	Time            int
	AutoDestroyTime int
	BodyPart        string //struct
	ReferencePrice  int
	CrystalType     CrystalType
	CrystalCount    int

	Stackable               bool
	Sellable                bool
	Dropable                bool
	Destroyable             bool
	Tradeable               bool
	Depositable             bool
	Elementable             bool
	Enchantable             int
	QuestItem               bool
	Freightable             bool
	Allow_self_resurrection bool
	Is_oly_restricted       bool
	For_npc                 bool

	Immediate_effect    bool
	Ex_immediate_effect bool

	DefaultAction       int
	UseSkillDisTime     int
	DefaultEnchantLevel int
	ReuseDelay          int
	SharedReuseGroup    int

	//Common = ((_itemId >= 11605) && (_itemId <= 12361));
	//HeroItem = ((_itemId >= 6611) && (_itemId <= 6621)) || ((_itemId >= 9388) && (_itemId <= 9390)) || (_itemId == 6842);
	//PvpItem = ((_itemId >= 10667) && (_itemId <= 10835)) || ((_itemId >= 12852) && (_itemId <= 12977)) || ((_itemId >= 14363) && (_itemId <= 14525)) || (_itemId == 14528) || (_itemId == 14529) || (_itemId == 14558) || ((_itemId >= 15913) && (_itemId <= 16024))
	//|| ((_itemId >= 16134) && (_itemId <= 16147)) || (_itemId == 16149) || (_itemId == 16151) || (_itemId == 16153) || (_itemId == 16155) || (_itemId == 16157) || (_itemId == 16159) || ((_itemId >= 16168) && (_itemId <= 16176)) || ((_itemId >= 16179) && (_itemId <= 16220));
}

type Items struct {
	Itemss Item
}

func main() {
	//defer profile.Start().Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfileHeap()).Stop()

	file, err := os.Open("./data/stats/items/items.json")
	if err != nil {
		log.Fatal("Failed to load config file")
	}
	var Items []Item

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Items)
	if err != nil {
		log.Fatal("Failed to decode config file")
	}

	models.NewWorld()
	data.Load()
	x := gameserver.New()
	x.Init()
	x.Start()
}
