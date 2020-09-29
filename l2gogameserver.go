package main

import (
	"l2gogameserver/gameserver"
	"l2gogameserver/gameserver/models"
)

func main() {
	models.Read()

	x := gameserver.New()
	x.Init()
	x.Start()
}
