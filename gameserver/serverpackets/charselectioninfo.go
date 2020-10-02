package serverpackets

import (
	"github.com/jackc/pgx"
	"l2gogameserver/gameserver/models"
	"log"
)

func NewCharSelectionInfo(db *pgx.Conn, client *models.Client) *models.Account {

	rows, err := db.Query("SELECT * FROM characters WHERE Login = $1", []byte(client.Account.Login))
	if err != nil {
		log.Fatal(err)
	}

	var account models.Account
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
			log.Fatal(err)
		}
		character.Coordinates = &coord
		account.Char = append(account.Char, &character)
	}

	client.Buffer.WriteH(0) //reserve
	client.Buffer.WriteSingleByte(0x09)
	client.Buffer.WriteD(int32(len(account.Char))) //size char in account

	// Can prevent players from creating new characters (if 0); (if 1, the client will ask if chars may be created (0x13) Response: (0x0D) )
	client.Buffer.WriteD(7)          //char max number
	client.Buffer.WriteSingleByte(0) // delim

	//todo блок который должен повторяться

	for _, char := range account.Char {

		client.Buffer.WriteS(string(char.CharName.Bytes)) // Pers name

		client.Buffer.WriteD(char.CharId)              // objId
		client.Buffer.WriteS(string(char.Login.Bytes)) // loginName

		client.Buffer.WriteD(0)           //TODO sessionId
		client.Buffer.WriteD(char.ClanId) //clanId
		client.Buffer.WriteD(0)           // Builder Level

		client.Buffer.WriteD(char.Sex)       //sex
		client.Buffer.WriteD(char.Race)      // race
		client.Buffer.WriteD(char.BaseClass) // baseclass

		client.Buffer.WriteD(0) // active ??

		client.Buffer.WriteD(char.Coordinates.X) //x 53
		client.Buffer.WriteD(char.Coordinates.Y) //y 57
		client.Buffer.WriteD(char.Coordinates.Z) //z 61

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
		//
		//
		m := make([]byte, 104)
		client.Buffer.WriteSlice(m)

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

	return &account
}
