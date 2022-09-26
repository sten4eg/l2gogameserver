package targets

import (
	"errors"
	"strings"
)

type TargetType int

const (
	AREA TargetType = iota
	AREA_CORPSE_MOB
	AREA_FRIENDLY
	AREA_SUMMON
	AREA_UNDEAD
	AURA
	AURA_CORPSE_MOB
	AURA_FRIENDLY
	AURA_UNDEAD_ENEMY
	BEHIND_AREA
	BEHIND_AURA
	CLAN
	CLAN_MEMBER
	COMMAND_CHANNEL
	CORPSE
	CORPSE_CLAN
	CORPSE_MOB
	ENEMY
	ENEMY_ONLY
	ENEMY_SUMMON
	FLAGPOLE
	FRONT_AREA
	FRONT_AURA
	GROUND
	HOLY
	NONE
	ONE
	OWNER_PET
	PARTY
	PARTY_CLAN
	PARTY_MEMBER
	PARTY_NOTME
	PARTY_OTHER
	PC_BODY
	PET
	SELF
	SERVITOR
	SUMMON
	TARGET_PARTY
	UNDEAD
	UNLOCKABLE
)

func (t *TargetType) UnmarshalJSON(data []byte) error {
	sData := strings.Trim(string(data), "\"")
	switch sData {
	case "AREA":
		*t = AREA
	case "AREA_CORPSE_MOB":
		*t = AREA_CORPSE_MOB
	case "AREA_FRIENDLY":
		*t = AREA_FRIENDLY
	case "AREA_SUMMON":
		*t = AREA_SUMMON
	case "AREA_UNDEAD":
		*t = AREA_UNDEAD
	case "AURA":
		*t = AURA
	case "AURA_CORPSE_MOB":
		*t = AURA_CORPSE_MOB
	case "AURA_FRIENDLY":
		*t = AURA_FRIENDLY
	case "AURA_UNDEAD_ENEMY":
		*t = AURA_UNDEAD_ENEMY
	case "BEHIND_AREA":
		*t = BEHIND_AREA
	case "BEHIND_AURA":
		*t = BEHIND_AURA
	case "CLAN":
		*t = CLAN
	case "CLAN_MEMBER":
		*t = CLAN_MEMBER
	case "COMMAND_CHANNEL":
		*t = COMMAND_CHANNEL
	case "CORPSE":
		*t = CORPSE
	case "CORPSE_CLAN":
		*t = CORPSE_CLAN
	case "CORPSE_MOB":
		*t = CORPSE_MOB
	case "ENEMY":
		*t = ENEMY
	case "ENEMY_ONLY":
		*t = ENEMY_ONLY
	case "ENEMY_SUMMON":
		*t = ENEMY_SUMMON
	case "FLAGPOLE":
		*t = FLAGPOLE
	case "FRONT_AREA":
		*t = FRONT_AREA
	case "FRONT_AURA":
		*t = FRONT_AURA
	case "GROUND":
		*t = GROUND
	case "HOLY":
		*t = HOLY
	case "NONE":
		*t = NONE
	case "ONE":
		*t = ONE
	case "OWNER_PET":
		*t = OWNER_PET
	case "PARTY":
		*t = PARTY
	case "PARTY_CLAN":
		*t = PARTY_CLAN
	case "PARTY_MEMBER":
		*t = PARTY_MEMBER
	case "PARTY_NOTME":
		*t = PARTY_NOTME
	case "PARTY_OTHER":
		*t = PARTY_OTHER
	case "PC_BODY":
		*t = PC_BODY
	case "PET":
		*t = PET
	case "SELF":
		*t = SELF
	case "SERVITOR":
		*t = SERVITOR
	case "SUMMON":
		*t = SUMMON
	case "TARGET_PARTY":
		*t = TARGET_PARTY
	case "UNDEAD":
		*t = UNDEAD
	case "UNLOCKABLE":
		*t = UNLOCKABLE
	default:
		return errors.New("неправильный TargetType скила, TargetType = " + string(data))
	}
	return nil
}
