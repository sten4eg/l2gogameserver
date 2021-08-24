package main

import (
	"l2gogameserver/data"
	"l2gogameserver/db"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
)

//todo 	var packet = packets.NewReader(data)
// в каждом пакете от клиента , возможно лучше сделать пул ?
func main() {
	//defer profile.Start().Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfileHeap()).Stop()
	//

	gameserver.Load()
	//strings.Join()
	gameserver.FindPath(-64072, 100856, -3584, -64072, 101048, -3584)

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
