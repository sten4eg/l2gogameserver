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
}
