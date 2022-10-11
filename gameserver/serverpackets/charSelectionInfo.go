package serverpackets

import (
	"context"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

const InfoAboutCharsByLogin = `SELECT login,object_id,level,max_hp,cur_hp,max_mp,cur_mp,face,hair_style,hair_color,sex,x,y,z,exp,sp,karma,pvp_kills,pk_kills,clan_id,race,class_id,base_class,title,online_time,nobless,vitality,char_name,first_enter_game FROM characters WHERE Login = $1`

// TODO убрать модель
func CharSelectionInfo(clientI interfaces.ReciverAndSender) *packets.Buffer {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return nil
	}

	buffer := packets.Get()

	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), InfoAboutCharsByLogin, client.Account.Login)
	if err != nil {
		logger.Error.Panicln(err)
	}

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
		)
		character.Inventory = models.NewInventory(character.ObjectId)
		if err != nil {
			logger.Error.Panicln(err)
		}
		character.Coordinates = &coord
		character.Conn = client
		client.Account.Char = append(client.Account.Char, character)
	}

	for index := range client.Account.Char {
		client.Account.Char[index].Paperdoll = models.RestoreVisibleInventory(client.Account.Char[index].ObjectId)
	}

	buffer.WriteSingleByte(0x09)
	buffer.WriteD(int32(len(client.Account.Char))) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	buffer.WriteD(7)          //char max number
	buffer.WriteSingleByte(0) // delim

	//todo блок который должен повторяться

	for index := range client.Account.Char {

		buffer.WriteS(client.Account.Char[index].GetName()) // Pers name

		buffer.WriteD(client.Account.Char[index].GetObjectId()) // objId
		buffer.WriteS(client.Account.Char[index].Login)         // loginName

		buffer.WriteD(0)                                 //TODO sessionId
		buffer.WriteD(client.Account.Char[index].ClanId) //clanId
		buffer.WriteD(0)                                 // Builder Level

		buffer.WriteD(client.Account.Char[index].GetSex())         //sex
		buffer.WriteD(int32(client.Account.Char[index].GetRace())) // race
		buffer.WriteD(client.Account.Char[index].GetBaseClass())   // baseclass

		buffer.WriteD(1) // active ??

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

		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)

		paperdoll := models.RestoreVisibleInventory(client.Account.Char[index].GetObjectId())
		for _, slot := range models.GetPaperdollOrder() {
			if paperdoll[slot].Item == nil {
				buffer.WriteD(0)
			} else {
				buffer.WriteD(int32(paperdoll[slot].Id))
			}
		}

		buffer.WriteD(client.Account.Char[index].GetHairStyle()) //hairStyle
		buffer.WriteD(client.Account.Char[index].GetHairColor()) //hairColor
		buffer.WriteD(client.Account.Char[index].GetFace())      // face

		buffer.WriteF(float64(client.Account.Char[index].GetMaxHp())) //max hp
		buffer.WriteF(float64(client.Account.Char[index].GetMaxMp())) // max mp

		buffer.WriteD(0)                                       // days left before
		buffer.WriteD(client.Account.Char[index].GetClassId()) //classId

		buffer.WriteD(0)          //auto-selected
		buffer.WriteSingleByte(0) // enchanted
		buffer.WriteD(0)          //augumented

		buffer.WriteD(0) // Currently on retail when you are on character select you don't see your transformation.

		// Implementing it will be waster of resources.
		buffer.WriteD(0)                                        // Pet ID
		buffer.WriteD(0)                                        // Pet Level
		buffer.WriteD(0)                                        // Pet Max Food
		buffer.WriteD(0)                                        // Pet Current Food
		buffer.WriteF(0)                                        // Pet Max HP
		buffer.WriteF(0)                                        // Pet Max MP
		buffer.WriteD(client.Account.Char[index].GetVitality()) // H5 Vitality

	}

	return buffer
}
