package models

import (
	"encoding/json"
	"log"
	"math/rand"
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

func GetCreationSpawn(classId int32) *Spawn {

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

	var spawn Spawn
	for _, v := range config {
		if v.ClassId == classId {
			spawn = v.Spawn[rand.Intn(len(v.Spawn))]
		}
	}
	return &spawn
}
