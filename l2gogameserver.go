package main

import (
	"l2gogameserver/data"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
	"log"
	"math"
)

func main() {
	log.Fatal(math.MinInt8)
	//	defer profile.Start().Stop()
	models.NewWorld()
	data.Load()
	x := gameserver.New()
	x.Init()
	x.Start()
}
