package npc

import (
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/types"
	"l2gogameserver/packets"
	"math"
	"math/rand"
	"time"
)

// --- КОНСТАНТЫ ---
const (
	PATROL_RADIUS           = 300  // Радиус патрулирования NPC от точки респауна
	BROADCAST_RADIUS        = 2000 // Радиус рассылки пакета движения игрокам
	OPCODE_MOVE_TO_LOCATION = 0x2f // Опкод пакета MoveToLocation
	NPC_SAY_RADIUS          = 1000 // Радиус рассылки NpcSay
	OPCODE_NPC_SAY          = 0x2c
	NPC_SAY_PATROL          = "Я иду патрулировать!"
	NPC_SAY_SEE_PLAYER      = "Я тебя вижу, человек!"
)

// BaseAI базовая реализация AI для NPC
type BaseAI struct {
	behaviorType           string
	lastUpdate             time.Time
	lastMovement           time.Time
	movementCooldown       time.Duration
	patrolRadius           int32
	agroRange              int32
	chaseRange             int32
	returnRange            int32
	spawnX, spawnY, spawnZ int32
	regionProvider         interfaces.RegionProvider
}

// NewBaseAI создает новый базовый AI
func NewBaseAI(behaviorType string, agroRange, chaseRange, returnRange int32, regionProvider interfaces.RegionProvider) *BaseAI {
	return &BaseAI{
		behaviorType:     behaviorType,
		movementCooldown: time.Second * 5, // Уменьшаем с 10 до 5 секунд
		patrolRadius:     100,
		agroRange:        agroRange,
		chaseRange:       chaseRange,
		returnRange:      returnRange,
		lastUpdate:       time.Now(),
		lastMovement:     time.Now(),
		regionProvider:   regionProvider,
	}
}

func (ai *BaseAI) GetBehaviorType() string {
	return ai.behaviorType
}

func (ai *BaseAI) Update(npc interfaces.Npcer, world interfaces.WorldRegioner) {
	now := time.Now()

	// Обновляем не чаще чем раз в секунду
	if now.Sub(ai.lastUpdate) < time.Second {
		return
	}
	ai.lastUpdate = now

	currentState := npc.GetCurrentState()

	// Логируем состояние NPC для отладки
	// logger.Info.Printf("NPC %s (ID: %d) состояние: %v, движется: %v",
	// npc.GetName(), npc.GetObjectId(), currentState, npc.IsMoving())

	switch currentState {
	case interfaces.NpcStateIdle:
		ai.handleIdleState(npc, world)
	case interfaces.NpcStateMoving:
		ai.handleMovingState(npc, world)
	case interfaces.NpcStateChasing:
		ai.handleChasingState(npc, world)
	case interfaces.NpcStateReturning:
		ai.handleReturningState(npc, world)
	}
}

func (ai *BaseAI) OnPlayerNearby(npc interfaces.Npcer, player interfaces.CharacterI, distance float64) {
	if distance <= float64(ai.agroRange) && npc.GetCurrentState() == interfaces.NpcStateIdle {
		// NPC замечает игрока и начинает преследование
		npc.SetCurrentState(interfaces.NpcStateChasing)
		npc.SetTarget(player.GetObjectId())
		// Сообщаем игрокам, что NPC заметил игрока
		ai.broadcastNpcSay(npc, NPC_SAY_SEE_PLAYER)
	}
}

func (ai *BaseAI) OnAttacked(npc interfaces.Npcer, attacker interfaces.CharacterI) {
	// NPC атакован - переходит в режим преследования
	npc.SetCurrentState(interfaces.NpcStateChasing)
	npc.SetTarget(attacker.GetObjectId())
	logger.Info.Printf("NPC %s (ID: %d) атакован игроком %s",
		npc.GetName(), npc.GetObjectId(), attacker.GetName())
}

func (ai *BaseAI) OnDeath(npc interfaces.Npcer) {
	npc.SetCurrentState(interfaces.NpcStateDead)
	logger.Info.Printf("NPC %s (ID: %d) погиб", npc.GetName(), npc.GetObjectId())
}

// handleIdleState обработка состояния покоя
func (ai *BaseAI) handleIdleState(npc interfaces.Npcer, world interfaces.WorldRegioner) {
	now := time.Now()

	// Проверяем, не пора ли начать патрулирование
	if now.Sub(ai.lastMovement) > ai.movementCooldown {
		// Случайное движение в радиусе патрулирования
		if rand.Float64() < 0.7 { // Увеличиваем шанс до 70% для более активного движения
			ai.startRandomMovement(npc)
		}
	}

	// Проверяем игроков поблизости
	ai.checkNearbyPlayers(npc, world)
}

// handleMovingState обработка состояния движения
func (ai *BaseAI) handleMovingState(npc interfaces.Npcer, world interfaces.WorldRegioner) {
	// Проверяем, достигли ли мы цели
	if !npc.IsMoving() {
		// Движение завершено, возвращаемся в состояние покоя
		npc.SetCurrentState(interfaces.NpcStateIdle)
		ai.lastMovement = time.Now()
		return
	}

	// Проверяем игроков поблизости во время движения
	ai.checkNearbyPlayers(npc, world)

	// Случайно завершаем движение (симуляция достижения цели)
	if rand.Float64() < 0.1 { // 10% шанс завершить движение
		npc.SetMoving(false)
		npc.SetCurrentState(interfaces.NpcStateIdle)
		ai.lastMovement = time.Now()
	}
}

// handleChasingState обработка состояния преследования
func (ai *BaseAI) handleChasingState(npc interfaces.Npcer, world interfaces.WorldRegioner) {
	targetId := npc.GetTarget()
	if targetId == 0 {
		// Цель потеряна, возвращаемся к спавну
		npc.SetCurrentState(interfaces.NpcStateReturning)
		return
	}

	// Ищем цель в соседних регионах
	target := ai.findTarget(npc, world, targetId)
	if target == nil {
		// Цель не найдена, возвращаемся к спавну
		npc.SetCurrentState(interfaces.NpcStateReturning)
		return
	}

	// Проверяем расстояние до цели
	distance := npc.CalculateDistanceTo(target.GetX(), target.GetY(), target.GetZ(), false, false)

	if distance > float64(ai.chaseRange) {
		// Цель слишком далеко, возвращаемся к спавну
		npc.SetCurrentState(interfaces.NpcStateReturning)
		return
	}

	// Двигаемся к цели
	ai.moveToTarget(npc, target)
}

// handleReturningState обработка состояния возвращения к спавну
func (ai *BaseAI) handleReturningState(npc interfaces.Npcer, world interfaces.WorldRegioner) {
	spawnX, spawnY, spawnZ := npc.GetSpawnLocation()

	distanceToSpawn := npc.CalculateDistanceTo(spawnX, spawnY, spawnZ, false, false)

	if distanceToSpawn < 50 {
		// Достигли спавна, возвращаемся в состояние покоя
		npc.SetCurrentState(interfaces.NpcStateIdle)
		npc.SetTarget(0)
		npc.SetMoving(false)
		return
	}

	// Двигаемся к спавну
	ai.moveToLocation(npc, spawnX, spawnY, spawnZ)
}

// startRandomMovement начинает случайное движение NPC в пределах PATROL_RADIUS от точки респауна
func (ai *BaseAI) startRandomMovement(npc interfaces.Npcer) {
	spawnX, spawnY, spawnZ := npc.GetSpawnLocation()

	// Генерируем случайную точку в радиусе PATROL_RADIUS
	angle := rand.Float64() * 2 * math.Pi
	radius := rand.Float64() * float64(PATROL_RADIUS)
	targetX := spawnX + int32(radius*math.Cos(angle))
	targetY := spawnY + int32(radius*math.Sin(angle))
	targetZ := spawnZ // Пока оставляем ту же высоту

	// Сообщаем игрокам, что NPC начинает патрулировать
	ai.broadcastNpcSay(npc, NPC_SAY_PATROL)

	ai.moveToLocation(npc, targetX, targetY, targetZ)
	npc.SetCurrentState(interfaces.NpcStateMoving)
	ai.lastMovement = time.Now()
}

// moveToLocation двигает NPC к указанной точке
func (ai *BaseAI) moveToLocation(npc interfaces.Npcer, targetX, targetY, targetZ int32) {
	currentX, currentY, currentZ := npc.GetCoordinates()

	// Создаем пакет движения
	location := &types.BackwardToLocation{
		TargetX: targetX,
		TargetY: targetY,
		TargetZ: targetZ,
		OriginX: currentX,
		OriginY: currentY,
		OriginZ: currentZ,
	}

	// Отправляем пакет движения всем игрокам в регионе
	ai.broadcastMovement(npc, location)

	// Обновляем координаты NPC
	npc.SetXYZ(targetX, targetY, targetZ)
	npc.SetMoving(true)
}

// moveToTarget двигает NPC к цели
func (ai *BaseAI) moveToTarget(npc interfaces.Npcer, target interfaces.CharacterI) {
	targetX, targetY, targetZ := target.GetXYZ()
	currentX, currentY, _ := npc.GetCoordinates()

	// Вычисляем направление к цели
	dx := float64(targetX - currentX)
	dy := float64(targetY - currentY)
	distance := math.Sqrt(dx*dx + dy*dy)

	if distance < 50 {
		// Достаточно близко для атаки
		npc.SetCurrentState(interfaces.NpcStateAttacking)
		return
	}

	// Нормализуем вектор направления
	if distance > 0 {
		dx /= distance
		dy /= distance
	}

	// Двигаемся на 50 единиц в направлении цели
	moveDistance := float64(50)
	newX := currentX + int32(dx*moveDistance)
	newY := currentY + int32(dy*moveDistance)
	newZ := targetZ // Пока оставляем ту же высоту что у цели

	ai.moveToLocation(npc, newX, newY, newZ)
}

// checkNearbyPlayers проверяет игроков поблизости
func (ai *BaseAI) checkNearbyPlayers(npc interfaces.Npcer, world interfaces.WorldRegioner) {
	// Проверяем всех игроков в соседних регионах
	for _, region := range world.GetNeighbors() {
		for _, player := range region.GetCharsInRegion() {
			if player == nil {
				continue
			}

			distance := npc.CalculateDistanceTo(player.GetX(), player.GetY(), player.GetZ(), false, false)
			ai.OnPlayerNearby(npc, player, distance)
		}
	}
}

// findTarget ищет цель по ID
func (ai *BaseAI) findTarget(npc interfaces.Npcer, world interfaces.WorldRegioner, targetId int32) interfaces.CharacterI {
	for _, region := range world.GetNeighbors() {
		for _, player := range region.GetCharsInRegion() {
			if player != nil && player.GetObjectId() == targetId {
				return player
			}
		}
	}
	return nil
}

// broadcastMovement отправляет пакет движения только игрокам в радиусе BROADCAST_RADIUS от NPC
func (ai *BaseAI) broadcastMovement(npc interfaces.Npcer, location *types.BackwardToLocation) {
	// Получаем регион NPC
	x, y, z := npc.GetCoordinates()
	region := ai.getRegionProvider().GetRegion(x, y, z)
	if region == nil {
		return
	}

	// Собираем всех игроков в радиусе BROADCAST_RADIUS
	players := make([]interfaces.CharacterI, 0)
	for _, reg := range region.GetNeighbors() {
		for _, player := range reg.GetCharsInRegion() {
			if player == nil {
				continue
			}
			dx := float64(player.GetX() - x)
			dy := float64(player.GetY() - y)
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist <= BROADCAST_RADIUS {
				players = append(players, player)
			}
		}
	}

	// Формируем пакет движения
	movementPacket := makeMoveToLocationPacket(location, npc)

	// Отправляем пакет только игрокам в радиусе
	for _, player := range players {
		player.EncryptAndSend(movementPacket)
	}
}

// makeMoveToLocationPacket создает пакет движения NPC (0x2F)
func makeMoveToLocationPacket(location *types.BackwardToLocation, npc interfaces.Npcer) []byte {
	buffer := getPacketBuffer()
	buffer.WriteSingleByte(OPCODE_MOVE_TO_LOCATION)
	buffer.WriteD(npc.GetObjectId())
	buffer.WriteD(location.TargetX)
	buffer.WriteD(location.TargetY)
	buffer.WriteD(location.TargetZ)
	buffer.WriteD(location.OriginX)
	buffer.WriteD(location.OriginY)
	buffer.WriteD(location.OriginZ)
	return buffer.Bytes()
}

// getPacketBuffer возвращает новый буфер пакета
func getPacketBuffer() *packets.Buffer {
	return packets.Get()
}

// getRegionProvider возвращает провайдер регионов
func (ai *BaseAI) getRegionProvider() interfaces.RegionProvider {
	return ai.regionProvider
}

// broadcastNpcSay отправляет пакет NpcSay всем игрокам в радиусе NPC_SAY_RADIUS
func (ai *BaseAI) broadcastNpcSay(npc interfaces.Npcer, text string) {
	x, y, z := npc.GetCoordinates()
	region := ai.getRegionProvider().GetRegion(x, y, z)
	if region == nil {
		return
	}
	players := make([]interfaces.CharacterI, 0)
	for _, reg := range region.GetNeighbors() {
		for _, player := range reg.GetCharsInRegion() {
			if player == nil {
				continue
			}
			dx := float64(player.GetX() - x)
			dy := float64(player.GetY() - y)
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist <= NPC_SAY_RADIUS {
				players = append(players, player)
			}
		}
	}
	// Формируем пакет NpcSay
	packet := makeNpcSayPacket(npc, text)
	// packet := serverpackets.NpcSay(npc.GetObjectId(), 0, npc.GetId(), text)
	for _, player := range players {
		player.EncryptAndSend(packet)
	}
}

// makeNpcSayPacket создает пакет NpcSay (0x2C)
func makeNpcSayPacket(npc interfaces.Npcer, npcString string) []byte {
	buffer := packets.Get()

	buffer.WriteSingleByte(0x30)
	buffer.WriteD(npc.GetObjectId())
	buffer.WriteD(1)
	buffer.WriteD(npc.GetId())
	buffer.WriteD(-1) // -1 тогда свой текст
	buffer.WriteS(npcString)
	return buffer.Bytes()
}

// PassiveAI пассивный AI для NPC с диалогами
type PassiveAI struct {
	*BaseAI
}

func NewPassiveAI(regionProvider interfaces.RegionProvider) *PassiveAI {
	return &PassiveAI{
		BaseAI: NewBaseAI("passive", 0, 0, 0, regionProvider),
	}
}

func (ai *PassiveAI) OnPlayerNearby(npc interfaces.Npcer, player interfaces.CharacterI, distance float64) {
	// Пассивные NPC не реагируют на приближение игроков
}

func (ai *PassiveAI) OnAttacked(npc interfaces.Npcer, attacker interfaces.CharacterI) {
	// Пассивные NPC не атакуют в ответ
}

// AggressiveAI агрессивный AI для монстров
type AggressiveAI struct {
	*BaseAI
}

func NewAggressiveAI(agroRange, chaseRange, returnRange int32, regionProvider interfaces.RegionProvider) *AggressiveAI {
	return &AggressiveAI{
		BaseAI: NewBaseAI("aggressive", agroRange, chaseRange, returnRange, regionProvider),
	}
}

func (ai *AggressiveAI) OnAttacked(npc interfaces.Npcer, attacker interfaces.CharacterI) {
	// Агрессивные NPC сразу переходят в режим преследования при атаке
	npc.SetCurrentState(interfaces.NpcStateChasing)
	npc.SetTarget(attacker.GetObjectId())
	logger.Info.Printf("Агрессивный NPC %s (ID: %d) атакован игроком %s",
		npc.GetName(), npc.GetObjectId(), attacker.GetName())
}
