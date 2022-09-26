package serverpackets

import (
	"context"
	"l2gogameserver/data/logger"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/interfaces"
	"l2gogameserver/gameserver/models"
	"l2gogameserver/packets"
)

func CharSelectionInfo(clientI interfaces.ReciverAndSender) []byte {
	client, ok := clientI.(*models.ClientCtx)
	if !ok {
		return []byte{}
	}

	buffer := packets.Get()

	dbConn, err := db.GetConn()
	if err != nil {
		logger.Error.Panicln(err)
	}
	defer dbConn.Release()

	sql := `SELECT login,object_id,level,max_hp,cur_hp,max_mp,cur_mp,face,hair_style,hair_color,sex,x,y,z,exp,sp,karma,pvp_kills,pk_kills,clan_id,race,class_id,base_class,title,online_time,nobless,vitality,char_name,first_enter_game FROM characters WHERE Login = $1`
	rows, err := dbConn.Query(context.Background(), sql, client.Account.Login)
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

	for _, v := range client.Account.Char {
		v.Paperdoll = models.RestoreVisibleInventory(v.ObjectId)
	}

	buffer.WriteSingleByte(0x09)
	buffer.WriteD(int32(len(client.Account.Char))) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	buffer.WriteD(7)          //char max number
	buffer.WriteSingleByte(0) // delim

	//todo блок который должен повторяться

	for _, char := range client.Account.Char {

		buffer.WriteS(char.CharName) // Pers name

		buffer.WriteD(char.ObjectId) // objId
		buffer.WriteS(char.Login)    // loginName

		buffer.WriteD(0)           //TODO sessionId
		buffer.WriteD(char.ClanId) //clanId
		buffer.WriteD(0)           // Builder Level

		buffer.WriteD(char.Sex)         //sex
		buffer.WriteD(int32(char.Race)) // race
		buffer.WriteD(char.BaseClass)   // baseclass

		buffer.WriteD(1) // active ??

		x, y, z := char.GetXYZ()
		buffer.WriteD(x) //x 53
		buffer.WriteD(y) //y 57
		buffer.WriteD(z) //z 61

		buffer.WriteF(float64(char.CurHp)) //currentHP
		buffer.WriteF(float64(char.CurMp)) //currentMP

		buffer.WriteD(char.Sp)                                               // SP
		buffer.WriteQ(int64(char.Exp))                                       // EXP
		buffer.WriteF(char.GetPercentFromCurrentLevel(char.Exp, char.Level)) // percent
		buffer.WriteD(char.Level)                                            // level

		buffer.WriteD(char.Karma)    // karma
		buffer.WriteD(char.PkKills)  // pk
		buffer.WriteD(char.PvpKills) //pvp

		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)
		buffer.WriteD(0)

		paperdoll := models.RestoreVisibleInventory(char.ObjectId)
		for _, slot := range models.GetPaperdollOrder() {
			buffer.WriteD(int32(paperdoll[slot].Id))
		}

		buffer.WriteD(char.HairStyle) //hairStyle
		buffer.WriteD(char.HairColor) //hairColor
		buffer.WriteD(char.Face)      // face

		buffer.WriteF(float64(char.MaxHp)) //max hp
		buffer.WriteF(float64(char.MaxMp)) // max mp

		buffer.WriteD(0)            // days left before
		buffer.WriteD(char.ClassId) //classId

		buffer.WriteD(1)          //auto-selected
		buffer.WriteSingleByte(0) // enchanted
		buffer.WriteD(0)          //augumented

		buffer.WriteD(0) // Currently on retail when you are on character select you don't see your transformation.

		// Implementing it will be waster of resources.
		buffer.WriteD(0)             // Pet ID
		buffer.WriteD(0)             // Pet Level
		buffer.WriteD(0)             // Pet Max Food
		buffer.WriteD(0)             // Pet Current Food
		buffer.WriteF(0)             // Pet Max HP
		buffer.WriteF(0)             // Pet Max MP
		buffer.WriteD(char.Vitality) // H5 Vitality

	}

	defer packets.Put(buffer)
	return buffer.Bytes()
}
