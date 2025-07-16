package npc

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
)

// ExampleUsage демонстрирует использование AI системы
func ExampleUsage() {
	logger.Info.Println("=== Пример использования AI системы для NPC ===")

	// 1. Создаем менеджер AI (передаем nil для RegionProvider в примере)
	// В реальном использовании нужно передать правильный RegionProvider
	aiManager := NewAIManager(nil)

	// 2. Загружаем NPC (обычно это делается в main)
	npcs := LoadNpc()

	// 3. Инициализируем AI для всех NPC
	aiManager.InitializeAI(npcs)

	// 4. Запускаем цикл обновления AI
	aiManager.StartAIUpdate()

	logger.Info.Println("AI система инициализирована и запущена")

	// 5. Пример обработки событий
	// Когда игрок приближается к NPC:
	// aiManager.OnPlayerNearby(player)

	// Когда NPC атакован:
	// aiManager.OnNPCAttacked(npcObjId, attacker)

	// Когда NPC умирает:
	// aiManager.OnNPCDeath(npcObjId)
}

// DemonstrateNPCMovement демонстрирует движение NPC
func DemonstrateNPCMovement() {
	logger.Info.Println("=== Демонстрация движения NPC ===")

	// Создаем тестового NPC
	testNpc := &Npc{
		Name:      "Тестовый NPC",
		NpcId:     1001,
		ObjId:     10001,
		Type:      "monster", // Агрессивный тип
		AgroRange: 200,
		CanMove:   1,
		Spawn: Locations{
			Locx: 1000,
			Locy: 1000,
			Locz: 100,
		},
	}

	// Инициализируем координаты
	testNpc.SetXYZ(testNpc.Spawn.Locx, testNpc.Spawn.Locy, testNpc.Spawn.Locz)

	// Создаем агрессивный AI
	ai := NewAggressiveAI(200, 400, 600, nil) // Передаем nil, так как это пример
	testNpc.SetAI(ai)

	// Устанавливаем начальное состояние
	testNpc.SetCurrentState(interfaces.NpcStateIdle)
	testNpc.SetMoving(false)

	logger.Info.Printf("Создан тестовый NPC: %s (ID: %d)", testNpc.Name, testNpc.ObjId)
	x, y, z := testNpc.GetCoordinates()
	logger.Info.Printf("Начальные координаты: (%d, %d, %d)", x, y, z)

	// Симулируем приближение игрока
	logger.Info.Println("Симулируем приближение игрока...")

	// В реальной системе это будет вызвано автоматически
	// когда игрок приближается к NPC
	ai.OnPlayerNearby(testNpc, nil, 150.0)

	logger.Info.Printf("Состояние NPC после обнаружения игрока: %v",
		testNpc.GetCurrentState())
}

// DemonstrateAITypes демонстрирует различные типы AI
func DemonstrateAITypes() {
	logger.Info.Println("=== Демонстрация различных типов AI ===")

	// Пассивный NPC (для диалогов)
	passiveNpc := &Npc{
		Name: "Торговец",
		Type: "merchant",
	}
	passiveAI := NewPassiveAI(nil) // Передаем nil, так как это пример
	passiveNpc.SetAI(passiveAI)

	logger.Info.Printf("Пассивный NPC: %s, AI тип: %s",
		passiveNpc.Name, passiveAI.GetBehaviorType())

	// Агрессивный NPC (монстр)
	aggressiveNpc := &Npc{
		Name:      "Орк",
		Type:      "monster",
		AgroRange: 300,
	}
	aggressiveAI := NewAggressiveAI(300, 600, 900, nil) // Передаем nil, так как это пример
	aggressiveNpc.SetAI(aggressiveAI)

	logger.Info.Printf("Агрессивный NPC: %s, AI тип: %s",
		aggressiveNpc.Name, aggressiveAI.GetBehaviorType())
}
