package interfaces

type SkillInterface interface {
	Identifier
	IsPassive() bool
	GetLevel() int
}
