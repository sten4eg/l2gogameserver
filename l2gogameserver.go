package main

import (
	"l2gogameserver/config"
	"l2gogameserver/data"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/gameserver/models/party"
	"l2gogameserver/gameserver/models/teleport"
	"l2gogameserver/loginserver"
	"l2gogameserver/server"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.ConfigureDB(cfg.GameServer.Database)
	if err != nil {
		log.Fatal(err)
	}

	idfactory.Load(dbConn)
	multisell.LoadMultisell()
	teleport.LoadLocationListTeleport()
	models.LoadStats()
	models.LoadSkills()
	items.LoadItems()
	models.NewWorld()
	data.Load()
	models.LoadNpc()

	party.LoadPartyDistributionTypes()

	err = loginserver.HandlerInit(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	server.New().Start(dbConn)

}
