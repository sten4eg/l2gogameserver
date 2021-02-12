package main

import (
	"l2gogameserver/data"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
)

func main() {
	//defer profile.Start().Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfileHeap()).Stop()

	items.LoadItems()
	models.NewWorld()
	data.Load()

	server := gameserver.New()
	server.Init()
	server.Start()
}
