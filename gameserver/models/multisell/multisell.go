package multisell

import (
	"encoding/json"
	"l2gogameserver/gameserver/models"
	"log"
	"os"
	"strconv"
)

type Item struct {
	ID                  int          `json:"id"`
	ApplyTaxes          bool         `json:"applyTaxes,omitempty"`
	MaintainEnchantment bool         `json:"maintainEnchantment,omitempty"`
	IsTaxIngredient     bool         `json:"isTaxIngredient,omitempty"`
	TaxPercent          int          `json:"taxPercent"`
	Ingredient          []Ingredient `json:"ingredient"`
	Production          []Production `json:"production"`
}

type Ingredient struct {
	Id      int `json:"id"`
	Count   int `json:"count"`
	Enchant int `json:"enchant,omitempty"`
}
type Production struct {
	Id      int `json:"id"`
	Count   int `json:"count"`
	Enchant int `json:"enchant,omitempty"`
}

//var Multisells []Item

func LoadMultisell() {
	//log.Println("Загрузка мультиселла")
	//file, err := os.Open("./server/data/multisell/101.json")
	//if err != nil {
	//	panic("Failed to load config file " + err.Error())
	//}
	//decoder := json.NewDecoder(file)
	//err = decoder.Decode(&Multisells)
	//if err != nil {
	//	panic("Failed to decode config file " + file.Name() + " " + err.Error())
	//}

	/**
	MuL := []Item{{
		ID:                  1,
		ApplyTaxes:          false,
		MaintainEnchantment: false,
		IsTaxIngredient:     false,
		TaxPercent:          0,
		Ingredient: []Ingredient{{
			Id:    57,
			Count: 100,
		}, {
			Id:    5859,
			Count: 1,
		}},
		Production: []Production{{
			Id:    5575,
			Count: 510,
		}, {
			Id:    6036,
			Count: 14,
		}},
	},
		{
			ID:                  2,
			ApplyTaxes:          false,
			MaintainEnchantment: false,
			IsTaxIngredient:     false,
			TaxPercent:          0,
			Ingredient: []Ingredient{{
				Id:    57,
				Count: 100,
			}, {
				Id:    5859,
				Count: 1,
			}},
			Production: []Production{{
				Id:    5575,
				Count: 510,
			}, {
				Id:    6036,
				Count: 14,
			}},
		},
	}

	file, _ := json.MarshalIndent(MuL, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
	log.Println(MuL)
	*/
}

func LoadMultisellFile(ms int) []Item {
	file, err := os.Open("./server/data/multisell/" + strconv.Itoa(ms) + ".json")
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}
	var Multisells []Item
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Multisells)
	if err != nil {
		panic("Failed to decode config file " + file.Name() + " " + err.Error())
	}
	return Multisells
}

func Get(client *models.Client, id int) []Item {
	log.Println("Чтение GMShop", id)
	return LoadMultisellFile(id)
}
