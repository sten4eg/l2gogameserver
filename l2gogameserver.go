package main

import (
	"l2gogameserver/data"
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
)

func main() {

	models.NewWorld()
	data.Load()
	x := gameserver.New()
	x.Init()
	x.Start()
}
