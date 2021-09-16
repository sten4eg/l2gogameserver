package main

import (
	"l2gogameserver/data"
	"l2gogameserver/db"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/utils"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//	gameserver.Load()
	//	gameserver.FindPath(-64072, 100856, -3584, -64072, 101048, -3584)

	setup()
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	server := gameserver.New()

	server.Start()
}

func setup() {
	db.ConfigureDB()
	models.LoadStats()
	models.LoadSkills()
	models.LoadItems()
	models.NewWorld()
	data.Load()

	utils.SetupServerPackets()
}
