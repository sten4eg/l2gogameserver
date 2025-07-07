package interfaces

type MacrosInterface interface {
	Identifier
	Namer
	GetDescription() string
	GetAcronym() string
	GetIcon() byte
	GetCount() byte
	GetCommands() []MacrosCommandInterface
}

type MacrosCommandInterface interface {
	Identifier
	GetIndex() byte
	GetType() byte
	GetSkillId() int32
	GetShortcutId() byte
	GetName() string
}
