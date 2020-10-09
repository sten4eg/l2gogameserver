package main

import (
	"l2gogameserver/data"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
)

func main() {
	//defer profile.Start().Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfileHeap()).Stop()
	models.NewWorld()
	data.Load()
	x := gameserver.New()
	x.Init()
	x.Start()
}
