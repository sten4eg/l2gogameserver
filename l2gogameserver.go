package main

import (
	"l2gogameserver/config"
	"l2gogameserver/data"
	"l2gogameserver/db"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/gameserver/models/teleport"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//	gameserver.Load()
	//	gameserver.FindPath(-64072, 100856, -3584, -64072, 101048, -3584)

	setup()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	gameserver.New().Start()

}

func setup() {
	config.LoadAllConfig()
	db.ConfigureDB()
	idfactory.Load()
	teleport.LoadLocationListTeleport()
	models.LoadStats()
	models.LoadSkills()
	items.LoadItems()
	models.NewWorld()
	data.Load()
	models.LoadNpc()
	multisell.LoadMultisell()
}
