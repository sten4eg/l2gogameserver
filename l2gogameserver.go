package main

import (
	"github.com/pkg/profile"
	"l2gogameserver/data"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
)

func main() {
	//defer profile.Start().Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	defer profile.Start().Stop()
	models.NewWorld()
	data.Load()
	x := gameserver.New()
	x.Init()
	x.Start()
}
