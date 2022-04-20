package teleport

import (
	"encoding/json"
	"log"
	"os"
)

type Location struct {
	Id      int  `json:"id"`
	X       int  `json:"x"`
	Y       int  `json:"y"`
	Z       int  `json:"z"`
	Price   int  `json:"price"`
	IsNoble bool `json:"isNoble"`
}

var Locations []Location

//Загрузка позиций локаций для телепортации
func LoadLocationListTeleport() {
	log.Println("Загрузка позиций к телепортации")
	file, err := os.Open("./datapack/data/teleport/locationToTeleport.json")
	if err != nil {
		panic("Failed to load config file " + err.Error())
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Locations)
	if err != nil {
		panic("Failed to decode config file " + file.Name() + " " + err.Error())
	}
}

// Вернуть данные о локации к телепорту
func GetTeleportID(id int) (Location, bool) {
	for _, location := range Locations {
		if location.Id == id {
			return location, true
		}
	}
	return Location{}, false
}
