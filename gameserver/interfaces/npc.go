package interfaces

type Npcer interface {
	UniquerId
	Identifier
	IsTargetable() bool
	GetCoordinates() (x, y, z int32)
	IsAttackable() int32
	GetHeading() int32
	GetCollisionRadius() float64
	GetCollisionHeight() float64
	GetSlotRhand() int32
	GetSlotLhand() int32
	GetMaxHp() int32
	// Новые методы для AI
	GetAI() NpcAI
	SetAI(NpcAI)
	GetSpawnLocation() (x, y, z int32)
	GetCurrentState() NpcState
	SetCurrentState(NpcState)
	IsMoving() bool
	SetMoving(bool)
	GetTarget() int32
	SetTarget(int32)
	GetAgroRange() int32
	GetCanMove() bool
	GetLevel() int32
	GetName() string
	// Методы для движения и позиционирования
	SetXYZ(x, y, z int32)
	CalculateDistanceTo(ox, oy, oz int32, includeZAxis, squared bool) float64

	// Новые методы для NpcInfo и отображения параметров
	GetRunSpd() int
	GetWalkSpd() int
	GetBaseAttackSpeed() int
	GetBaseMagicAttack() float64
	GetBaseDefend() float64
	GetBaseMagicDefend() float64
	GetBaseAttackRange() int
	GetBaseCritical() int
	GetBasePhysicalAttack() float64
	GetOrgHp() float64
	GetOrgMp() float64
	GetOrgHpRegen() float64
	GetOrgMpRegen() float64
	// Добавляйте другие методы по необходимости
}

// NpcAI интерфейс для AI поведения NPC
type NpcAI interface {
	Update(npc Npcer, world WorldRegioner)
	OnPlayerNearby(npc Npcer, player CharacterI, distance float64)
	OnAttacked(npc Npcer, attacker CharacterI)
	OnDeath(npc Npcer)
	GetBehaviorType() string
}

// NpcState состояние NPC
type NpcState int

const (
	NpcStateIdle NpcState = iota
	NpcStateMoving
	NpcStateChasing
	NpcStateAttacking
	NpcStateReturning
	NpcStateDead
)
