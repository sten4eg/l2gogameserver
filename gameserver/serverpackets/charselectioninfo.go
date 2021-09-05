package serverpackets

import (
	"context"
	"l2gogameserver/db"
	"l2gogameserver/gameserver/models"
)

func CharSelectionInfo(client *models.Client) {

	dbConn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Release()

	rows, err := dbConn.Query(context.Background(), "SELECT * FROM characters WHERE Login = $1", []byte(client.Account.Login))
	if err != nil {
		panic(err)
	}

	//
	client.Account.Char = client.Account.Char[:0]
	for rows.Next() {
		var character models.Character
		var coord models.Coordinates
		err = rows.Scan(
			&character.Login,
			&character.CharId,
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
		)
		if err != nil {
			panic(err)
		}
		character.Coordinates = &coord
		character.Conn = client
		models.SetupStats(&character)
		client.Account.Char = append(client.Account.Char, &character)
	}

	for _, v := range client.Account.Char {
		v.Paperdoll = models.RestoreVisibleInventory(v.CharId)
	}

	client.Buffer.WriteSingleByte(0x09)
	client.Buffer.WriteD(int32(len(client.Account.Char))) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	client.Buffer.WriteD(7)          //char max number
	client.Buffer.WriteSingleByte(0) // delim

	//todo блок который должен повторяться

	for _, char := range client.Account.Char {

		client.Buffer.WriteS(string(char.CharName.Bytes)) // Pers name

		client.Buffer.WriteD(char.CharId)              // objId
		client.Buffer.WriteS(string(char.Login.Bytes)) // loginName

		client.Buffer.WriteD(0)           //TODO sessionId
		client.Buffer.WriteD(char.ClanId) //clanId
		client.Buffer.WriteD(0)           // Builder Level

		client.Buffer.WriteD(char.Sex)         //sex
		client.Buffer.WriteD(int32(char.Race)) // race
		client.Buffer.WriteD(char.BaseClass)   // baseclass

		client.Buffer.WriteD(1) // active ??

		x, y, z := char.GetXYZ()
		client.Buffer.WriteD(x) //x 53
		client.Buffer.WriteD(y) //y 57
		client.Buffer.WriteD(z) //z 61

		client.Buffer.WriteF(float64(char.CurHp)) //currentHP
		client.Buffer.WriteF(float64(char.CurMp)) //currentMP

		client.Buffer.WriteD(char.Sp)                                               // SP
		client.Buffer.WriteQ(int64(char.Exp))                                       // EXP
		client.Buffer.WriteF(char.GetPercentFromCurrentLevel(char.Exp, char.Level)) // percent
		client.Buffer.WriteD(char.Level)                                            // level

		client.Buffer.WriteD(char.Karma)    // karma
		client.Buffer.WriteD(char.PkKills)  // pk
		client.Buffer.WriteD(char.PvpKills) //pvp

		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)
		client.Buffer.WriteD(0)

		paperdoll := models.RestoreVisibleInventory(char.CharId)
		for _, slot := range models.GetPaperdollOrder() {
			client.Buffer.WriteD(int32(paperdoll[slot].Id))
		}

		client.Buffer.WriteD(char.HairStyle) //hairStyle
		client.Buffer.WriteD(char.HairColor) //hairColor
		client.Buffer.WriteD(char.Face)      // face

		client.Buffer.WriteF(float64(char.MaxHp)) //max hp
		client.Buffer.WriteF(float64(char.MaxMp)) // max mp

		client.Buffer.WriteD(0)            // days left before
		client.Buffer.WriteD(char.ClassId) //classId

		client.Buffer.WriteD(1)          //auto-selected
		client.Buffer.WriteSingleByte(0) // enchanted
		client.Buffer.WriteD(0)          //augumented

		client.Buffer.WriteD(0) // Currently on retail when you are on character select you don't see your transformation.

		// Implementing it will be waster of resources.
		client.Buffer.WriteD(0)             // Pet ID
		client.Buffer.WriteD(0)             // Pet Level
		client.Buffer.WriteD(0)             // Pet Max Food
		client.Buffer.WriteD(0)             // Pet Current Food
		client.Buffer.WriteF(0)             // Pet Max HP
		client.Buffer.WriteF(0)             // Pet Max MP
		client.Buffer.WriteD(char.Vitality) // H5 Vitality

	}
	client.SaveAndCryptDataInBufferToSend(true)
}
