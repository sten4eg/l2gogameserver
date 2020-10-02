package main

import (
	"l2gogameserver/data"
	"l2gogameserver/gameserver"
)

func main() {
	data.Load()
	x := gameserver.New()
	x.Init()
	x.Start()
}
