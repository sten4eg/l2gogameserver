package main

import (
	"fmt"
	"l2gogameserver/config"
	"l2gogameserver/data"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/gameserver/models/teleport"
	"l2gogameserver/loginserver"
	"l2gogameserver/server"
	"log"
	"math"
)

func main() {
	fmt.Println((11 - 20) << 15)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//	gameserver.Load()
	//	gameserver.FindPath(-64072, 100856, -3584, -64072, 101048, -3584)

	setup()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	err := loginserver.HandlerInit()
	if err != nil {
		log.Fatal(err)
	}

	const MAP_MIN_X = (config.GeoFirstX - 20) << 15
	const MAP_MAX_X = ((config.GeoLastX - 20 + 1) << 15) - 1
	const MAP_MIN_Y = (config.GeoFirstY - 18 + 1) << 15
	const MAP_MAX_Y = ((config.GeoLastY - 18 + 1) << 15) - 1
	const MAP_MIN_Z = -16384
	const MAP_MAX_Z = 16384

	const WORLD_SIZE_X = config.GeoLastX - config.GeoFirstX + 1
	const WORLD_SIZE_Y = config.GeoLastY - config.GeoFirstY + 1

	const SHIFT_BY = config.SHIFT_BY
	const SHIFT_BY_Z = config.SHIFT_BY_Z

	var OFFSET_X = math.Abs(MAP_MIN_X >> SHIFT_BY)
	var OFFSET_Y = math.Abs(MAP_MIN_Y >> SHIFT_BY)
	var OFFSET_Z = math.Abs(MAP_MIN_Z >> SHIFT_BY_Z)

	var REGIONS_X = int32((MAP_MAX_X >> SHIFT_BY) + OFFSET_X)
	var REGIONS_Y = int32((MAP_MAX_Y >> SHIFT_BY) + OFFSET_Y)
	var REGIONS_Z = int32((MAP_MAX_Z >> SHIFT_BY_Z) + OFFSET_Z)

	_, _, _ = REGIONS_Z, REGIONS_Y, REGIONS_X
	server.New().Start()

}

func setup() {
	config.LoadAllConfig()
	db.ConfigureDB()
	idfactory.Load()
	multisell.LoadMultisell()
	teleport.LoadLocationListTeleport()
	models.LoadStats()
	models.LoadSkills()
	items.LoadItems()
	models.NewWorld()
	data.Load()
	models.LoadNpc()
}
