package models

import (
	"encoding/json"
	"log"
	"os"
)

type Spawn struct {
	X int32
	Y int32
	Z int32
}

type Location struct {
	ClassId int32
	Spawn   []Spawn
}

func Read() *[]Location {

	var config []Location
	file, err := os.Open("./data/stats/char/pcCreationPoint.json")
	if err != nil {
		log.Fatal("Failed to load config file")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Failed to decode config file")
	}
	return &config
}
