package models

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
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
func (c *Character) AddMacros(macro interfaces.MacrosInterface) {
	//Если макрос найден, тогда делаем замену параметров его
	if c.CheckMacros(macro.GetId()) {
		RemoveMacros(macro.GetId(), c.GetObjectId(), c.Conn.db)
		c.saveMacros(macro, macro.GetId())
		return
	} else {
		c.saveMacros(macro, c.GetMacroId())
	}
}

// Удаление макроса
// todo Временное положение дел, необходимо будет НЕ удалять, а изменять, но пока и так сгодиться.
func RemoveMacros(id, charId int32, db *sql.DB) {
	sqlMacros := `DELETE FROM "macros" WHERE "id" = $1 and char_id = $2`
	sqlCommands := `DELETE FROM "macros_commands" WHERE "command_id" = $1 and char_id = $2`

	_, err := db.Exec(sqlMacros, id, charId)
	if err != nil {
		logger.Error.Panicln(err)
	}
	_, err = db.Exec(sqlCommands, id, charId)
	if err != nil {
		logger.Error.Panicln(err)
	}
}

// Сохранение макроса
func (c *Character) saveMacros(macro interfaces.MacrosInterface, id int32) {
	//Макроса нет, добавляем в базу

	sql := `INSERT INTO "macros" ("char_id", "id", "icon", "name", "description", "acronym")
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	_, err := c.Conn.db.Exec(sql, c.ObjectId, id, macro.GetIcon(), macro.GetName(), macro.GetDescription(), macro.GetAcronym())

	if err != nil {
		logger.Info.Println(err)
		return
	}

	sql = `INSERT INTO "macros_commands" ("command_id", "char_id", "index", "type", "skill_id", "shortcut_id", "name") VALUES ($1, $2, $3, $4, $5, $6, $7)`
	for _, command := range macro.GetCommands() {
		_, err = c.Conn.db.Exec(sql, id, c.ObjectId, command.GetIndex(), command.GetType(), command.GetSkillId(), command.GetShortcutId(), command.GetName())
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

	sql := `SELECT id,icon,name,description,acronym FROM macros WHERE char_id=$1`
	rows, err := c.Conn.db.Query(sql, c.ObjectId)
	if err != nil {
		logger.Info.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		m := Macro{}
		err = rows.Scan(&m.Id, &m.Icon, &m.Name, &m.Description, &m.Acronym)
		if err != nil {
			logger.Info.Println(err)
			return
		}
		m.Commands = MacrosesCommands
		Macroses = append(Macroses, m)

		if c.MacroId < m.Id {
			c.MacroId = m.Id
		}
	}
	for index, macros := range Macroses {
		MacrosesCommands = nil
		sqlCommand := `SELECT * FROM macros_commands WHERE command_id=$1 and char_id=$2`
		rowsCommand, err := c.Conn.db.Query(sqlCommand, macros.Id, c.GetObjectId())
		if err != nil {
			logger.Info.Println(err.Error())
			return
		}
		for rowsCommand.Next() {
			cCom := MacroCommand{}
			var test int32
			err = rowsCommand.Scan(&cCom.Id, &cCom.Index, &cCom.Type, &cCom.SkillID, &cCom.ShortcutID, &cCom.Name, &test)
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

func (m *Macro) GetId() int32 {
	return m.Id
}

func (m *Macro) GetName() string {
	return m.Name
}

func (m *Macro) GetDescription() string {
	return m.Description
}

func (m *Macro) GetAcronym() string {
	return m.Acronym
}

func (m *Macro) GetIcon() byte {
	return m.Icon
}

func (m *Macro) GetCount() byte {
	return m.Count
}

func (m *Macro) GetCommands() []interfaces.MacrosCommandInterface {
	commands := make([]interfaces.MacrosCommandInterface, len(m.Commands))
	for index := range m.Commands {
		commands[index] = &m.Commands[index]
	}
	return commands
}

func (mc *MacroCommand) GetId() int32 {
	return mc.Id
}

func (mc *MacroCommand) GetIndex() byte {
	return mc.Index
}

func (mc *MacroCommand) GetType() byte {
	return mc.Type
}

func (mc *MacroCommand) GetSkillId() int32 {
	return mc.SkillID
}

func (mc *MacroCommand) GetShortcutId() byte {
	return mc.ShortcutID
}

func (mc *MacroCommand) GetName() string {
	return mc.Name
}
