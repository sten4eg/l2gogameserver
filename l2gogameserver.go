package main

import (
	"l2gogameserver/config"
	"l2gogameserver/db"
	"l2gogameserver/loginserver"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//	gameserver.Load()
	//	gameserver.FindPath(-64072, 100856, -3584, -64072, 101048, -3584)

	setup()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	err := loginserver.HandlerInit()
	if err != nil {
		log.Fatal(err)
	}

	for {
	}
	// 	server.New().Start()

}

func setup() {
	config.LoadAllConfig()
	db.ConfigureDB()
	//idfactory.Load()
	//multisell.LoadMultisell()
	//teleport.LoadLocationListTeleport()
	//models.LoadStats()
	//models.LoadSkills()
	//items.LoadItems()
	//models.NewWorld()
	//data.Load()
	//models.LoadNpc()
}
