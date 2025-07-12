package main

import (
	"fmt"
	"l2gogameserver/config"
	"l2gogameserver/data"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/idfactory"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/gameserver/models/multisell"
	"l2gogameserver/gameserver/models/party"
	"l2gogameserver/gameserver/models/teleport"
	"l2gogameserver/gameserver/npc"
	"l2gogameserver/gameserver/world"
	"l2gogameserver/loginserver"
	"l2gogameserver/server"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func stp(file *os.File) {
	time.Sleep(5 * time.Minute)
	pprof.StopCPUProfile()
	_ = file.Close()
	fmt.Println("cpu profiling STOPPED")
	fmt.Println("cpu profiling STOPPED")
	fmt.Println("cpu profiling STOPPED")
	fmt.Println("cpu profiling STOPPED")
	fmt.Println("cpu profiling STOPPED")
}
func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	world.GetRegionFromXY(12446, 16683)

	f, err := os.OpenFile("test", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	go stp(f)

	// Инициализируем конфигурацию
	configManager := config.GetConfigManager()
	if err := configManager.Load("./config/config.yaml"); err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	cfg := configManager.Get()

	if cfg.GameServer.Debug.IsShowPackets() {
		logger.Info.Println("Показ пакетов включен")
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
	data.Load()
	npcs := npc.LoadNpc()
	worldS := world.NewWorld()
	worldS.AddNpc(npcs)

	party.LoadPartyDistributionTypes()

	err = loginserver.HandlerInit(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	server.New(dbConn).Start()

}
