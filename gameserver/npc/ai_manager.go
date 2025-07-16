package npc

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"time"
)

// AIManager управляет AI для всех NPC
type AIManager struct {
	npcs           map[int32]interfaces.Npcer
	regionProvider interfaces.RegionProvider
}

// Глобальный экземпляр AIManager для доступа из других пакетов
var AIManagerInstance *AIManager

// NewAIManager создает новый менеджер AI
func NewAIManager(regionProvider interfaces.RegionProvider) *AIManager {
	return &AIManager{
		npcs:           make(map[int32]interfaces.Npcer),
		regionProvider: regionProvider,
	}
}

// InitializeAI инициализирует AI для всех NPC
func (am *AIManager) InitializeAI(npcs map[int32]map[int32]Npc) {
	logger.Info.Println("Инициализация AI для NPC...")

	for npcId, npcInstances := range npcs {
		for objId, npc := range npcInstances {
			// Инициализируем координаты NPC
			npc.SetXYZ(npc.Spawn.Locx, npc.Spawn.Locy, npc.Spawn.Locz)

			// Определяем тип AI на основе типа NPC
			ai := am.createAIForNPC(&npc)
			npc.SetAI(ai)

			// Устанавливаем начальное состояние
			npc.SetCurrentState(interfaces.NpcStateIdle)
			npc.SetMoving(false)
			npc.SetTarget(0)

			// Добавляем в менеджер
			am.npcs[objId] = &npc

			//logger.Info.Printf("Инициализирован AI для NPC %s (ID: %d, ObjID: %d)",
			//	npc.Name, npcId, objId)
			_ = npcId
		}
	}

	logger.Info.Printf("Инициализировано AI для %d NPC", len(am.npcs))
}

// createAIForNPC создает подходящий AI для NPC
func (am *AIManager) createAIForNPC(npc *Npc) interfaces.NpcAI {
	// Определяем тип NPC (0 - диалоговый, 1 - монстр)
	npcType := GetDialogNPC(npc.Type)

	switch npcType {
	case 0: // NPC с диалогами (пассивные)
		return NewPassiveAI(am.regionProvider)
	case 1: // Монстры (агрессивные)
		agroRange := int32(npc.AgroRange)
		if agroRange == 0 {
			// Если agroRange == 0, NPC полностью пассивный, не реагирует на игроков
			return NewPassiveAI(am.regionProvider)
		}
		chaseRange := agroRange * 2
		returnRange := agroRange * 3
		return NewAggressiveAI(agroRange, chaseRange, returnRange, am.regionProvider)
	default:
		// По умолчанию пассивный AI
		return NewPassiveAI(am.regionProvider)
	}
}

// UpdateAI обновляет AI для всех NPC
func (am *AIManager) UpdateAI() {
	for _, npc := range am.npcs {
		if npc == nil {
			continue
		}

		// Получаем регион NPC
		x, y, z := npc.GetCoordinates()
		region := am.regionProvider.GetRegion(x, y, z)

		if region == nil {
			continue
		}

		// Обновляем AI
		ai := npc.GetAI()
		if ai != nil {
			ai.Update(npc, region)
		}
	}
}

// StartAIUpdate запускает цикл обновления AI
func (am *AIManager) StartAIUpdate() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			am.UpdateAI()
		}
	}()

	logger.Info.Println("Запущен цикл обновления AI для NPC")
}

// GetNPC возвращает NPC по ID объекта
func (am *AIManager) GetNPC(objId int32) interfaces.Npcer {
	return am.npcs[objId]
}

// OnPlayerNearby обрабатывает приближение игрока к NPC
func (am *AIManager) OnPlayerNearby(player interfaces.CharacterI) {
	playerX, playerY, playerZ := player.GetXYZ()
	playerRegion := am.regionProvider.GetRegion(playerX, playerY, playerZ)

	if playerRegion == nil {
		return
	}

	// Проверяем всех NPC в соседних регионах
	for _, region := range playerRegion.GetNeighbors() {
		for _, npc := range region.GetNpcInRegion() {
			if npc == nil {
				continue
			}

			// Вычисляем расстояние до игрока
			distance := npc.CalculateDistanceTo(playerX, playerY, playerZ, false, false)

			// Уведомляем AI о приближении игрока
			ai := npc.GetAI()
			if ai != nil {
				ai.OnPlayerNearby(npc, player, distance)
			}
		}
	}
}

// OnNPCAttacked обрабатывает атаку на NPC
func (am *AIManager) OnNPCAttacked(npcObjId int32, attacker interfaces.CharacterI) {
	npc := am.GetNPC(npcObjId)
	if npc == nil {
		return
	}

	ai := npc.GetAI()
	if ai != nil {
		ai.OnAttacked(npc, attacker)
	}
}

// OnNPCDeath обрабатывает смерть NPC
func (am *AIManager) OnNPCDeath(npcObjId int32) {
	npc := am.GetNPC(npcObjId)
	if npc == nil {
		return
	}

	ai := npc.GetAI()
	if ai != nil {
		ai.OnDeath(npc)
	}
}
