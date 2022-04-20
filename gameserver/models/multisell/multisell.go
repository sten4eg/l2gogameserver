package multisell

import (
	"encoding/json"
	"l2gogameserver/data"
	"l2gogameserver/gameserver/interfaces"
	"log"
	"os"
	"strconv"
)

type MultiList struct {
	ID     int    //id мультиселла (название (числом) файла)
	Config Config `json:"config"`
	Item   []Item `json:"items"`
}
type Config struct {
	Trader        []int `json:"trader"`        // Массив ID трейдеров продают мультиселлы
	Showall       bool  `json:"showall"`       //показывать айтемы, которые нельзя купить из-за отсутствия инградиентов
	Notax         bool  `json:"notax"`         // налогообложение
	Keepenchanted bool  `json:"keepenchanted"` //Сохраняется ли заточка при обмене
	Bbsallowed    bool  `json:"bbsallowed"`    // Разрешить вызов мультиселла из комьюнити
	Disabled      bool  `json:"disabled"`      //Отключить мутильселл
}

type Item struct {
	Ingredient []Ingredient `json:"ingredient"`
	Production []Production `json:"production"`
}

type Ingredient struct {
	Id                int  `json:"id"`
	Count             int  `json:"count"`
	Enchant           int  `json:"enchant"`
	MantainIngredient bool `json:"mantainIngredient"` //Сохранить инградиенты (к примеру рецепт у крафта нпц)
	FireAttr          int  `json:"fireAttr"`
	WaterAttr         int  `json:"waterAttr"`
	EarthAttr         int  `json:"earthAttr"`
	WindAttr          int  `json:"windAttr"`
	HolyAttr          int  `json:"holyAttr"`
	UnholyAttr        int  `json:"unholyAttr"`
}

type Production struct {
	Id         int `json:"id"`
	Count      int `json:"count"`
	Enchant    int `json:"enchant"`
	FireAttr   int `json:"fireAttr"`
	WaterAttr  int `json:"waterAttr"`
	EarthAttr  int `json:"earthAttr"`
	WindAttr   int `json:"windAttr"`
	HolyAttr   int `json:"holyAttr"`
	UnholyAttr int `json:"unholyAttr"`
}

var Multisells []MultiList

func LoadMultisell() {
	log.Println("Загрузка мультиселлов")
	msells := data.Find("./datapack/data/multisell", ".json")
	for _, msPath := range msells {
		var msell MultiList
		file, err := os.Open(msPath)
		if err != nil {
			panic("Failed to load config file " + err.Error())
		}
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&msell)
		if err != nil {
			panic("Failed to decode config file " + file.Name() + " " + err.Error())
		}
		msell.ID, err = strconv.Atoi(data.FileNameWithoutExtension(msPath))
		if err != nil {
			panic(err)
		}
		Multisells = append(Multisells, msell)
	}
	//log.Println(Multisells)
	//fmt.Printf("%+v\n", Multisells)
}

func Get(client interfaces.ReciverAndSender, id int) (MultiList, bool) {
	log.Println("Чтение GMShop", id)
	for _, multisell := range Multisells {
		if multisell.ID == id {
			return multisell, true
		}
	}
	return MultiList{}, false
}
