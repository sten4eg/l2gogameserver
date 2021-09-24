package main

import (
	"l2gogameserver/data"
	"l2gogameserver/db"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/models"
	"log"
)

func main() {
	//lolkek
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//	gameserver.Load()
	//	gameserver.FindPath(-64072, 100856, -3584, -64072, 101048, -3584)

	setup()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	gameserver.New().Start()

}

func setup() {
	db.ConfigureDB()
	idfactory.Load()
	models.LoadStats()
	models.LoadSkills()
	models.LoadItems()
	models.NewWorld()
	data.Load()
	models.LoadNpc()

}
