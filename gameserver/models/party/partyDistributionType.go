package party

import "l2gogameserver/data/logger"

type PartyDistributionType struct {
	id          int32
	sysStringId int32
}

var partyDistributionTypes map[int32]PartyDistributionType

func (pdt *PartyDistributionType) GetId() int32 {
	return pdt.id
}

func (pdt *PartyDistributionType) GetSysStringId() int32 {
	return pdt.sysStringId
}

func LoadPartyDistributionTypes() {
	partyDistributionTypes = make(map[int32]PartyDistributionType)

	partyDistributionTypes[0] = FindersKeepers
	partyDistributionTypes[1] = Random
	partyDistributionTypes[2] = RandomIncludingSpoil
	partyDistributionTypes[3] = ByTurn
	partyDistributionTypes[4] = ByTurnIncludingSpoil

	logger.Info.Printf("Загружено %d PartyDistributionTypes", len(partyDistributionTypes))
}

func GetPartyDistributionTypeById(id int32) (*PartyDistributionType, bool) {
	partyDistributionType, ok := partyDistributionTypes[id]
	if ok {
		return &partyDistributionType, true
	}
	return nil, false
}

var FindersKeepers = PartyDistributionType{id: 0, sysStringId: 487}
var Random = PartyDistributionType{id: 1, sysStringId: 488}
var RandomIncludingSpoil = PartyDistributionType{id: 2, sysStringId: 798}
var ByTurn = PartyDistributionType{id: 3, sysStringId: 799}
var ByTurnIncludingSpoil = PartyDistributionType{id: 4, sysStringId: 800}
