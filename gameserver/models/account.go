package models

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
)

type Account struct {
	Char             []interfaces.CharacterI
	CharSlotSelected int32
	Login            string
}

const DeleteCharacterShortcuts = `DELETE FROM character_shortcuts WHERE char_id = $1`
const DeleteCharacterSkills = `DELETE FROM character_skills WHERE char_id = $1`
const DeleteCharacter = `DELETE FROM characters WHERE object_id = $1`
const DeleteCharacterItems = `DELETE FROM items WHERE owner_id = $1`

func (a *Account) AddCharacter(character interfaces.CharacterI) {
	a.Char = append(a.Char, character)
}
func (a *Account) Len() int {
	return len(a.Char)
}

func (a *Account) SetLogin(login string) {
	a.Login = login
}

func (a *Account) GetLogin() string {
	return a.Login
}

func (a *Account) GetChar(id int) interfaces.CharacterI {
	return a.Char[id]
}

func (a *Account) MarkToDeleteChar(slot int32, db *sql.DB) int8 {
	if slot < 0 || int(slot) >= len(a.Char) {
		return -1
	}

	objId := a.Char[slot].GetObjectId()

	// TODO чекнуть в бд клан персонажа
	var answer int8

	if answer == 0 { // clan == nil
		a.deleteCharByObjId(objId, db)
	}

	return answer
}

func (a *Account) GetCurrentChar() interfaces.CharacterI {
	return a.Char[a.CharSlotSelected]
}
func (a *Account) SetSelectedChar(slot int) {
	a.CharSlotSelected = int32(slot)
}

func (a *Account) deleteCharByObjId(objId int32, db *sql.DB) {
	if objId < 0 {
		return
	}

	_, err := db.Exec(DeleteCharacterShortcuts, objId)
	if err != nil {
		logger.Error.Panicln(err)
	}

	_, err = db.Exec(DeleteCharacterItems, objId)
	if err != nil {
		logger.Error.Panicln(err)
	}

	_, err = db.Exec(DeleteCharacter, objId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}
