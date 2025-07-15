package serverpackets

import (
	"database/sql"
	"l2gogameserver/data/logger"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/gameserver/models/initial"
	"l2gogameserver/gameserver/models/items"
	"l2gogameserver/packets"
	"log"
)

const InfoAboutCharsByLogin = `SELECT login,object_id,level,max_hp,cur_hp,max_mp,cur_mp,face,hair_style,hair_color,sex,x,y,z,exp,sp,karma,pvp_kills,pk_kills,clan_id,race,class_id,base_class,title,online_time,nobless,vitality,char_name,first_enter_game,last_enter_world FROM characters WHERE Login = $1 ORDER BY object_id`

// TODO убрать модель
func CharSelectionInfo(clientI interfaces.NewClientCtxInterface, db *sql.DB) *packets.Buffer {
	isSelectPlayer := false

	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return nil
	}

	buffer := packets.Get()

	rows, err := db.Query(InfoAboutCharsByLogin, client.Account.Login)
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer rows.Close()
	//
	client.Account.Char = client.Account.Char[:0]
	for rows.Next() {
		var character = models.GetNewCharacterModel()
		var coord models.Coordinates
		err = rows.Scan(
			&character.Login,
			&character.ObjectId,
			&character.Level,
			&character.MaxHp,
			&character.CurHp,
			&character.MaxMp,
			&character.CurMp,
			&character.Face,
			&character.HairStyle,
			&character.HairColor,
			&character.Sex,
			&coord.X,
			&coord.Y,
			&coord.Z,
			&character.Exp,
			&character.Sp,
			&character.Karma,
			&character.PvpKills,
			&character.PkKills,
			&character.ClanId,
			&character.Race,
			&character.ClassId,
			&character.BaseClass,
			&character.Title,
			&character.OnlineTime,
			&character.Nobless,
			&character.Vitality,
			&character.CharName,
			&character.FirstEnterGame,
			&character.LastEnterWorld,
		)
		character.Inventory = models.NewInventory(character.ObjectId)
		// Устанавливаем ссылку на персонажа в инвентаре
		character.Inventory.SetCharacter(character)
		if err != nil {
			logger.Error.Panicln(err)
		}
		character.Coordinates = &coord
		character.Conn = client
		client.Account.Char = append(client.Account.Char, character)
	}

	// Узнаем последнего активного персонажа, чтоб он в лобби был по умолчанию
	lastActiveChar := client.GetAccount().GetLastActiveChar()

	for _, char := range client.Account.Char {
		if char.IsFirstEnterGame() {
			//Выдача предметов после создания персонажа
			eq := initial.GetEquipmentByClass(char.GetBaseClass())
			if eq != nil {
				for _, item := range eq.Items {
					itemData, _ := items.GetItemInfo(item.Id)
					AddItem2 := char.GetInventory().AddItem2(int32(item.Id), item.Count, itemData.IsStackable(), db)
					if item.Equipped {
						char.GetInventory().GetItemByObjectId(AddItem2.GetObjectId()).UseEquippableItem()
						for _, itemInterface := range char.GetInventory().GetItems() {
							itemInterface.UpdateDB(db)
						}
					}
				}

				// ======================= ВАШ ДЕБАГ-МАРКЕР №2 =======================
				log.Printf("====== ПРОВЕРКА ИНВЕНТАРЯ ПОСЛЕ ЭКИПИРОВКИ ДЛЯ ЧАРА '%s' (ID: %d) ======", char.GetName(), char.GetObjectId())
				for _, invItem := range char.GetInventory().GetItems() {
					// Выводим только те предметы, которые должны быть надеты
					if invItem.GetLocation() == "PAPERDOLL" {
						log.Printf("ПРОВЕРКА: Предмет '%s' (ObjectID: %d) находится в слоте %d (Location: %s)", invItem.GetId(), invItem.GetObjectId(), invItem.GetLocData(), invItem.GetLocation())
					}
				}
				log.Println("=========================== КОНЕЦ ПРОВЕРКИ ===========================")
				// ========================================================================

			}

			char.SaveFirstInGamePlayer(db)
		}
	}

	buffer.WriteSingleByte(0x09)
	buffer.WriteD(int32(len(client.Account.Char))) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	buffer.WriteD(7) //char max number

	// Если достигнут лимит — клиенту нельзя создавать новых персонажей (0), иначе (1) клиент будет спрашивать (0x13), сервер ответит (0x0D)
	canCreate := byte(0x00)
	if len(client.Account.Char) < 7 {
		canCreate = 0x01
	}
	buffer.WriteSingleByte(canCreate)

	//todo блок который должен повторяться
	for index := range client.Account.Char {
		isSelectPlayer = false
		if lastActiveChar != nil {
			if lastActiveChar.GetObjectId() == client.Account.Char[index].GetObjectId() {
				isSelectPlayer = true
			}
		}
		buffer.WriteS(client.Account.Char[index].GetName()) // Pers name

		buffer.WriteD(client.Account.Char[index].GetObjectId()) // objId
		buffer.WriteS(client.Account.Login)                     // loginName

		buffer.WriteD(0)                                      //TODO sessionId
		buffer.WriteD(client.Account.Char[index].GetClanId()) //clanId
		buffer.WriteD(0x00)                                   // Builder Level

		buffer.WriteD(client.Account.Char[index].GetSex())         //sex
		buffer.WriteD(int32(client.Account.Char[index].GetRace())) // race

		if client.GetAccount().GetCurrentChar().GetClassId() == client.GetAccount().GetCurrentChar().GetBaseClass() {
			buffer.WriteD(client.Account.Char[index].GetClassId())
		} else {
			buffer.WriteD(client.Account.Char[index].GetBaseClass())
		}

		buffer.WriteD(0x01) // active ??

		x, y, z := client.Account.Char[index].GetXYZ()
		buffer.WriteD(x) //x 53
		buffer.WriteD(y) //y 57
		buffer.WriteD(z) //z 61

		buffer.WriteF(float64(client.Account.Char[index].GetCurrentHp())) //currentHP
		buffer.WriteF(float64(client.Account.Char[index].GetCurrentMp())) //currentMP

		buffer.WriteD(client.Account.Char[index].GetCurrentSp())
		currentExp := client.Account.Char[index].GetCurrentExp()

		buffer.WriteQ(int64(currentExp))
		level := client.Account.Char[index].GetLevel()
		buffer.WriteF(client.Account.Char[index].GetPercentFromCurrentLevel(currentExp, level)) // percent
		buffer.WriteD(level)                                                                    // level

		buffer.WriteD(client.Account.Char[index].GetKarma()) // karma
		buffer.WriteD(client.Account.Char[index].GetPK())    // pk
		buffer.WriteD(client.Account.Char[index].GetPVP())   //pvp

		buffer.WriteD(0x00)
		buffer.WriteD(0x00)
		buffer.WriteD(0x00)
		buffer.WriteD(0x00)
		buffer.WriteD(0x00)
		buffer.WriteD(0x00)
		buffer.WriteD(0x00)

		paperdoll := models.BuildPaperdollFromInventory(client.Account.Char[index].GetInventory())
		if !client.GetAccount().GetCurrentChar().IsFirstEnterGame() {
			paperdoll = models.RestoreVisibleInventory(client.Account.Char[index].GetObjectId(), db)
		}

		for _, slot := range models.GetPaperdollOrder() {
			if paperdoll[slot].GetItemInfo() == nil {
				buffer.WriteD(0x00)
			} else {
				buffer.WriteD(paperdoll[slot].GetItemInfo().GetId())
			}
		}

		buffer.WriteD(client.Account.Char[index].GetHairStyle()) //hairStyle
		buffer.WriteD(client.Account.Char[index].GetHairColor()) //hairColor
		buffer.WriteD(client.Account.Char[index].GetFace())      // face

		buffer.WriteF(float64(client.Account.Char[index].GetMaxHp())) //max hp
		buffer.WriteF(float64(client.Account.Char[index].GetMaxMp())) // max mp

		buffer.WriteD(0x00) // days left before

		buffer.WriteD(client.Account.Char[index].GetClassId()) //classId

		if isSelectPlayer {
			buffer.WriteD(0x01)
		} else {
			buffer.WriteD(0x00)
		}

		buffer.WriteSingleByte(0x00) // enchanted
		buffer.WriteD(0x00)          //augumented

		buffer.WriteD(0x00) // Currently on retail when you are on character select you don't see your transformation.

		// Implementing it will be waster of resources.
		buffer.WriteD(0x00)                                     // Pet ID
		buffer.WriteD(0x00)                                     // Pet Level
		buffer.WriteD(0x00)                                     // Pet Max Food
		buffer.WriteD(0x00)                                     // Pet Current Food
		buffer.WriteF(0x00)                                     // Pet Max HP
		buffer.WriteF(0x00)                                     // Pet Max MP
		buffer.WriteD(client.Account.Char[index].GetVitality()) // H5 Vitality

	}

	return buffer
}
