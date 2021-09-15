package main

import (
	"github.com/pkg/profile"
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
	server := gameserver.New()
	defer profile.Start(profile.MemProfileHeap, profile.ProfilePath(".")).Stop()

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
