package interfaces

type SkillInterface interface {
	Identifier
	IsPassive() bool
	GetLevel() int
	GetName() string
	GetPower() int
	GetCastRange() int
	GetCoolTime() int
	GetHitTime() int
	GetOverHit() bool
	GetReuseDelay() int
	GetOperateType() int
	GetTargetType() int
	GetIsMagic() int
	GetMagicLvl() int
	GetMpConsume1() int
	GetMpConsume2() int
}

type SkillHolderInterface interface {
	GetSkill() SkillInterface
}
