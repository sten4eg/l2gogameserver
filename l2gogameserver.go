package main

import (
	"l2gogameserver/data"
	"l2gogameserver/db"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
)

func main() {

	//	gameserver.Load()
	//	gameserver.FindPath(-64072, 100856, -3584, -64072, 101048, -3584)

	setup()

	db.ConfigureDB()
	server := gameserver.New()
	server.Init()
	server.Start()
}

func setup() {
	models.LoadStats()
	models.LoadSkills()
	items.LoadItems()
	models.NewWorld()
	data.Load()

}
