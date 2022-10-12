package models

import (
	"context"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
)

type MacroCommand struct {
	Id         int32
	Index      byte
	Type       byte
	SkillID    int32
	ShortcutID byte
	Name       string
}

type Macro struct {
	Id          int32
	Name        string
	Description string
	Acronym     string
	Icon        byte
	Count       byte
	Commands    []MacroCommand
}

// AddMacros Добавление нового макроса
// Также эта функция служит, и для изменения
// макроса (если пользователь добавил или поменял содержимое макроса)
func (c *Character) AddMacros(macro Macro) {
	//Если макрос найден, тогда делаем замену параметров его
	if c.CheckMacros(macro.Id) {
		RemoveMacros(macro.Id)
		c.saveMacros(macro)
		return
	} else {
		c.saveMacros(macro)
	}
}

// Удаление макроса
// todo Временное положение дел, необходимо будет НЕ удалять, а изменять, но пока и так сгодиться.
func RemoveMacros(id int32) {
	sqlMacros := `DELETE FROM "macros" WHERE "id" = $1`
	sqlCommands := `DELETE FROM "macros_commands" WHERE "command_id" = $1`
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	_, err = dbConn.Exec(context.Background(), sqlMacros, id)
	if err != nil {
		logger.Error.Panicln(err)
	}
	_, err = dbConn.Exec(context.Background(), sqlCommands, id)
	if err != nil {
		logger.Error.Panicln(err)
	}
}

// Сохранение макроса
func (c *Character) saveMacros(macro Macro) {
	//Макроса нет, добавляем в базу
	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	lastInsertId := 0
	sql := `INSERT INTO "macros" ("char_id", "icon", "name", "description", "acronym")
			VALUES ($1, $2, $3, $4, $5) RETURNING id`
	_ = dbConn.QueryRow(context.Background(), sql, c.ObjectId, macro.Icon, macro.Name, macro.Description, macro.Acronym).Scan(&lastInsertId)
	sql = `INSERT INTO "macros_commands" ("command_id", "index", "type", "skill_id", "shortcut_id", "name") VALUES ($1, $2, $3, $4, $5, $6)`
	for _, command := range macro.Commands {
		_, err = dbConn.Exec(context.Background(), sql, lastInsertId, command.Index, command.Type, command.SkillID, command.ShortcutID, command.Name)
		if err != nil {
			logger.Info.Println(err)
			return
		}
	}
	c.LoadCharactersMacros() //todo
}

// CheckMacros Проверка существования макроса
func (c *Character) CheckMacros(id int32) bool {
	for _, macro := range c.Macros {
		if macro.Id == id {
			return true
		}
	}
	return false
}

// LoadCharactersMacros Загрузка всех макросов игрока
func (c *Character) LoadCharactersMacros() {
	var Macroses []Macro
	var MacrosesCommands []MacroCommand
	c.Macros = nil

	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	sql := `SELECT id,icon,name,description,acronym FROM macros WHERE char_id=$1`
	rows, err := dbConn.Query(context.Background(), sql, c.ObjectId)
	if err != nil {
		logger.Info.Println(err.Error())
		return
	}
	for rows.Next() {
		m := Macro{}
		err = rows.Scan(&m.Id, &m.Icon, &m.Name, &m.Description, &m.Acronym)
		if err != nil {
			logger.Info.Println(err)
			return
		}
		m.Commands = MacrosesCommands
		Macroses = append(Macroses, m)
	}
	for index, macros := range Macroses {
		MacrosesCommands = nil
		sqlCommand := `SELECT * FROM macros_commands WHERE command_id=$1`
		rowsCommand, err := dbConn.Query(context.Background(), sqlCommand, macros.Id)
		if err != nil {
			logger.Info.Println(err.Error())
			return
		}
		for rowsCommand.Next() {
			cCom := MacroCommand{}
			err = rowsCommand.Scan(&cCom.Id, &cCom.Index, &cCom.Type, &cCom.SkillID, &cCom.ShortcutID, &cCom.Name)
			if err != nil {
				logger.Info.Println(err.Error())
				return
			}
			MacrosesCommands = append(MacrosesCommands, cCom)
		}
		macros.Count = uint8(index)
		macros.Commands = MacrosesCommands
		c.Macros = append(c.Macros, macros)
	}
}

// MacrosesCount Кол-во имеющихся макросов у персонажа
func (c *Character) MacrosesCount() uint8 {
	return uint8(len(c.Macros))
}
