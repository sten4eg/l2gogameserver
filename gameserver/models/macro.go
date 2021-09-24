package models

import (
	"context"
	"l2gogameserver/db"
	"log"
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
	Id       int32
	Name     string
	Desc     string
	Acronym  string
	Icon     byte
	Count    byte
	Commands []MacroCommand
}

//Добавление нового макроса
//Также эта функция служит и для изменения
//макроса(если пользователь добавил или поменял содержимое макроса)
func (c *Character) AddMacros(macro Macro) {
	//Если макрос найден, тогда делаем замену параметров его
	if c.CheckMacros(macro.Id) {
		for _, m := range c.Macros {
			if m.Id == macro.Id {
				removeMacros(macro.Id)
				c.saveMacros(macro)
				return
			}
		}
	} else {
		c.saveMacros(macro)
	}
}

//Удаление макроса
//Временное положение дел, необходимо будет НЕ удалять, а изменять, но пока и так сгодиться.
func removeMacros(id int32) {
	sqlMacros := `DELETE FROM "macros" WHERE "id" = $1`
	sqlCommands := `DELETE FROM "macros_commands" WHERE "command_id" = $1`
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
=======
	defer dbConn.Release()
>>>>>>> 67ebec2007b68bf2c47d3ecf1ae277e36cfd3071
	dbConn.Exec(context.Background(), sqlMacros, id)
	dbConn.Exec(context.Background(), sqlCommands, id)
}

//Сохранение макроса
func (c *Character) saveMacros(macro Macro) {
	//Макроса нет, добавляем в базу
	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()
	lastInsertId := 0
	sql := `INSERT INTO "macros" ("char_id", "icon", "name", "desc", "acronym")
			VALUES ($1, $2, $3, $4, $5) RETURNING id`
	_ = dbConn.QueryRow(context.Background(), sql, c.CharId, macro.Icon, macro.Name, macro.Desc, macro.Acronym).Scan(&lastInsertId)
	sql = `INSERT INTO "macros_commands" ("command_id", "index", "type", "skill_id", "shortcut_id", "name") VALUES ($1, $2, $3, $4, $5, $6)`
	for _, command := range macro.Commands {
		_, err = dbConn.Exec(context.Background(), sql, lastInsertId, command.Index, command.Type, command.SkillID, command.ShortcutID, command.Name)
		if err != nil {
			log.Println(err)
			return
		}
	}
	c.LoadCharactersMacros() //todo
}

//Проверка существования макроса
func (c *Character) CheckMacros(id int32) bool {
	for _, macro := range c.Macros {
		if macro.Id == id {
			return true
		}
	}
	return false
}

//Загрузка всех макросов игрока
func (c *Character) LoadCharactersMacros() {
	var Macroses []Macro
	var MacrosesCommands []MacroCommand
	c.Macros = nil

	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	sql := `SELECT * FROM "macros" WHERE char_id=$1 `
	rows, err := dbConn.Query(context.Background(), sql, c.CharId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for rows.Next() {
		m := Macro{}
		err = rows.Scan(nil, &m.Id, &m.Icon, &m.Name, &m.Desc, &m.Acronym)
		if err != nil {
			log.Println(err)
			return
		}
		m.Commands = MacrosesCommands
		Macroses = append(Macroses, m)
	}
	for index, macros := range Macroses {
		MacrosesCommands = nil
		sqlCommand := `SELECT * FROM "macros_commands" WHERE command_id=$1`
		rowsCommand, err := dbConn.Query(context.Background(), sqlCommand, macros.Id)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for rowsCommand.Next() {
			cCom := MacroCommand{}
			err = rows.Scan(&cCom.Id, &cCom.Index, &cCom.Type, &cCom.SkillID, &cCom.ShortcutID, &cCom.Name)
			if err != nil {
				log.Println(err)
				return
			}
			MacrosesCommands = append(MacrosesCommands, cCom)
		}
		macros.Count = uint8(index)
		macros.Commands = MacrosesCommands
		c.Macros = append(c.Macros, macros)
	}
}

//Кол-во имеющихся макросов у персонажа
func (c *Character) MacrosesCount() uint8 {
	return uint8(len(c.Macros))
}
