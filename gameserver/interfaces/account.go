package interfaces

import "database/sql"

type AccountInterface interface {
	AddCharacter(CharacterI)
	GetChar(slot int) CharacterI
	Len() int
	GetLogin() string
	SetLogin(login string)
	MarkToDeleteChar(slot int32, db *sql.DB) int8
	GetCurrentChar() CharacterI
	SetSelectedChar(slot int)
}
