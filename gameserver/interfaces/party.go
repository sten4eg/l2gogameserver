package interfaces

type PartyInterface interface {
	GetMemberCount() int
	AddPartyMember(character CharacterI) bool
	GetLeaderObjectId() int32
	GetDistributionType() PartyDistributionTypeInterface
	SetMembers(members []CharacterI)
	GetMembers() []CharacterI
	GetLeader() CharacterI
	IsMemberInParty(character CharacterI) bool
	IsLeader(i CharacterI) bool
	IsDisbanding() bool
	SetDisbanding(bool)
	GetMemberIndex(CharacterI) int
	BroadcastParty([]byte)
	AddMember(member CharacterI)
	RemoveMember(member CharacterI)
}

type PartyDistributionTypeInterface interface {
	Identifier
	GetSysStringId() int32
}
